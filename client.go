package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	menu()
}

func menu() {
	client := http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
	}

	url := "http://172.31.85.85:80/testing"

	var choose int = 0
	for {
		fmt.Println("Session 8")
		fmt.Println("1. Receive data (get method)")
		fmt.Println("2. Send data (post method)")
		fmt.Println(">> ")
		fmt.Scanf("%d", &choose)
		switch choose {
		case 1:
			receiveMsg(client, url)
		case 2:
			sendMsg(client, url)
		}
	}
}

func msgExtractor(r *http.Response) (string, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil

}

func receiveMsg(client http.Client, url string) {
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	msg, err := msgExtractor(resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
	return
}

func sendMsg(client http.Client, url string) {
	data := map[string]interface{}{
		"massage": "recived",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	buff := bytes.NewBuffer(jsonData)

	resp, err := http.Post(url, "testing/jason", buff)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
