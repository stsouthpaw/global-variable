package worker

import (
	"fmt"
	"os"
	"time"

	"github.com/stsouthpaw/global-variable/pkg/global"
)

func StartWorker () {
	g := new(global.Global)
	c := make(chan string)

	go worker1(g)
	go worker2(g, c)
	go worker3(g, c)
	go worker4(g)
	
    for {
		state := <-c 
        if (state == "close") {
			fmt.Println("final")
			close(c)  

			os.Exit(200)
		}
	}

}

func worker1 (g *global.Global) {
	for {
		time.Sleep(15 * time.Second)
		count := g.GetCount()
		if (count == 0) {
			g.SetIncrement()
		}
	}
}

func worker2 (g *global.Global, c chan string) {
	for {
		time.Sleep(5 * time.Second)
		count := g.GetCount()
		if (count > 0) {
			g.SetIncrement()
			c <- "work"
		}
	}
}

func worker3 (g *global.Global, c chan string) {
	for {
		time.Sleep(10 * time.Second)
		count := g.GetCount()
		if (count > 10) {
			c <- "close"
		}
	}

}

func worker4 (g *global.Global) {
	for {
		time.Sleep(5 * time.Second)
		fmt.Println(g.GetCount())
	}
}