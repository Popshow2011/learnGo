package main

import (
	"fmt"
	"test/3-struct/file"
)

func main() {
	file := file.NewConstructor("notebook.json")
	fmt.Println(file)
}
