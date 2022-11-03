package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/hartsfield/gmailer"
)

var (
	urls []string
)

func getURLs() {
	file, err := os.Open("./config.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	getURLs()

	for range time.Tick(time.Minute * 5) {
		for _, url := range urls {
			r, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
			}

			if r.StatusCode != 200 {
				sendMail(url)
			}
		}

	}
}

func sendMail(url string) {
	msg := gmailer.Message{
		Recipient: "johnathanhartsfield@gmail.com",
		Subject:   "ALERT! " + url + " is down",
		Body:      url + " appears to be down",
	}
	msg.Send(onMessageSent)
}

func onMessageSent() {
	fmt.Println("sent mail")
}
