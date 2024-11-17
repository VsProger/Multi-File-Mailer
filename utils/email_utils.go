package utils

import "strings"

func ParseEmails(emails string) []string {
	emailList := strings.Split(emails, ",")

	for i := range emailList {
		emailList[i] = strings.TrimSpace(emailList[i])
	}

	return emailList
}
