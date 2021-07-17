package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	id              int
	leftCS, rightCS *ChopS
}

func (p Philo) eat(c chan *Philo, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		c <- &p
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("starting to eat", p.id+1)

		p.rightCS.Unlock()
		p.leftCS.Unlock()

		fmt.Println("finishing eating", p.id+1)
	}
	wg.Done()
}

func host(c chan *Philo, wg *sync.WaitGroup) {
	for {
		if len(c) == 2 {
			<-c
			<-c
			//time delay
			time.Sleep(20 * time.Millisecond)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	c := make(chan *Philo, 2)

	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i, CSticks[i], CSticks[(i+1)%5]}
	}

	go host(c, &wg)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat(c, &wg)
	}
	wg.Wait()
}
