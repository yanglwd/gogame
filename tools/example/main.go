package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	ActorNum    = 1000
	LoopNum     = 100
	InteractNum = 10
)

func main() {
	// ch := make(chan int, 100)
	// close(ch)
	// <-ch
	// fmt.Println("ch closed")
	// time.Sleep(10 * time.Second)

	actors := make([]*Actor, ActorNum)
	for i := range ActorNum {
		actors[i] = NewActor(NewPlayer(i))
		actors[i].Start()
	}

	begin := time.Now()
	// for range LoopNum {
	wg := sync.WaitGroup{}
	for i := range actors {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range LoopNum {
				for j := range InteractNum {
					tmp := time.Now()
					target := actors[(i+j)%ActorNum]
					target.Attack()
					target.Heal()
					if cost := time.Since(tmp); cost > 100*time.Millisecond {
						fmt.Println("actor", i, "interact with actor", (i+j)%ActorNum, "cost:", cost)
					}
				}
				sec := rand.Int31n(10)
				time.Sleep(time.Duration(sec) * time.Second)
			}
		}()
	}
	wg.Wait()
	// }
	fmt.Println("all actor interactive cost:", time.Since(begin))

	for _, actor := range actors {
		actor.Stop()
	}
}
