package main

import (
	"bufio"
	"fmt"
	"game/creature"
	"os"
)

func main() {
	/*fmt.Println("Enter command:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	fmt.Println("You entered:", input)*/
	//watch()
	hfme := creature.New()
	hfme.Eat()
	//day()
}

/*func watch() {
	fmt.Println("Д:", hole)
	fmt.Println("З:", hp)
	fmt.Println("У:", rep)
	fmt.Println("В:", mass)
}*/

func Day() {
	//new:
	fmt.Println("Наступил новый день.")
	//watch()
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
		//dig()
	case "2":
		//eat()
	case "3":
		//fight()
	case "4":
		//night()
	default:
		goto vozvrat
	}
	/*switch check() {
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
	}*/

}
