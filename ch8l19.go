package main

import "strings"
import "fmt"

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	for i, e := range messages {
		// fmt.Println(tagger(i))
		messages[i].tags = tagger(e)
		fmt.Println(e)
	}
	fmt.Println(messages)
	return messages
}

func tagger(msg sms) []string {
	tags := []string{}
	lowerStr := strings.ToLower(msg.content)
	if strings.Contains(lowerStr, "urgent") {
		tags = append(tags, "Urgent")
	}
	if strings.Contains(lowerStr, "sale") {
		tags = append(tags, "Promo")
	}
	return tags
}
