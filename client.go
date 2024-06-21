package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
	}
	
	url := "http://localhost:8080/testing"

	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

	data := map[string]interface{}{
		"massage": "recived",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	buff := bytes.NewBuffer(jsonData)

	resp, err = http.Post(url, "testing/jason", buff)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
