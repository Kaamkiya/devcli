package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
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

func writeArticle() {
	scanner := bufio.NewScanner(os.Stdin)

	var title string
	var path string
	var published bool
	var description string

	fmt.Print("What's the title of your article? ")
	scanner.Scan()
	title = scanner.Text()

	fmt.Print("Where is the file that your article is currently in? (eg ./article.md, ../learn-c.markdown)")
	scanner.Scan()
	path = scanner.Text()

	fmt.Print("Would you like to publish your article now? (Y/n)")
	scanner.Scan()
	input := scanner.Text()[0]
	if input == 'y' || input == 'Y' {
		published = true
	} else {
		published = false
	}

	fmt.Print("Enter a description for your article: ")
	scanner.Scan()
	description = scanner.Text()

	markdownBytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Could not open file \"" + path + "\"")
		panic(err)
	}

	article := []byte(fmt.Sprintf(`{
  "title": "%s",
  "body_markdown": "%s",
  "published": %t,
  "description": "%s"
}`, title, markdownBytes, published, description))

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://dev.to/api/articles/", bytes.NewBuffer(article))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", os.Getenv("DEV_API_KEY"))
	if err != nil {
		panic(err)
	}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
}
