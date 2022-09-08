package main

import (
	"bufio"
	"fmt"
	"os"
)

func userInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func chooseDifficulty() int {
	fmt.Printf("Укажите уровень сложности по трехбалльной шкале, где 1 - самый простой, 3 - самый сложный\nВвод: ")
	switch userInput() {
	case "1":
		fmt.Println("Вы выбрали самый низкий уровень сложности")
		return 1
	case "2":
		fmt.Println("Вы выбрали средний уровень сложности")
		return 2
	case "3":
		fmt.Println("Вы выбрали самый высокий уровень сложности")
		return 3
	case "q":
		fmt.Println("До свидания")
		return 0
	default:
		fmt.Println("Некорректный ввод. Попробуйте снова")
		return chooseDifficulty()
	}
}

func chooseMapSize() int {
	fmt.Printf("Укажите размер карты. 1 - маленькая, 2 - большая\nВвод: ")
	switch userInput() {
	case "1":
		fmt.Println("Вы выбрали малый размер карты")
		return 1
	case "2":
		fmt.Println("Вы выбрали большой размер карты")
		return 2
	case "q":
		fmt.Println("До свидания")
		return 0
	default:
		fmt.Println("Некорректный ввод. Попробуйте снова")
		return chooseMapSize()
	}
}

func setShipsAmount() int {
	fmt.Printf("Укажите число кораблей от 1 до 3\nВвод: ")
	switch userInput() {
	case "1":
		fmt.Println("Вы будете играть с одним кораблем")
		return 1
	case "2":
		fmt.Println("Вы будете играть с двумя кораблями")
		return 2
	case "3":
		fmt.Println("Вы будете играть с тремя кораблями")
		return 3
	case "q":
		fmt.Println("До свидания")
		return 0
	default:
		fmt.Println("Некорректный ввод. Попробуйте снова")
		return chooseMapSize()
	}
}

func setInitialField(size int) string {
	var field string
	for i := 0; i < size*5; i++ {
		if i != 0 {
			field += fmt.Sprintf("%d", i)
		} else {
			field += " "
			for j := 1; j < size*5; j++ {
				field += string(64 + j)
			}
			field += "\n"
			continue
		}
		for j := 1; j < size*5; j++ {
			field += " "
		}
		field += "\n"
	}

	return field
}

func main() {
	difficulty := chooseDifficulty()

	if difficulty == 0 {
		return
	}

	mapSize := chooseMapSize()

	if mapSize == 0 {
		return
	}

	shipsAmount := setShipsAmount()

	if shipsAmount == 0 {
		return
	}

	field := setInitialField(mapSize)

	fmt.Println(field)

	coordinatesAmount := shipsAmount * 3

	var playerCoordinatesSlice []string

	for i := 0; i < coordinatesAmount; i++ {
		playerCoordinatesSlice = append(playerCoordinatesSlice, "")
	}

	for i := 0; i < coordinatesAmount; i++ {
		coordinate := userInput()

		if coordinate == "q" {
			return
		} else {
			playerCoordinatesSlice[i] = coordinate
		}
	}

	fmt.Println("Успешно выбрана сложность, размер карты, число кораблей и координаты:", difficulty, mapSize, shipsAmount, playerCoordinatesSlice)
}
