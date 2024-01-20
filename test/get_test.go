package test

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/time/rate"
	"log"
	"net/http"
	"select_menu/enums"
	"select_menu/models"
	"strconv"
	"testing"
)

func TestGet(t *testing.T) {
	var url string
	rateLimiter := rate.NewLimiter(1, 3)
	params := "page/"
	for i := 0; i < 20; i++ {
		//生成url
		for {
			if rateLimiter.Allow() { // 桶满了阻塞在这里
				break
			}
		}
		url = "https://home.meishichina.com/recipe/tanggeng/"
		if i != 0 {
			url = url + params + strconv.Itoa(i+1)
		}
		//发送请求
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("request failed")
			return
		}
		defer resp.Body.Close()
		//提取需要内容
		if resp.StatusCode == 200 {
			doc, err2 := goquery.NewDocumentFromReader(resp.Body)
			if err2 != nil {
				fmt.Println("parse html failed")
				return
			}
			doc.Find(".space_left .detail").Each(func(i int, s *goquery.Selection) {
				title, _ := s.Find("a").Attr("title")
				material := s.Find(".subcontent").Text()
				link, _ := s.Find("a").Attr("href")
				//写入数据库
				data := &models.Food{
					Name:     title,
					Material: material,
					Url:      link,
					Status:   enums.FoodStatusSoup,
				}
				err3 := models.DB.Create(data).Error
				if err3 != nil {
					log.Printf("creat data error" + err.Error())
				}
				//fmt.Printf("菜名：%s   ,配料：%s   ,链接:%s \n", title, material, link)
			})
		} else {
			fmt.Println("response status code: ", resp.StatusCode)
		}
	}
}
func TestDelete(t *testing.T) {

	err := models.DB.Debug().Where("material=?", "NULL").Delete(new(models.Food)).Error
	if err != nil {
		log.Printf("Delete data Error:" + err.Error())

	}

}
