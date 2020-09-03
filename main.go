package main

import (
    "fmt"
		"math"
		"kw101/go-playground/api"
)

// A Month specifies a month of the year (January = 1, ...).
type Month int

const (
    January Month = 1 + iota
    February
    March
    April
    May
    November
    December
)

type geometry interface {
    area() float64
		perim() float64
		changeSize(width, height float64)
}

type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}

func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}
func (r *rect) changeSize(width, height float64) {
	r.width = width
	r.height = height
}
// func (r *rect) changeSize(width, height float64) {
// 	r.width = width
// 	r.height = height
// }

func tryTry () (int, int) {
	return 1, 3
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}
func (c *circle) changeSize(width, height float64) {
}

// meaure geometry and print size
func measure(g geometry) {
		g.changeSize(1, 2)
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func measureP(g geometry) {
	g.changeSize(1, 2)
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	model.Demo()
	model.Demo2()
	fmt.Println(February)
	a, b := tryTry ()
	fmt.Println(a)
	fmt.Println(b)
	// testgo.
    // r := rect{width: 3, height: 4}
		// c := circle{radius: 5}
		// rn := new(rect)
		
		// fmt.Println(r)
		// fmt.Println(c)
		// fmt.Println(rn)
    // measure(&r)
		// measure(&c)
		// measure(rn)
		// fmt.Println(r)
		// fmt.Println(c)
		// fmt.Println(rn)
}