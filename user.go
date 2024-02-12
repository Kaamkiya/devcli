package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

/* USER SPECIFIED WITH API KEY */
func followingTags() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://dev.to/api/follows/tags", nil)
	req.Header.Set("api-key", os.Getenv("DEV_API_KEY"))
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	
	var tags string
	scanner := bufio.NewScanner(res.Body)
	for scanner.Scan() {
		tags += scanner.Text() + "\n"
	}
	
	tagsList := make([]map[string]string, 100)
	json.Unmarshal([]byte(tags), &tagsList)

	for _, tag := range tagsList {
		fmt.Println(tag["name"])
	}
}

func followers() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://dev.to/api/followers/users", nil)
	req.Header.Set("api-key", os.Getenv("DEV_API_KEY"))
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var rawFollowers string
	scanner := bufio.NewScanner(res.Body)
	for scanner.Scan() {
		rawFollowers += scanner.Text() + "\n"
	}

	// 1000 is the max amount of users we can request
	followers := make([]map[string]string, 1000)
	json.Unmarshal([]byte(rawFollowers), &followers)
	for _, follower := range followers {
		fmt.Println(follower["name"])
	}
}

/* ANY USER */
func articles() {}
