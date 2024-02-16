package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math"
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

	tags, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	tagsList := make([]map[string]string, 1000)
	json.Unmarshal(tags, &tagsList)

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

	rawFollowers, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	
	// 1000 is the max amount of users we can request
	followers := make([]map[string]string, 1000)
	json.Unmarshal(rawFollowers, &followers)
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

	rawList, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	list := make([]map[string]interface{}, 1000)
	json.Unmarshal(rawList, &list)

	for _, e := range list {
		article := e["article"].(map[string]interface{})
		title := article["title"].(string)
		id := article["id"].(float64)
		fmt.Printf("%s - \033[38;5;245mdevcli read %d \033[0m\n", title, math.Round(id))
	}
}

func articles() {}
