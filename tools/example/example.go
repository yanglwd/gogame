//go:generate go run ../actorgen/ -input example.go -channel 256 -timeout 1000

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
