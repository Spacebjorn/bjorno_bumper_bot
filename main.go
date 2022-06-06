package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Ok struct {
	Ok bool
}

type User struct {
	Id        int
	FirstName string
	LastName  string
	UserName  string
}

type Message struct {
	MessageId int
	From      User
}

type Chat struct {
	Ok            bool
	Id            int
	Type          string
	Title         string
	PinnedMessage Message
}

func getChat(baseUrl string, chatId string) (resp *http.Response) {
	requestUrl := baseUrl + "getChat"
	resp = makeRequest(requestUrl)
	return resp
}
func makeRequest(requestUrl string) (resp *http.Response) {
	resp, err := http.Get(requestUrl)
	if err != nil {
		fmt.Printf("ERROR ERROR ERROR PANIC")
		fmt.Printf("%q", err)
		os.Exit(1)
	}
	return resp
}

func getBot(baseUrl string) {
	requestUrl := baseUrl + "getMe"
	makeRequest(requestUrl)
}

func getAPIToken() (string, bool) {
	apiToken, ok := os.LookupEnv("API_TOKEN")
	return apiToken, ok
}

func createBaseUrl(apiToken string) string {
	return fmt.Sprintf("https://api.telegram.org/bot%s/", apiToken)
}

func main() {
	fmt.Println("Starting Telegram bot...")
	apiToken, ok := getAPIToken()
	if !ok {
		log.Fatal("API_TOKEN environment variable not set!")
	}
	baseUrl := createBaseUrl(apiToken)
	getBot(baseUrl)
}
