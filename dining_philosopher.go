package main

import (
	"fmt"
	"sync"
	"time"
)

// Dining philosopher, a famous concurrency problem

type Philosopher struct {
	id          int
	leftFork    *sync.Mutex
	rightFork   *sync.Mutex
	eatingState *sync.Mutex
}

func (p Philosopher) think() {
	fmt.Printf("Philosopher %d is thinking\n", p.id)
	time.Sleep(time.Millisecond * time.Duration(p.id*100))
}

func (p Philosopher) eat() {
	fmt.Printf("Philosopher %d is eating\n", p.id)
	time.Sleep(time.Millisecond * time.Duration(p.id*100))
}

func (p Philosopher) dine(host *sync.WaitGroup, semaphore chan struct{}) {
	defer host.Done()

	for i := 0; i < 3; i++ {
		p.think()

		semaphore <- struct{}{}

		p.leftFork.Lock()
		p.rightFork.Lock()

		p.eat()

		p.rightFork.Unlock()
		p.leftFork.Unlock()

		<-semaphore
	}
}

func main() {
	const numPhilosophers = 5
	const maxEatingPhilosophers = 2

	var forks []*sync.Mutex
	for i := 0; i < numPhilosophers; i++ {
		forks = append(forks, &sync.Mutex{})
	}

	semaphore := make(chan struct{}, maxEatingPhilosophers)

	var philosophers []Philosopher
	for i := 0; i < numPhilosophers; i++ {
		philosopher := Philosopher{
			id:          i + 1,
			leftFork:    forks[i],
			rightFork:   forks[(i+1)%numPhilosophers],
			eatingState: &sync.Mutex{},
		}
		philosophers = append(philosophers, philosopher)
	}

	var wg sync.WaitGroup
	for _, philosopher := range philosophers {
		wg.Add(1)
		go philosopher.dine(&wg, semaphore)
	}

	wg.Wait()

	close(semaphore)

	fmt.Println("Dining completed!")
}
