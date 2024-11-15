package utils

import "strings"

func ParseEmails(emails string) []string {
	// Разделяем строку по запятым
	emailList := strings.Split(emails, ",")

	// Удаляем пробелы вокруг каждого адреса
	for i := range emailList {
		emailList[i] = strings.TrimSpace(emailList[i])
	}

	return emailList
}
