package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStartUp(t *testing.T) {
	t.Run("get API key", func(t *testing.T) {
		_, ok := getAPIToken()
		if !ok {
			t.Errorf("Error API_TOKEN environment variable cannot be nill")
		}
	})

	t.Run("Create base URL", func(t *testing.T) {
		apiToken, _ := getAPIToken()

		got := createBaseUrl(apiToken)
		want := fmt.Sprintf("https://api.telegram.org/bot%s/", apiToken)

		if got != want {
			t.Errorf(" Got: %q want: %q", got, want)
		}
	})

	t.Run("Test connection to Telegram", func(t *testing.T) {
		apiToken, _ := getAPIToken()
		baseUrl := createBaseUrl(apiToken)
		requestUrl := baseUrl + "getMe"
		resp := makeRequest(requestUrl)

		gotStatus, wantStatus := resp.StatusCode, 200
		if gotStatus != wantStatus {
			t.Errorf("Response code failed got: %d, want: %d", gotStatus, wantStatus)
		}

		var got Ok
		err := json.NewDecoder(resp.Body).Decode(&got)
		if err != nil {
			t.Errorf("Got error unmarshalling the response: %q", err)
		}

		if !got.Ok {
			t.Errorf("There was an error making the request to Telegram")
		}

	})

	t.Run("Get Chat information", func(t *testing.T) {
		apiToken, _ := getAPIToken()
		baseUrl := createBaseUrl(apiToken)
		chatId := os.Getenv("CHAT_ID")
		resp := getChat(baseUrl, chatId)

		var got Chat
		err := json.NewDecoder(resp.Body).Decode(&got)
		if err != nil {
			t.Errorf("There was an error in retrieving the chat")
		}

		if !got.Ok {
			t.Errorf("There was an error making the request to Telegram")
		}
		//if got.Id != chatId {
		//	t.Errorf("How did we get the wrong chat???")
		//}
		fmt.Println(got)
	})

}
