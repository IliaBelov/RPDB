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

func (c *Creature) Param() {
	fmt.Printf("HP:%v \n", c.hp)
	fmt.Printf("HOLE:%v\n", c.hole)
	fmt.Printf("REP:%v\n", c.rep)
	fmt.Printf("MASS:%v\n", c.mass)
}

func (c *Creature) Eat() {
vozvrat:
	fmt.Println("Какой травки поедим,жухлой(Ж) или зелёной(З)?")
	fmt.Println("1. Жухлой ")
	fmt.Println("2. Зелёной ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	grass := scanner.Text()
	switch grass {
	case "1":
		{
			c.hp += 10
			c.mass += 15
		}
	case "2":
		{
			if c.rep < 30 {
				c.hp -= 30
			} else {
				c.hp += 30
				c.mass += 30
			}
		}
	default:
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
	switch dig {
	case "1":
		{
			fmt.Println("Вы интенсивно копаете...")
			c.hole += 5
			c.hp -= 30
		}
	case "2":
		{
			fmt.Println("Вы лениво копаете...")
			c.hole += 5
			c.hp -= 30
		}
	default:
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
	switch enemy {
	case "1":
		{
			c.randfight(30)
		}
	case "2":
		{
			c.randfight(50)
		}
	case "3":
		{
			c.randfight(70)
		}
	default:
		goto vozvrat
	}

}

func (c *Creature) randfight(enemy int) {
	if rand.Intn(c.mass+enemy) > c.mass {
		fmt.Println("Вы проиграли битву")
		c.hp = c.hp - enemy/2
		c.rep = c.rep - ((c.mass + enemy) / 5)
	} else {
		fmt.Println("Вы выйграли битву")
		c.hp = c.hp - (enemy / 3)
		c.rep = c.rep + (c.mass+enemy)/4
	}
}

func (c *Creature) Night() {
	c.hole -= 2
	c.hp += 20
	c.rep -= 2
	c.mass -= 5
	//watch()
}

func (c *Creature) Сheck() int {
	if c.hole <= 0 || c.hp <= 0 || c.rep <= 0 || c.mass <= 0 {
		return -1
	} else if c.rep >= 100 {
		return 1
	} else {
		return 0
	}
}
