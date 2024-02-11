package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
)

func readArticle(articleName string) {
	res, err := http.Get("https://dev.to/api/articles/" + articleName)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	scanner := bufio.NewScanner(res.Body)
	var articleApi string
	for scanner.Scan() {
		articleApi += scanner.Text() + "\n"
	}
	article := make(map[string]any)
	api := []byte(articleApi)
	json.Unmarshal(api, &article)
	fmt.Println(article["body_markdown"])
}
