package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

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

	fmt.Println(tags)
}
