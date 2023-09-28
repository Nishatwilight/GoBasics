// package main

// import "fmt"

// type ninjaStar struct {
// 	owner string
// }

// type ninjaSword struct {
// 	owner string
// 	level int
// }

// func (ns ninjaStar) attack() {
// 	fmt.Printf("Throwing Ninja Star %s\n", ns)
// }

// func (ninjasword ninjaSword) attack() {
// 	fmt.Printf("The owner of the sword is %s\n %s is at level of %d\n", ninjasword.owner, ninjasword.owner, ninjasword.level)
// }

//	func main() {
//		fmt.Println("Lets Learn Interface")
//		star := []ninjaStar{{owner: "Nisha"}, {owner: "King"}, {owner: "Queen"}}
//		for _, stars := range star {
//			// fmt.Printf("value of index is %d and the Value is %s\n", index, star)
//			stars.attack()
//		}
//		swords := []ninjaSword{{owner: "small N", level: 1}, {owner: "small K", level: 2}, {owner: "small A", level: 3}, {owner: "small S", level: 4}}
//		for _, sword := range swords {
//			sword.attack()
//		}
//	}
package main

import "fmt"

type attacker interface {
	attack()
}
type ninjaStar struct {
	Owner string `json:"owner"`
}
type ninjaSword struct {
	Owner string  `json:"sword_owner"`
	Level float32 `json:"sword_level"`
}
type ninjaWar struct {
	Opponent string `json:"opponent_name"`
	Level    int    `json:"opponent_level"`
}

func (ninStar ninjaStar) attack() {
	fmt.Printf("The owner of ninjaStar is %s\n", ninStar)
}
func (ninSword ninjaSword) attack() {
	fmt.Printf("The owner of ninjaSword is %s. His level is %f\n", ninSword.Owner, ninSword.Level)
}
func (ninWar ninjaWar) attack() {
	fmt.Printf("The opponent for  of ninjaStar is  %s and His level is %d\n", ninWar.Opponent, ninWar.Level)
}

func main() {
	fmt.Println("Lets Learn Actual Interface")
	attackers := []attacker{
		ninjaStar{Owner: "Nisha"},
		ninjaSword{Owner: "Someone Weak", Level: 8.86},
		ninjaWar{Opponent: "Someone Strong", Level: 9},
	}
	for _, attacker := range attackers {
		attacker.attack()
	}
}
