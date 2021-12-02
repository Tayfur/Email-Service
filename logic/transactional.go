package logic

import (
	"SendEmail-Service/pkg/rabbitmq"
	"log"
)

func Transactional(email []byte,template string)error{
	if email==nil{
		log.Println("email is empty")
		return nil
	}
	err:=rabbitmq.Publisher(email,template)
	return err
}