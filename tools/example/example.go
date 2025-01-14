//go:generate go run ../actorgen/ -input example.go -channel 256 -timeout 1000 -debug true

package main

func NewPlayer(id int) *Player {
	return &Player{
		Id: id,
		HP: 100,
	}
}

type Player struct {
	Id int
	HP int
}

func (p *Player) ID() int {
	return p.Id
}

func (p *Player) Attack() {
	p.HP -= 10
}

func (p *Player) Heal() bool {
	p.HP += 10
	return true
}
