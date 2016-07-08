// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

type GeoContent struct {
	LatLong string `json:"latLon"`
	Utm     string `json:"utm"`
	Mgrs    string `json:"mgrs"`
}

type ResultGeoLoc struct {
	Results GeoContent `json:"result"`
}

func getGeoLoc(body []byte) (*ResultGeoLoc, error) {
	var s = new(ResultGeoLoc)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}

func main() {

	// line bot
	strID := os.Getenv("ChannelID")
	numID, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		log.Fatal("Wrong environment setting about ChannelID")
	}

	bot, err = linebot.NewClient(numID, os.Getenv("ChannelSecret"), os.Getenv("MID"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {

	received, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	for _, result := range received.Results {
		content := result.Content()
		log.Println("-->", content)

		//Log detail receive content
		if content != nil {
			log.Println("RECEIVE Msg:", content.IsMessage, " OP:", content.IsOperation, " type:", content.ContentType, " from:", content.From, "to:", content.To, " ID:", content.ID)
		}
		// user add friend
		if content != nil && content.IsOperation && content.OpType == linebot.OpTypeAddedAsFriend {
			out := fmt.Sprintf("Bot แปลงพิกัด Eggyo\nวิธีใช้\nเพียงแค่กดแชร์ Location ที่ต้องการ ระบบจะทำการแปลง Location เป็นพิกัดระบบต่างๆ และหาความสูงจากระดับน้ำทะเลให้\nกด #help เพื่อดูวิธีใช้อื่นๆ :))")
			//result.RawContent.Params[0] is who send your bot friend added operation, otherwise you cannot get in content or operation content.
			_, err = bot.SendText([]string{result.RawContent.Params[0]}, out)
			if err != nil {
				log.Println(err)
			}
		}

		if content != nil && content.IsMessage && content.ContentType == linebot.ContentTypeText {
			text, err := content.TextContent()
			var processedText = messageCheck(text.Text)
			_, err = bot.SendText([]string{content.From}, processedText)

			if err != nil {
				log.Println(err)
			}
		}
		if content != nil && content.ContentType == linebot.ContentTypeLocation {
			_, err = bot.SendText([]string{content.From}, "ระบบกำลังประมวลผล...")

			loc, err := content.LocationContent()

			// add eggyo geo test
			resp, err := http.Get("http://eggyo-geo-node.herokuapp.com/geo/" + FloatToString(loc.Latitude) + "/" + FloatToString(loc.Longitude))
			if err != nil {
				println(err.Error())
				return
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			log.Println(string(body))

			var elev = callGoogleElev(loc.Latitude, loc.Longitude)
			geo, err := getGeoLoc([]byte(body))
			_, err = bot.SendText([]string{content.From}, "LatLong :"+geo.Results.LatLong)
			_, err = bot.SendText([]string{content.From}, "Utm :"+geo.Results.Utm)
			_, err = bot.SendText([]string{content.From}, "Mgrs :"+geo.Results.Mgrs)
			_, err = bot.SendText([]string{content.From}, "Altitude :"+elev)

			if err != nil {
				log.Println(err)
			}
		}
	}
}
func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

//eggyo
