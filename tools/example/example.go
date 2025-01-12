//go:generate go run ../actorgen/ example.go

package main

type Player struct {
	HP int
}

func (p *Player) Attack() {
	p.HP -= 10
}

func (p *Player) Heal() bool {
	p.HP += 10
	return true
}
