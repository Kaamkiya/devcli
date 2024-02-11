package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"

	md "github.com/MichaelMure/go-term-markdown"
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
	article := make(map[string]interface{})
	json.Unmarshal([]byte(articleApi), &article)

	title := article["title"].(string)
	body := article["body_markdown"].(string)	
	output := md.Render(string(body), 80, 6)
	fmt.Println("\033[1m" + title + "\033[0m")
	fmt.Println(string(output))
}
