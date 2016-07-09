package main

import "strings"

func messageCheck(msg string) string {
	var result = ""
	if msg[0:4] == "#ask" {
		// train
		msg = strings.Trim(msg, "#ask ")
		msg = strings.Replace(msg, " #ans ", ":", 1)
		var msgArray = strings.Split(msg, ":")
		if checkNewMessage(msgArray[0]) == true {
			addNewMessageFromUser(msgArray[0], msgArray[1])
		} else {
			addReplyMessageFromUser(msgArray[0], msgArray[1])
		}
		result = "ข้าจำได้แล้ว ลองทักใหม่ซิ อิอิ"

	} else if msg == "#help" {
		result = "ควย เอ้ย! คอย"
	} else {
		result = getReplyMessageFromUser(msg)
		//result = msg
	}

	return result
}
