package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	md "github.com/MichaelMure/go-term-markdown"
)

type Article struct {
	TypeOf                 string    `json:"type_of"`
	ID                     int       `json:"id"`
	Title                  string    `json:"title"`
	Description            string    `json:"description"`
	ReadablePublishDate    string    `json:"readable_publish_date"`
	Slug                   string    `json:"slug"`
	Path                   string    `json:"path"`
	URL                    string    `json:"url"`
	CommentsCount          int       `json:"comments_count"`
	PublicReactionsCount   int       `json:"public_reactions_count"`
	CollectionID           any       `json:"collection_id"`
	PublishedTimestamp     time.Time `json:"published_timestamp"`
	PositiveReactionsCount int       `json:"positive_reactions_count"`
	CoverImage             string    `json:"cover_image"`
	SocialImage            string    `json:"social_image"`
	CanonicalURL           string    `json:"canonical_url"`
	CreatedAt              time.Time `json:"created_at"`
	EditedAt               time.Time `json:"edited_at"`
	CrosspostedAt          any       `json:"crossposted_at"`
	PublishedAt            time.Time `json:"published_at"`
	LastCommentAt          time.Time `json:"last_comment_at"`
	ReadingTimeMinutes     int       `json:"reading_time_minutes"`
	TagList                string    `json:"tag_list"`
	Tags                   []string  `json:"tags"`
	BodyHTML               string    `json:"body_html"`
	MarkdownBody           string    `json:"body_markdown"`
	User                   struct {
		Name            string `json:"name"`
		Username        string `json:"username"`
		TwitterUsername string `json:"twitter_username"`
		GithubUsername  string `json:"github_username"`
		UserID          int    `json:"user_id"`
		WebsiteURL      string `json:"website_url"`
		ProfileImage    string `json:"profile_image"`
		ProfileImage90  string `json:"profile_image_90"`
	} `json:"user"`
}

type Comment struct {
	TypeOf    string    `json:"type_of"`
	IDCode    string    `json:"id_code"`
	CreatedAt time.Time `json:"created_at"`
	BodyHTML  string    `json:"body_html"`
	User      struct {
		Name            string `json:"name"`
		Username        string `json:"username"`
		TwitterUsername any    `json:"twitter_username"`
		GithubUsername  string `json:"github_username"`
		UserID          int    `json:"user_id"`
		WebsiteURL      any    `json:"website_url"`
		ProfileImage    string `json:"profile_image"`
		ProfileImage90  string `json:"profile_image_90"`
	} `json:"user"`
	Children []Comment `json:"children"`
}

func readArticle(articleName string) {
	res, err := http.Get("https://dev.to/api/articles/" + articleName)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	article := Article{}
	json.Unmarshal(body, &article)

	output := md.Render(article.MarkdownBody, 80, 6)

	fmt.Println("\033[1m" + article.Title + "\033[0m")
	fmt.Println(string(output))

	if includes(os.Args, "--show-comments") || includes(os.Args, "-sc") {
		commentsRes, err := http.Get(fmt.Sprintf("https://dev.to/api/comments?a_id=%d", article.ID))
		if err != nil {
			panic(err)
		}
		defer commentsRes.Body.Close()
		rawComments, err := io.ReadAll(commentsRes.Body)
		if err != nil {
			panic(err)
		}
		commentsList := make([]Comment, 1000)
		json.Unmarshal(rawComments, &commentsList)
		fmt.Printf("%d comments: \n", len(commentsList))
		for _, comment := range commentsList {
			fmt.Println("  " + comment.User.Name + ":")
			fmt.Println("    " + comment.BodyHTML)
			fmt.Println()
		}
	}

	fmt.Printf("\033[4;38;5;245mPublished on %s\n", article.ReadablePublishDate)
	fmt.Printf("See the original article here: \033[38;5;74m %s \033[0m\n", article.URL)
}

func writeArticle() {
	fmt.Println("\033[33;1mNot Working\033[0m.")
	fmt.Println("This function does not yet work. It is temporarily unavailable")
	os.Exit(0)
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
  "description": "%s",
  "series": null,
  "main_image": null,
  "canonical_url": null,
  "organization_id": null,
  "tags": "discuss, healthydebate"
}`, title, string(markdownBytes), published, description))

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
	fmt.Println(res)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(body)
}

func recentlyPosted() {
	res, err := http.Get("https://dev.to/api/articles/latest")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var articles []Article
	body, err := io.ReadAll(res.Body)
	json.Unmarshal(body, &articles)

	for _, article := range articles {
		fmt.Printf("%s \033[38;5;245m devcli read %s \033[0m\n", article.Title, article.Path[1:])
	}
}
