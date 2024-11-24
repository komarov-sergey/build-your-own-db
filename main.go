package main

import (
	"build-your-own-db/btree"
	"build-your-own-db/kvstore"
	"fmt"
)

func main() {
	btree.Example()
	kvstore.Example()

	fmt.Println("Проект успешно инициирован!")
}
