package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	ActorNum    = 100000
	LoopNum     = 1
	InteractNum = 100
)

func main() {
	actors := make([]*Actor, ActorNum)
	for i := range ActorNum {
		actors[i] = NewActor(&Player{HP: 100})
		actors[i].Start()
	}

	begin := time.Now()
	for range LoopNum {
		wg := sync.WaitGroup{}
		tmp := time.Now()
		for i := range actors {
			wg.Add(1)
			go func(actor *Actor) {
				defer wg.Done()
				for j := range InteractNum {
					// tmp := time.Now()
					target := actors[(i+j)%ActorNum]
					target.Attack()
					target.Heal()
					// if cost := time.Since(tmp); cost > 100*time.Millisecond {
					// 	fmt.Println("actor", i, "interact with actor", (i+j)%ActorNum, "cost:", cost)
					// }
				}
			}(actors[i])
		}
		wg.Wait()
		fmt.Println("once loop cost:", time.Since(tmp))
	}
	fmt.Println("all actor interactive cost:", time.Since(begin))

	for _, actor := range actors {
		actor.Stop()
	}
}
