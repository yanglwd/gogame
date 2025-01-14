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
	// ch := make(chan int, 100)
	// go func() {
	// 	close(ch)
	// }()
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
					fmt.Println("actor", i, "interact with actor", (i+j)%ActorNum, "cost:", time.Since(tmp))
				}
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
