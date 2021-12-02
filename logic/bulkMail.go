package logic

import (
	"SendEmail-Service/pkg/db"
	"SendEmail-Service/pkg/models"
	"SendEmail-Service/pkg/rabbitmq"
	"log"
)

func BulkMail(template string) error {
	// database baglaniyoruz
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	var User []models.User
	db.Find(&User)
	for _, t := range User {
		// to parametresini queue'a gondermek icin byte dizisi olarak aliyoruz
		var to []byte
		to = []byte(t.Email)
		rabbitmq.Publisher(to, template)
		// db deki userların email adreslerini alıp rabbitmq'e gönderiyoruz , her mail icin bir mesaj gönderiyoruz!!
	}

	return nil
}
