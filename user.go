package main

import (
	"fmt"
	"net/http"
)

func followingTags() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://dev.to/api/follows/tags", nil)
	req.Header.Set("api-key", os.Getenv("DEV_API_KEY"))
	res, _ := client.Do(req)
}
