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
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"net/url"
  	"io/ioutil"
	"github.com/line/line-bot-sdk-go/linebot"
        "encoding/json"  
)

var bot *linebot.Client

type GeoContent struct {
	LatLong string `json:"latLon"`
	Utm string `json:"utm"`
        Mgrs string `json:"mgrs"`
}

type ResultGeoLoc struct {
	Results GeoContent `json:"result"`
}
func getGeoLoc(body []byte) (*ResultGeoLoc, error) {
    var s = new(ResultGeoLoc)
    err := json.Unmarshal(body, &s)
    if(err != nil){
        fmt.Println("whoops:", err)
    }
    return s, err
}

func main() {
	// fixie
	fixieUrl, err := url.Parse(os.Getenv("FIXIE_URL"))
  	customClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(fixieUrl)}}
  	resp, err := customClient.Get("http://welcome.usefixie.com")
  	if (err != nil) {
    		println(err.Error())
    		return
  	}
  	defer resp.Body.Close()
  	body, err := ioutil.ReadAll(resp.Body)
  	println(string(body))
	
	// end fixie
	
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
		/*
		if content != nil && content.IsMessage && content.ContentType == linebot.ContentTypeText {
			text, err := content.TextContent()
			_, err = bot.SendText([]string{content.From}, "Bot received : "+text.Text)

			if err != nil {
				log.Println(err)
			}
		}*/
		if content != nil && content.ContentType == linebot.ContentTypeLocation {
			loc, err := content.LocationContent()

			// add eggyo geo test
			resp, err := http.Get("http://eggyo-geo-node.herokuapp.com/geo/" + FloatToString(loc.Latitude) + "/" + FloatToString(loc.Longitude))
			if (err != nil) {
    				println(err.Error())
    				return
  			}
  			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
                       log.Println(string(body))
			
			var elev = callGoogleElev(loc.Latitude,loc.Longitude)
                       geo, err := getGeoLoc([]byte(body))
			_, err = bot.SendText([]string{content.From}, "LatLong :" + geo.Results.LatLong)
			_, err = bot.SendText([]string{content.From}, "Utm :" + geo.Results.Utm)
			_, err = bot.SendText([]string{content.From}, "Mgrs :" + geo.Results.Mgrs)
			_, err = bot.SendText([]string{content.From}, "Altitude :" + elev)

                        
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
func callGoogleElev(lat float64,lon float64) string {
	resp, err := http.Get("https://maps.googleapis.com/maps/api/elevation/json?locations=" + FloatToString(lat) + "," + FloatToString(lon) + "&key=AIzaSyAn9cWoce9zGEfGjDzMg6r_uTTUw3WoMOg")
	if (err != nil) {
    		println(err.Error())
    		return string(err.Error())
  	}
  	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
       elev, err := getElev([]byte(body))

	return FloatToString(elev.Results[0].Elev)
}
type ELEV struct {
	Elev float64 `json:”elevation”`
	Loc string `json:”location”`
        Res float64 `json:”resolution”`
}

type ResultElev struct {
	Results []ELEV `json:"results”`
}
func getElev(body []byte) (*ResultElev, error) {
    var s = new(ResultElev)
    err := json.Unmarshal(body, &s)
    if(err != nil){
        fmt.Println("whoops:", err)
    }
    return s, err
}
//eggyo
