package utils

import (
	"log"
	"time"
)

func LogServiceStatus() {
	for {
		log.Println("service is running!")
		time.Sleep(30 * time.Second)
	}
}
