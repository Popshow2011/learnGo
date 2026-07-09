package main

import (
	"fmt"
	"test/3-struct/api"
	"test/3-struct/bins"
	"test/3-struct/config"
	"test/3-struct/storage"
)

const str = "asdasasd"

func main() {

	config := config.NewConfig()
	var store storage.Storage = storage.NewStorage(config.BaseUrl, config.Key)
	apiService := api.NewApi(store)

	bin := *bins.NewBin("123", "MyBin", false)
	fmt.Println(bin)
	// err := apiService.CreateBin(bin)
	// fmt.Println(err, store)

	rBin, err := apiService.GetBin("6a4ff2f2da38895dfe478d5b")
	if err != nil {
		fmt.Println(err)
	}

	err = apiService.DeleteBin("6a4ff2f2da38895dfe478d5b")
	allBins, err := apiService.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rBin, allBins)

}
