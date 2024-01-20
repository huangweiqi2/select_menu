package test

import (
	"context"
	"fmt"
	"github.com/k0kubun/pp/v3"
	"log"
	"select_menu/models"
	"sync"
	"testing"
	"time"
)

func TestFind(t *testing.T) {
	data := make([]models.Food, 0)
	err := models.DB.Model(new(models.Food)).Where("status=?", 1).Find(&data).Error
	if err != nil {
		log.Printf("query data error:" + err.Error())
	}

	fmt.Println(data)
	//ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*20)
	for _, food := range data {
		pp.Println("food:", food)
		//fmt.Printf("sdf:%v\n", food)
		//name := food.Name
		//models.RDB.SAdd(ctx, "foodName", name)
		//fmt.Printf("第%d道菜的菜名,%s\n", i, name)
	}

}

func TestName1(t *testing.T) {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.TODO())
	//notifyChan := make(chan struct{})
	wg.Add(2)
	go sdfdsfds(ctx, &wg)
	go sdfdsfds(ctx, &wg)
	time.Sleep(time.Second * 5)

	cancel()
	//go sdfdsfds(ctx, &wg)
	wg.Wait()

}

func sdfdsfds(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("外部取消")
			return
		default:
			fmt.Println("sleep")

			time.Sleep(time.Second)
		}

	}

}

func TestSaveFood(t *testing.T) {
	food := models.Food{
		Name:     "dsfsd",
		Material: "sdfsgdsfsfs",
		Url:      "",
	}
	models.DB.Save(&food)

}
