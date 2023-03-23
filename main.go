package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func getUserID(username string, BearerToken string) string {
	req, err := http.NewRequest("GET", "https://api.twitter.com/2/users/by/username/"+username, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+BearerToken)

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	var jmap map[string]interface{}
	err = json.Unmarshal(body, &jmap)
	if err != nil {
		log.Fatal(err)
	}

	id := jmap["data"].(map[string]interface{})["id"].(string)
	return id
}

func main() {
	token, err := os.ReadFile("BearerToken.txt")

	if err != nil {
		log.Fatal(err)
	}

	BearerToken := string(token)
	var username string
	fmt.Printf("[+] Username: ")
	fmt.Scanln(&username)
	fmt.Printf("%v", getUserID(username, BearerToken))
}
