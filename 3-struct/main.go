package main

import (
	"fmt"
	"test/3-struct/api"
	"test/3-struct/bins"
	"test/3-struct/storage"
)

const str = "asdasasd"

func main() {
	var store storage.Storage = storage.NewStorage("data.json")

	apiService := api.NewApi(store)

	bin := bins.NewBin("123", "MyBin", false)
	apiService.SaveBin(bin)

	loadedBins, _ := apiService.LoadBins()
	fmt.Println(loadedBins)
}
