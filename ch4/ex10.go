// Exercise 4.10:
// Modify issues to report the results in age categories, say less than a month old,
// less than a year old, and more than a year old.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	IssuesURL = "https://api.github.com/search/issues"
	Month     = 30 * 24
	Year      = 365 * 24
)

type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	fmt.Println(IssuesURL + "?q=" + q)
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssueSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	var withinMonth, withinYear, overYear []*Issue
	for _, issue := range result.Items {
		if t := time.Since(issue.CreatedAt).Hours(); t <= Month {
			withinMonth = append(withinMonth, issue)
			withinYear = append(withinYear, issue)
		} else if t <= Year {
			withinYear = append(withinYear, issue)
		} else {
			overYear = append(overYear, issue)
		}
	}

	fmt.Printf("%d issues less than a month:\n", len(withinMonth))
	for _, issue := range withinMonth {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			issue.Number, issue.User.Login, issue.Title)
	}

	fmt.Printf("%d issues less than a year:\n", len(withinYear))
	for _, issue := range withinYear {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			issue.Number, issue.User.Login, issue.Title)
	}

	fmt.Printf("%d issues more than a year:\n", len(overYear))
	for _, issue := range overYear {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			issue.Number, issue.User.Login, issue.Title)
	}
}
