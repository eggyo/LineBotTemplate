package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ObjectId struct {
	ObjId string `json:"$oid"`
}
type USER struct {
	LineID string   `json:"lineID"`
	ID     ObjectId `json:"_id"`
}
type MESSAGE struct {
	Msg       string   `json:"msg"`
	From      string   `json:"fromUserId"`
	ReplyBool bool     `json:"replyBool"`
	ReplyMsg  []string `json:"replyMsg"`
	ID        ObjectId `json:"_id"`
}
type ResultsMESSAGE struct {
	ContentMsg []MESSAGE
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
	var sendingMsg = `{"lineID":"` + ID + `"}`
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
	var sendingMsg = `{"msg":"` + msg + `","fromUserId":"` + fromUserId + `","replyBool":false,"replyMsg":[""]}`
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
	body, err := ioutil.ReadAll(resp.Body)
	log.Println("obj :", string(body))

	msgObj, err := messageGet([]byte(body))
	log.Println("objId msg :", msgObj.ID.ObjId)
}

func getReplyMessageFromUser(msg string) string {
	var q = `&q={"msg":"` + msg + `","replyBool":true}`
	resp, err := http.Get(msgDb_url + q)
	log.Println("Query :", msgDb_url+q)

	if err != nil {
		println(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Println("reply body:", string(body))

	msgObjs, err := messageArrayGet([]byte(body))
	log.Println("reply :", msgObjs)
	if len(msgObjs.ContentMsg) == 0 {
		return "ข้าไม่เข้าใจ ถ้าอยากสอนชั้น ให้ทำตามนี้\n\nพิมพ์ #สอน ข้อความ #ตอบ ข้อความที่จะให้ตอบ\n\nเช่น\n#สอน หวัดดี #ตอบ จ้า"
	} else {
		return msgObjs.ContentMsg[0].ReplyMsg[0]
	}
}

func messageGet(body []byte) (*MESSAGE, error) {
	var s = new(MESSAGE)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}
func messageArrayGet(body []byte) (*ResultsMESSAGE, error) {
	var s = new(ResultsMESSAGE)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}
