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
	
	tagsList := make([]map[string]string, 1000)
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
func displayUser(username string) {
	res, err := http.Get("https://dev.to/api/users/by_username?url=" + username)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	scanner := bufio.NewScanner(res.Body)
	
	var userApi string
	for scanner.Scan() {
		userApi += scanner.Text()
	}
	user := make(map[string]string)
	json.Unmarshal([]byte(userApi), &user)
	fmt.Println("\033[38;5;245m   Username:              \033[0m" + user["username"])
	fmt.Println("\033[38;5;245m   Name:                  \033[0m" + user["name"])
	fmt.Println("\033[38;5;245m   Joined:                \033[0m" + user["joined_at"])
	fmt.Println("\033[38;5;245m   Bio:                   \033[0m" + user["summary"])

	if user["github_username"] != "" {
		fmt.Println("\033[38;5;245m   Github Profile:        \033[0mhttps://github.com/" + user["github_username"])
	}
	if user["twitter_username"] != "" {
		fmt.Println("\033[38;5;245m   Twitter Username:      \033[0m" + user["twitter_username"])
	}	
}

func readingList() {
	client := http.Client{}
	req, _ := http.NewRequest("GET", "https://dev.to/api/readinglist", nil)
	req.Header.Set("api-key", os.Getenv("DEV_API_KEY"))
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	scanner := bufio.NewScanner(res.Body)
	
	var rawList string
	for scanner.Scan() {
		rawList += scanner.Text()
	}

	list := make([]map[string]interface{}, 1000)
	json.Unmarshal([]byte(rawList), &list)

	for _, e := range list {
		article := e["article"].(map[string]interface{})
		title := article["title"].(string)
		path := article["path"].(string)
		fmt.Println(title + " - \033[38;5;245mdevcli read " + path[1:] + "\033[0m")
	}
}

func articles() {}
