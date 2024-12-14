package main

import "fmt"

func main() {
	var p, k, sh string

	fmt.Scan(&p, &k, &sh)
	if p == "Нет" && k == "Нет" && sh == "Нет" {
		fmt.Print("Кот")
	} else if p == "Нет" && k == "Нет" && sh == "Да" {
		fmt.Print("Жираф")
	} else if p == "Нет" && k == "Да" && sh == "Нет" {
		fmt.Print("Курица")
	} else if p == "Нет" && k == "Да" && sh == "Да" {
		fmt.Print("Страус")
	} else if p == "Да" && k == "Нет" && sh == "Нет" {
		fmt.Print("Дельфин")
	} else if p == "Да" && k == "Нет" && sh == "Да" {
		fmt.Print("Плезиозавры")
	} else if p == "Да" && k == "Да" && sh == "Нет" {
		fmt.Print("Пингвин")
	} else if p == "Да" && k == "Да" && sh == "Да" {
		fmt.Print("Утка")
	}
}
