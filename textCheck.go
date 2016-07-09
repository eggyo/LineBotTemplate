package main

import "strings"

func messageCheck(msg string) string {
	var result = ""
	if msg[0:4] == "#สอน" {
		// train
		msg = strings.Trim(msg, "#สอน ")
		msg = strings.Replace(msg, " #ตอบ ", ":", 1)
		var msgArray = strings.Split(msg, ":")
		addNewMessageFromUser(msgArray[0], msgArray[1])

	}
	if msg == "#help" {
		result = "ควย"
	} else {
		result = getReplyMessageFromUser(msg)
		//result = msg
	}

	return result
}
