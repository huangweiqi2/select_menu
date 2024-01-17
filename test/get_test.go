package test

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"testing"
)

func TestGet(t *testing.T) {
	var url string
	params := "page/"
	for i := 0; i < 10; i++ {
		url = "https://home.meishichina.com/recipe/recai/"
		if i != 0 {
			url = url + params + strconv.Itoa(i+1)
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("request failed")
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode == 200 {
			doc, err2 := goquery.NewDocumentFromReader(resp.Body)
			if err2 != nil {
				fmt.Println("parse html failed")
				return
			}
			doc.Find(".space_left .detail").Each(func(i int, s *goquery.Selection) {
				title, _ := s.Find("a").Attr("title")
				subcontent := s.Find(".subcontent").Text()
				link, _ := s.Find("a").Attr("href")
				fmt.Printf("菜名：%s   ,配料：%s   ,链接:%s \n", title, subcontent, link)
			})
		} else {
			fmt.Println("response status code: ", resp.StatusCode)
		}
	}

}
