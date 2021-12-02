package parseHtml

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)

func WelcomeTemplate() string {
	var templateBuffer bytes.Buffer
	t,err := template.ParseFiles("./Templates/Welcome.html")
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	templateBuffer.Write([]byte(fmt.Sprintf("Subject: Welcome to Tayfur\n%s\n\n", mimeHeaders)))
	t.Execute(&templateBuffer, nil)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	return templateBuffer.String()

}