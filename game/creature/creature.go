package creature

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

type Creature struct {
	hole int
	hp   int
	rep  int
	mass int
}

func New() *Creature {
	return &Creature{
		hole: 10,
		hp:   100,
		rep:  20,
		mass: 30,
	}
}

func (c *Creature) Eat() {
vozvrat:
	fmt.Println("Какой травки поедим,жухлой(Ж) или зелёной(З)?")
	fmt.Println("1. Жухлой")
	fmt.Println("2. Зелёной")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	grass := scanner.Text()
	if grass == "1" {
		c.hp += 10
		c.mass += 15
	} else if grass == "2" {
		if c.rep < 30 {
			c.hp -= 30
		} else if c.rep >= 30 {
			c.hp += 30
			c.mass += 30
		}
	} else {
		goto vozvrat
	}

}

func (c *Creature) Dig() {
vozvrat:
	fmt.Println("Как будем копать?")
	fmt.Println("1. Интенсивно")
	fmt.Println("2. Лениво")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	dig := scanner.Text()
	if dig == "1" {
		fmt.Println("Вы интенсивно копаете...")
		c.hole += 5
		c.hp -= 30
	} else if dig == "2" {
		fmt.Println("Вы лениво копаете...")
		c.hole += 5
		c.hp -= 30
	} else {
		goto vozvrat
	}

}

func (c *Creature) Fight() {
vozvrat:
	fmt.Println("С кем будем драться?")
	fmt.Println("1. Со слабым существом (вес 30)")
	fmt.Println("2. Со средним существом (вес 50)")
	fmt.Println("3. С сильным существом (вес 70)")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	enemy := scanner.Text()
	if enemy == "1" {
		c.randfight(30)
	}
	if enemy == "2" {

		c.randfight(50)
	}
	if enemy == "3" {
		c.randfight(70)
	} else {
		goto vozvrat
	}

}

func (c *Creature) randfight(enemy int) {
	if rand.Intn(c.mass+enemy) > c.mass {
		fmt.Println("Вы проиграли битву")
		c.hp -= enemy
	} else {
		fmt.Println("Вы выйграли битву")
		c.hp -= enemy / 3
	}
}

func (c *Creature) Night() {
	c.hole -= 2
	c.hp += 20
	c.rep -= 2
	c.mass -= 5
	//watch()
}

/*func (c *Creature) check() int {
	if c.hole == 0 || c.hp == 0 || c.rep == 0 || c.mass == 0 {

		return -1

	} else if c.rep >= 100 {

		return 1
	} else {
		return 0
	}
}*/
