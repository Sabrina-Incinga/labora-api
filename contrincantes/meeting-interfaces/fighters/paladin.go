package fighters

import (
	"math/rand"
)

type Paladin struct {
	BaseFighter
	InitialLife int
}

func NewPaladin(life int) *Paladin {
	paladin := &Paladin{
		BaseFighter: BaseFighter{
			Life: life,
		},
		InitialLife: life,
	}
	return paladin
}

func (p Paladin) ThrowAttack() int {
	intensityRatio := float64(p.Life)/float64(p.InitialLife)
	randNumber := rand.Intn(10)

	return int(float64(randNumber)*intensityRatio)
}

func (p *Paladin) ReceiveAttack(intensity int) {
	if p.Life <= intensity {
		p.Life = 0
		return
	} 
	p.Life -= intensity
}

func (p *Paladin) GetName() string {
	return "Paladin"
}

