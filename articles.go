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
	publishDate := article["readable_publish_date"].(string)
	url := article["canonical_url"].(string)

	output := md.Render(string(body), 80, 6)
	
	fmt.Println("\033[1m" + title + "\033[0m")
	fmt.Println(string(output))
	fmt.Printf("\033[4;38;5;245mPublished on %s\n", publishDate)
	fmt.Printf("See the original article here: \033[38;5;74m%s\033[0m\n", url)
}
