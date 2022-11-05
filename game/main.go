package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var (
	hole int = 10
	hp   int = 100
	rep  int = 20
	mass int = 30
)

func main() {
	/*fmt.Println("Enter command:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	fmt.Println("You entered:", input)*/
	watch()
	day()
}

func fight() {
vozvrat:
	fmt.Println("С кем будем драться?")
	fmt.Println("1. Со слабым существом (вес 30)")
	fmt.Println("2. Со средним существом (вес 50)")
	fmt.Println("3. С сильным существом (вес 70)")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	enemy := scanner.Text()
	if enemy == "1" {
		randfight(30)
	}
	if enemy == "2" {
		randfight(50)
	}
	if enemy == "3" {
		randfight(70)
	} else {
		goto vozvrat
	}

}

func randfight(enemy int) {
	if rand.Intn(mass+enemy) > mass {
		fmt.Println("Вы проиграли битву")
		hp -= enemy
	} else {
		fmt.Println("Вы выйграли битву")
		hp -= enemy / 3
	}
}
func dig() {
vozvrat:
	fmt.Println("Как будем копать?")
	fmt.Println("1. Интенсивно")
	fmt.Println("2. Лениво")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	dig := scanner.Text()
	if dig == "1" {
		fmt.Println("Вы интенсивно копаете...")
		hole = hole + 5
		hp = hp - 30
	} else if dig == "2" {
		fmt.Println("Вы лениво копаете...")
		hole = hole + 5
		hp = hp - 30
	} else {
		goto vozvrat
	}

}
func eat() {
vozvrat:
	fmt.Println("Какой травки поедим,жухлой(Ж) или зелёной(З)?")
	fmt.Println("1. Жухлой")
	fmt.Println("2. Зелёной")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	grass := scanner.Text()
	if grass == "1" {
		hp = hp + 10
		mass = mass + 15
	} else if grass == "2" {
		if rep < 30 {
			hp = hp - 30
		} else if rep >= 30 {
			hp = hp + 30
			mass = mass + 30
		}
	} else {
		goto vozvrat
	}

}
func watch() {
	fmt.Println("Д:", hole)
	fmt.Println("З:", hp)
	fmt.Println("У:", rep)
	fmt.Println("В:", mass)
}
func night() {
	hole = hole - 2
	hp = hp + 20
	rep = rep - 2
	mass = mass - 5
	watch()
}
func check() int {
	if hole == 0 || hp == 0 || rep == 0 || mass == 0 {

		return -1

	} else if rep >= 100 {

		return 1
	} else {
		return 0
	}
}
func day() {
new:
	fmt.Println("Наступил новый день.")
	watch()
vozvrat:
	fmt.Println("Что будем делать?")
	fmt.Println("1. Копать нору")
	fmt.Println("2. Поесть травки")
	fmt.Println("3. Пойти подраться")
	fmt.Println("4. Поcпать")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	switch input {
	case "1":
		dig()
	case "2":
		eat()
	case "3":
		fight()
	case "4":
		night()
	default:
		goto vozvrat
	}
	switch check() {
	case 0:
		{
			night()
			goto new
		}
	case 1:
		fmt.Println("Вы выйграли!")
	case -1:
		fmt.Println("ВЫ ПРОИГРАЛИ!")
	default:
		fmt.Println("ВЫЛЕТ")
	}

}
