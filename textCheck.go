package main

func messageCheck(msg string) string {
	var result = ""
	if msg == "#help" {
		result = "ควย"
	} else {
		result = getReplyMessageFromUser(msg)
	}

	return result
}
