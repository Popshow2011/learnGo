package main

import (
	"flag"
	"test/3-struct/api"
	"test/3-struct/bins"
	"test/3-struct/config"
	"test/3-struct/storage"
)

func main() {

	config := config.NewConfig()
	var store storage.Storage = storage.NewStorage(config.BaseUrl, config.Key)
	apiService := api.NewApi(store)

	create := flag.Bool("create", false, "Метод для создания BIN")
	update := flag.Bool("update", false, "Метод для обновления BINа")
	delete := flag.Bool("delete", false, "Метод для удаления BINа")
	get := flag.Bool("get", false, "Метод для получения BINа")
	list := flag.Bool("list", false, "Метод для получения всех BIN")
	file := flag.String("file", "", "Введите имя файла")
	name := flag.String("name", "", "Введите имя BIN")
	id := flag.String("id", "", "Введите ID Bin")

	flag.Parse()

	switch {
	case *create:
		apiService.CreateBin(bin)
	case *update:
		apiService.PutBin(bin, id)
	case *delete:
		apiService.DeleteBin(id)
	case *get:
		apiService.GetBin(id)
	case *list:
		apiService.GetAll()

	}

	bin := bins.NewBin("123", "MyBin", false)

}
