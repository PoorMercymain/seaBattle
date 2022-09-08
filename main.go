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

func initMemoryOfPlayerCoordinatesSlice(coordinatesAmount int) *[]string {
	var playerCoordinatesSlice []string

	for i := 0; i < coordinatesAmount; i++ {
		playerCoordinatesSlice = append(playerCoordinatesSlice, "")
	}

	return &playerCoordinatesSlice
}

func initCoordinatesAmount(shipsAmount int) int {
	return shipsAmount * 3
}

func isInt(char string) bool {
	return char >= "0" && char <= "9"
}

func checkCoordinateStringLength(coordinateString string) bool {
	return len(coordinateString) != 2
}

func isItTimeForExit(inputString string) bool {
	return inputString == "q"
}

func isSymbol(char string) bool {
	return (char >= "A" && char <= "Z") || (char >= "a" && char <= "z")
}

func areCoordinatesCorrect(coordinatesString string) bool {
	return checkCoordinateStringLength(coordinatesString) || !isInt(coordinatesString[0:1]) || !isSymbol(coordinatesString[1:2])
}

func toInt(someString string) (int, error) {
	return atoi(someString)
}

func insert(symbol string, someString string, index int) string {
	firstStringBuffer := someString[0:index]
	secondStringBuffer := someString[index+1 : len(someString)]
	return firstStringBuffer + symbol + secondStringBuffer
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

	coordinatesAmount := initCoordinatesAmount(shipsAmount)

	var playerCoordinatesSlice []string = *initMemoryOfPlayerCoordinatesSlice(coordinatesAmount)

	for i := 0; i < coordinatesAmount; i++ {
		coordinate := userInput()

		if areCoordinatesCorrect(coordinate) || coordinate == "q" {
			fmt.Println("Ошибка. Координаты введены в неверном формате.\n" +
				"Правильный формат: два символа, первый символ - цифра, второй - латинская буква.\n" +
				"Например: 2A")
			return
		} else {
			playerCoordinatesSlice[i] = coordinate
		}
		index, _ := toInt(playerCoordinatesSlice[i])
		field[index] = string(219)

	}

	fmt.Println("Успешно выбрана сложность, размер карты, число кораблей и координаты:", difficulty, mapSize, shipsAmount, playerCoordinatesSlice)
}
