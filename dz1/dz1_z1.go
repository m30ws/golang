package main

import (
	"errors"
	"fmt"
	"math"
)

var errNoData = errors.New("no data in list")

type ShopItem struct {
	Name		string
	Price		int
	Callory		int64
	Quantity	int
}

func main() {
	shopList := []ShopItem{
		{"Hosomaki Ginger", 54, 250, 1},
		{"Mlijecna cokolada", 8, 600, 3},
		{"Banana", 10, 150, 6},
	}

	tCal := totalCallory(shopList)
	fmt.Println(tCal) // 2950
	best, err := mostHealthy(shopList)
	fmt.Println(best) // [{Banana 10 150 6}]
	fmt.Println(err) // nil
}

func totalCallory(shoppingList []ShopItem) (total int) {
	for _, elem := range shoppingList {
		total += int(elem.Callory) * elem.Quantity
	}
	return
}

func mostHealthy(shoppingList []ShopItem) (items []ShopItem, err error) {
	if shoppingList == nil {
		err = errNoData
		return
	}

	minCal := int64(math.MaxInt64)
	for _, elem := range shoppingList {
		if elem.Callory < minCal {
			items = make([]ShopItem, 0, cap(shoppingList))
			minCal = elem.Callory
		}
		if elem.Callory == minCal {
			items = append(items, elem)
		}
	}

	return
}