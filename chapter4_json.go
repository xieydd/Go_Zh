package main

import (
	"encoding/json"
	"fmt"
	//"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"text/template"
	"time"
)

const IssueURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number   int
	HTMLURL  string `json:"html_url"`
	Title    string
	State    string
	User     *User
	CreateAt time.Time `json:"created_at"`
	Body     string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

const templ = `{{.TotalCount}} issues:
{{range .Items}}--------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age: {{.CreateAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

//var report = template.Must(template.New("issuelist").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ))

var issueList = template.Must(template.New("issuelist1").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
<th>#</th>th>
<th>State</th>th>
<th>User</th>th>
<th>Title</th>th>
</tr>tr>
{{range .Items}}
<tr>
<td><a href='{{.HTMLURL}}'>{{.Number}}</a>a></td>td>
<td>{{.State}}</td>td>
<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a>a></td>td>
<td><a href='{{.HTMLURL}}'>{{.Title}}</a>a></td>td>
</tr>tr>
{{end}}
</table>table>
`))

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s       %s\n", item.Number, item.User.Login, item.Title, item.CreateAt.UTC().Format(time.UnixDate))
	}

	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssueURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
