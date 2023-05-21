package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/vituchon/labora-golang-course/meeting-interfaces/fighters"
)

func main() {
	var police fighters.Police = fighters.Police{
		BaseFighter: fighters.BaseFighter{
			Life: 10,
		},
		Armour: 5,
	}
	var criminal fighters.Criminal = fighters.Criminal{
		BaseFighter: fighters.BaseFighter{
			Life: 10,
		},
	}
	var paladin fighters.Paladin = *fighters.NewPaladin(20)
	var contenders []fighters.Contender = make([]fighters.Contender, 3)

	randomValueBetweenOneAndZero := rand.Intn(3)
	contenders[randomValueBetweenOneAndZero] = &police
	contenders[(randomValueBetweenOneAndZero+1)%3] = &criminal
	contenders[(randomValueBetweenOneAndZero+2)%3] = &paladin

	var areAllAlive = police.IsAlive() && criminal.IsAlive() && paladin.IsAlive()
	for areAllAlive {
		firstContenderPosition := rand.Intn(3)
		intensity := contenders[firstContenderPosition].ThrowAttack()
		
		attackReceiverPosition := rand.Intn(3)
		for attackReceiverPosition == firstContenderPosition {
			attackReceiverPosition = rand.Intn(3)
		}
		contenders[attackReceiverPosition].ReceiveAttack(intensity)
		fmt.Println(contenders[firstContenderPosition].GetName(), " tira golpe con intensidad =", intensity, "al contendiente: ", contenders[attackReceiverPosition].GetName())

		if contenders[attackReceiverPosition].IsAlive() {
			intensity := contenders[attackReceiverPosition].ThrowAttack()
			fmt.Println(contenders[attackReceiverPosition].GetName(), " tira golpe con intensidad =", intensity, "al contendiente: ", contenders[firstContenderPosition].GetName())
			contenders[firstContenderPosition].ReceiveAttack(intensity)
		}

		fmt.Printf("PoliceLife=%d, CriminalLife=%d, PaladinLife=%d\n", police.Life, criminal.Life, paladin.Life)
		areAllAlive = police.IsAlive() && criminal.IsAlive() && paladin.IsAlive()
		time.Sleep(3 * time.Second)
	}
}

func main_legacy() {

	var police fighters.Police = fighters.Police{
		BaseFighter: fighters.BaseFighter{
			Life: 10,
		},
		Armour: 5,
	}
	var criminal fighters.Criminal = fighters.Criminal{
		BaseFighter: fighters.BaseFighter{
			Life: 10,
		},
	}

	randomValueBetweenOneAndZero := rand.Intn(2)
	policeHitFirst := randomValueBetweenOneAndZero == 1

	var areBothAlive = police.IsAlive() && criminal.IsAlive()
	for areBothAlive {
		if policeHitFirst {
			intesity := police.ThrowAttack()
			fmt.Println("Policia tira golpe con intensidad =", intesity)
			criminal.ReceiveAttack(intesity)

			if criminal.IsAlive() {
				intesity := criminal.ThrowAttack()
				fmt.Println("Criminal tira golpe con intensidad =", intesity)
				police.ReceiveAttack(intesity)
			}
		} else {
			intesity := criminal.ThrowAttack()
			fmt.Println("Criminal tira golpe con intensidad =", intesity)
			police.ReceiveAttack(intesity)

			if police.IsAlive() {
				intesity := police.ThrowAttack()
				fmt.Println("Policia tira golpe con intensidad =", intesity)
				criminal.ReceiveAttack(intesity)
			}
		}
		fmt.Printf("PoliceLife=%d, CriminalLife=%d\n", police.Life, criminal.Life)
		areBothAlive = police.IsAlive() && criminal.IsAlive()
		time.Sleep(3 * time.Second)
	}

}
