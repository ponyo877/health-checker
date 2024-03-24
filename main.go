package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/ponyo877/health-checker/model"
)

var (
	targetEndpoint string
	notifyEndpoint string
	message        string
	grace          int
	interval       int
)

func main() {
	var err error
	targetEndpoint = os.Getenv("TARGET_ENDPOINT")
	notifyEndpoint = os.Getenv("NOTIFY_ENDPOINT")
	message = os.Getenv("MESSAGE")
	grace, err = strconv.Atoi(os.Getenv("GRACE"))
	if err != nil {
		log.Fatal(err)
	}
	interval, err = strconv.Atoi(os.Getenv("INTERVAL"))
	if err != nil {
		log.Fatal(err)
	}

	c := model.NewChecker(targetEndpoint, grace)
	h := model.NewHooker(notifyEndpoint, message)
	for {
		if err := c.Check(); err != nil {
			if err := h.Notify(); err != nil {
				log.Fatal(err)
			}
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
