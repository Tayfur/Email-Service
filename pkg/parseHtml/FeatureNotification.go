package parseHtml

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)
// htmli parse ediyoruz bazi htmller parametre alabilir bunun ornegi icin weeklyreporttemplate bakmanizi oneririm
func FeatureNotificationTemplate() string {
	var templateBuffer bytes.Buffer
	t,err := template.ParseFiles("./Templates/FeatureNotification.html")
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	//subject gonderilen mailin titleni iceriyor
	templateBuffer.Write([]byte(fmt.Sprintf("Subject: Lates Features \n%s\n\n", mimeHeaders)))
	t.Execute(&templateBuffer, nil)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	return templateBuffer.String()

}