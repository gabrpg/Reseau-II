package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

/**
Atelier 8
*/

func main() {
	rand.Seed(time.Now().UnixNano())
	num4()
}

/**
Numéro 1 -> Atomic
*/

func num1() {
	var compte uint32 = 0
	var multiple uint32 = 0
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		i := i
		go func(compte *uint32, multiple *uint32, wg *sync.WaitGroup) {
			defer wg.Done()
			if i%3 == 0 || i%5 == 0 {
				atomic.AddUint32(multiple, 1)
			}
			atomic.AddUint32(compte, 1)
		}(&compte, &multiple, &wg)
	}
	wg.Wait()
	fmt.Println("Compte:", compte, "Mutiple de 3 ou 5:", multiple)
}

/**
Numéro 2 -> Mutex
*/

func num2() {
	var m sync.Mutex
	var wg sync.WaitGroup

	wg.Add(10)
	go print(&wg, &m, random())
	go print(&wg, &m, random())
	go print(&wg, &m, random())
	go print(&wg, &m, random())
	go print(&wg, &m, random())
	go print(&wg, &m, random())
	go print(&wg, &m, random())
	go print(&wg, &m, random())
	go print(&wg, &m, random())
	go print(&wg, &m, random())
	wg.Wait()
}

func print(wg *sync.WaitGroup, m *sync.Mutex, nb int) {
	defer wg.Done()
	msg := "Goroutine #"
	m.Lock()
	for i := 0; i < len(msg); i++ {
		fmt.Print(string(msg[i]))
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Print(nb)
	m.Unlock()
	fmt.Println()
}

/**
Atelier 3 -> RWMutex
*/

func num3() {
	var wg sync.WaitGroup
	var m sync.RWMutex
	var nb = random()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read(&wg, &m, &nb)
		wg.Wait()
	}

	wg.Add(2)
	go write(&wg, &m, &nb)
	go write(&wg, &m, &nb)
	wg.Wait()
}

func read(wg *sync.WaitGroup, m *sync.RWMutex, nb *int) {
	defer wg.Done()
	m.RLock()
	time.Sleep(2 * time.Second)
	fmt.Println("Read:", *nb)
	m.RUnlock()
}

func write(wg *sync.WaitGroup, m *sync.RWMutex, nb *int) {
	defer wg.Done()
	m.Lock()
	*nb += 5
	time.Sleep(1 * time.Second)
	fmt.Println("Write:", *nb)
	m.Unlock()
}

func random() int {
	min := 1
	max := 10
	return rand.Intn(max-min+1) + min
}

/**
Atelier 4 -> Sémaphore
*/

func num4() {
	wg := SemaphoreWaitGroup{sem: make(chan bool, 4)}
	for i := 1; i <= 20; i++ {
		wg.Add(1)
		go toilet(i, &wg)
	}
	wg.Wait()
}

func toilet(id int, wg WaitGroup) {
	defer wg.Done()
	fmt.Println("Étudiant", id, "prend une toilette")
	time.Sleep(3 * time.Second)
	fmt.Println("Étudiant", id, "libère sa toilette")
}
