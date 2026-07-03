package main

import "test/3-struct/storage"

const str = "asdasasd"

func main() {
	store := storage.NewStorage("data.json")
	store.Write([]byte(str))
	store.Read()
}
