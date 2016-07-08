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

var url = "https://api.mlab.com/api/1/databases/heroku_h1g317z7/collections/_User?apiKey=1S26M0Ti2t7gKunYRJiGNg8aeIMXnptN"

func getAllUser() {

	resp, err := http.Get(url)
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
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
