package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"runtime"
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
		ducks *list.List
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
	//fmt.Println("aaaaa")
	for {

		for b.ducks.Len() > 0 {
			runtime.Gosched() //让出CPU
		}
		duck := &Duck{
			Number: rand.Intn(1000),
			Weight: rand.Float64(),
		}
		fmt.Println("生产。。。")
		b.ducks.PushBack(duck)
		fmt.Println(duck.Weight, duck.Number)
	}
}

func (p *ContosmerIpl) buyDuck(b *Bascket) {
	//fmt.Println("ccccccc")
	for {
		for b.ducks.Len() <= 0 {
			runtime.Gosched()
		}

		element := b.ducks.Back()
		fmt.Println("消费。。。")
		duck := element.Value.(*Duck)
		fmt.Println(duck.Weight, duck.Number)
		b.ducks.Remove(element)
	}

}

func main() {
	wg := new(sync.WaitGroup)
	p := new(ProducerImpl)
	c := new(ContosmerIpl)

	b := &Bascket{
		ducks: list.New(),
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
