package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type (
	Producer interface {
		makeDuck(b *Bascket)
	}

	Contosmer interface {
		buyDuck(b *Bascket)
	}

	Bascket struct {
		ducks chan Duck
	}

	ProducerImpl struct {
		Producer
	}

	ContosmerIpl struct {
		Contosmer
	}

	Duck struct {
		Number int
		Weight float64
	}
)

func (p *ProducerImpl) makeDuck(b *Bascket) {
	for {
		duck := Duck{
			Number: rand.Intn(1000),
			Weight: rand.Float64(),
		}

		b.ducks <- duck
		fmt.Println("生产。。。")
		fmt.Println(duck.Weight, duck.Number)
	}
}

func (c *ContosmerIpl) buyDuck(b *Bascket) {
	for {

		duck, _ := <-b.ducks
		fmt.Println("消费。。。")
		fmt.Println(duck.Weight, duck.Number)

	}
}

func main() {
	wg := new(sync.WaitGroup)
	p := new(ProducerImpl)
	c := new(ContosmerIpl)

	b := &Bascket{
		ducks: make(chan Duck),
	}
	//fmt.Println("main")
	wg.Add(2)

	go func() {
		p.makeDuck(b)
		wg.Done()
	}()

	go func() {
		c.buyDuck(b)
		wg.Done()
	}()
	wg.Wait()
}
