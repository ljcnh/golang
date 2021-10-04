package prac10

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
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
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func SearchIssuesAll(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

// terms : 查询问题
// Year , Month, day : 距离当前时间所差的年，月，日
// Year,Month,day 要传入负值!!! 原本想改为默认传入负值的 但是没必要
// isIn 在日期之内,还是日期之外

func SearchIssuesLimited(terms []string, Year int, Month int, day int, isIn bool) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	var ans IssuesSearchResult
	now := time.Now()
	limitedTime := now.AddDate(Year, Month, day)
	//fmt.Printf("limitedTime: #%v \n", limitedTime)
	count := 0
	if isIn {
		for _, item := range result.Items {
			if limitedTime.Before(item.CreatedAt) {
				ans.Items = append(ans.Items, item)
				count++
			}
		}
	} else {
		for _, item := range result.Items {
			if limitedTime.After(item.CreatedAt) {
				ans.Items = append(ans.Items, item)
				count++
			}
		}
	}
	ans.TotalCount = count
	return &ans, nil
}
