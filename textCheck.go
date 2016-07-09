package main

import "strings"

func messageCheck(msg string) string {
	var result = ""
	if isTrainingCommand(msg) {
		// train
		messsageDeploy(msg)
		result = "ข้าจำได้แล้ว ลองทักใหม่ซิ อิอิ"

	} else if msg == "#help" {
		result = "ควย เอ้ย! คอย"
	} else if isBotCommand(msg) {
		botCommandProcessing(msg)
	} else {
		result = getReplyMessageFromUser(msg)
		//result = msg
	}

	return result
}
func isTrainingCommand(msg string) bool {
	if len(msg) > 6 && msg[0:4] == "#ask" {
		return true
	} else {
		return false
	}
}
func messsageDeploy(msg string) {
	msg = strings.Trim(msg, "#ask ")
	msg = strings.Replace(msg, " #ans ", ":", 1)
	var msgArray = strings.Split(msg, ":")
	if checkNewMessage(msgArray[0]) == true {
		addNewMessageFromUser(msgArray[0], msgArray[1])
	} else {
		addReplyMessageFromUser(msgArray[0], msgArray[1])
	}
}
func isBotCommand(msg string) bool {
	if len(msg) > 6 && msg[0:4] == "#bot" {
		return true
	} else {
		return false
	}
}
func botCommandProcessing(msg string) {
	var result = ""
	msg = strings.Trim(msg, "#bot ")
	var msgArray = strings.Split(msg, " ")
	if msgArray[0] == "nof" {
		logNof = msgArray[1]
		if msgArray[1] == "close" {
			result = "close log notification"
		} else {
			result = "open log notification"
		}
	} else if msgArray[0] == "help" {
		// show help command list
		result = "#bot nof close -> ปิดแจ้งเตือน\n#bot nof open -> เปิดแจ้งเตือน\n#bot send <userId> <msg> -> ส่งข้อความให้คนอื่น\n#bot broadcast <msg> -> ประกาศ\n"
	} else if msgArray[0] == "send" {
		// send content to user
	} else if msgArray[0] == "broadcast" {
		// broadcast msg to all user
	} else if msgArray[0] == "get" {
		// send content to user
	} else {

	}
	bot.SendText([]string{eggyoID}, result)

}
