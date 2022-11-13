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

	cr := creature.New()
	var live bool = true
	for live {
	vozvrat:
		fmt.Println("Наступил новый день")
		cr.Param()
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
			cr.Dig()
		case "2":
			cr.Eat()
		case "3":
			cr.Fight()
		case "4":
			cr.Night()
		default:
			goto vozvrat
		}
		switch cr.Сheck() {
		case 0:
			{
				fmt.Println("Вы спите")
				cr.Night()
			}
		case 1:
			{
				fmt.Println("Вы выйграли!")
				live = false
			}
		case -1:
			{
				fmt.Println("ВЫ ПРОИГРАЛИ!")
				live = false
			}
		default:
			{
				fmt.Println("ВЫЛЕТ")
				live = false
			}
		}

	}
	fmt.Println("Конец игры")
}
