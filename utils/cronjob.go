package utils

import (
	"log"
	"sync"
	"time"
)

func LogServiceStatus() {
	for {
		log.Println("service is running!")
		time.Sleep(30 * time.Second)
	}
}

func ProductMessageReceiver(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range ch {
		log.Println("Status:", msg)
	}
}
