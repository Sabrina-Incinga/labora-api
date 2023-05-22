package fighters

type Contender interface {
	ThrowAttack() int
	ReceiveAttack(intensity int)
	IsAlive() bool
	GetName() string
}
