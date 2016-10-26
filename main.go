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
/*
var bot *linebot.Client
var eggyoID = "ufa92a3a52f197e19bfddeb5ca0595e93"
var logNof = "open"

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
}*/

func main() {

	bot, err := linebot.New(
		os.Getenv("ChannelSecret"),
		os.Getenv("MID"),
	)
	if err != nil {
		log.Fatal(err)
	}
	// Setup HTTP Server for receiving requests from LINE platform
		http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
			events, err := bot.ParseRequest(req)
			if err != nil {
				if err == linebot.ErrInvalidSignature {
					w.WriteHeader(400)
				} else {
					w.WriteHeader(500)
				}
				return
			}
			for _, event := range events {
				if event.Type == linebot.EventTypeMessage {
					switch message := event.Message.(type) {
					case *linebot.TextMessage:
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
							log.Print(err)
						}
					}
				} 
			}
		})
		// This is just sample code.
		// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
		if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
			log.Fatal(err)
		}
}
/*
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
			out := fmt.Sprintf("Bot แปลงพิกัด Eggyo\nวิธีใช้\nเพียงแค่กดแชร์ Location ที่ต้องการ ระบบจะทำการแปลง Location เป็นพิกัดระบบต่างๆ และหาความสูงจากระดับน้ำทะเลให้\n\nหรือจะพูดคุยกับ bot ก็ได้\nกด #help เพื่อดูวิธีใช้อื่นๆ \nติดต่อผู้พัฒนา LINE ID : eggyo")
			//result.RawContent.Params[0] is who send your bot friend added operation, otherwise you cannot get in content or operation content.
			_, err = bot.SendText([]string{content.From}, out)
			if logNof == "open" {
				bot.SendText([]string{eggyoID}, "bot has a new friend :"+content.From)
			}

			addNewUser(content.From)

			if err != nil {
				log.Println(err)
			}
		}

		if content != nil && content.IsMessage && content.ContentType == linebot.ContentTypeText {

			text, err := content.TextContent()
			if logNof == "open" {
				bot.SendText([]string{eggyoID}, "bot get msg:"+text.Text+"\nfrom :"+content.From)
			}
			// reply message
			var processedText = messageCheck(text.Text)
			_, err = bot.SendText([]string{content.From}, processedText)

			if err != nil {
				log.Println(err)
			}
		}
		if content != nil && content.ContentType == linebot.ContentTypeLocation {
			_, err = bot.SendText([]string{content.From}, "ระบบกำลังประมวลผล...")

			loc, err := content.LocationContent()

			// add eggyo geo test//
			resp, err := http.Get("http://eggyo-geo-node.herokuapp.com/geo/" + FloatToString(loc.Latitude) + "," + FloatToString(loc.Longitude))
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
			_, err = bot.SendText([]string{content.From}, "Utm :"+geo.Results.Utm+"\n\nMgrs :"+geo.Results.Mgrs+"\n\nAltitude :"+elev)
			if logNof == "open" {
				bot.SendText([]string{eggyoID}, "bot get loc:"+geo.Results.Mgrs+"\nfrom :"+content.From)
			}
			if err != nil {
				log.Println(err)
			}
		}
	}
	*/
}
