package menu

import (
	"fmt"
	"math/rand"
	"select_menu/models"
	"strconv"
	"sync"
	"testing"
)

func TestName(t *testing.T) {
	var foods []models.Food
	for i := 0; i < 100; i++ {
		foods = append(foods, models.Food{Name: strconv.Itoa(i)})
	}

	foodChan := make(chan models.Food)
	var wg sync.WaitGroup

	wg.Add(1)
	go func(wg *sync.WaitGroup, foodChan chan<- models.Food, foods []models.Food) {
		defer func() {
			fmt.Println("close foodChan")
			close(foodChan)
			wg.Done()
		}()
		for _, food := range foods {
			foodChan <- food
		}
	}(&wg, foodChan, foods)

	f := func(wg *sync.WaitGroup, foodChan <-chan models.Food) (foodRateChan chan FoodRate) {
		foodRateChan = make(chan FoodRate)
		wg.Add(1)
		go func(wg *sync.WaitGroup, foodChan <-chan models.Food, foodRateChan chan<- FoodRate) {

			defer func() {
				fmt.Println("close foodRateChan")
				close(foodRateChan)
				wg.Done()
			}()
			for food := range foodChan {
				fmt.Println("duqu foodChan:", food.Name)
				foodRateChan <- FoodRate{
					Rate: rand.Intn(10),
					Food: food,
				}
			}
		}(wg, foodChan, foodRateChan)
		return foodRateChan
	}

	chans := make([]<-chan FoodRate, 0, 3)
	for i := 0; i < 3; i++ {
		chans = append(chans, f(&wg, foodChan))
	}
	for i, rates := range chans {
		for rate := range rates {
			fmt.Println("i:", i, rate.Food.Name)
		}
	}

	wg.Wait()

}
