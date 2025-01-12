package main

//go:generate go run . example.go
//go:generate gofmt -w .

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
