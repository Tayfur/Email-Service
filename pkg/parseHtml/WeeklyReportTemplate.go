package parseHtml

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)
func WeeklyReportTemplate() string {
	var templateBuffer bytes.Buffer
	t,err := template.ParseFiles("./Templates/WeeklyReport.html")
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	templateBuffer.Write([]byte(fmt.Sprintf("Subject:Weekly Report  \n%s\n\n", mimeHeaders)))
	t.Execute(&templateBuffer, WeeklyReport{
		Name:    "tayfur",
		WorksapeceName: "Facebook",
		Cpu:80,
		Ram:80,
		Storage: 200,
		Price: 1000,
	})

	if err != nil {
		log.Fatal(err)
		return ""
	}

	return templateBuffer.String()

}