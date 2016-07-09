package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

type USER struct {
	Name   string
	LineID string
}

var userDb_url = "https://api.mlab.com/api/1/databases/heroku_h1g317z7/collections/_User?apiKey=1S26M0Ti2t7gKunYRJiGNg8aeIMXnptN"
var msgDb_url = "https://api.mlab.com/api/1/databases/heroku_h1g317z7/collections/Message?apiKey=1S26M0Ti2t7gKunYRJiGNg8aeIMXnptN"

func getAllUser() {

	resp, err := http.Get(userDb_url)
	if err != nil {
		println(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Println("mLab User", string(body))

}
func addNewUser(ID string) {
	var sendingMsg = `{"lineID":` + ID + `}`
	var jsonStr = []byte(sendingMsg)
	req, err := http.NewRequest("POST", userDb_url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func addMessageFromUser(msg string, fromUserId string) {
	var sendingMsg = `{"msg":` + msg + `,"fromUserId":` + fromUserId + `,"replyBool":false,"replyMsg":""}`
	log.Println(sendingMsg)
	var jsonStr = []byte(sendingMsg)
	req, err := http.NewRequest("POST", msgDb_url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
