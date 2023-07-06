package modules

import "fmt"

type Shape interface {
	area() float32
	perimeter() float32
	radius(arg int) float32
}

type Rect struct {
	length, breadth float32
}

type Circle struct {
	circular float32
}

func (r Rect) area() float32 {
	return r.breadth * r.length
}

func (r Rect) perimeter() float32 {
	return r.breadth * r.length * 0.5
}

func (r Rect) radius(num int) float32 {
	return float32(num)
}

func calculate(s Shape) {
	fmt.Println("Shape:", s)
	fmt.Println("Area:", s.area())
	fmt.Println("Perimeter:", s.perimeter())
	fmt.Println("Radius:", s.radius(2))
}

func Interface() {
	rect := Rect{20, 20}
	calculate(rect)

	_toyota := CarStruct{"Toyota", 2000}
	toyota(_toyota)
}

type Car interface {
	brand() string
	year() int
}

type CarStruct struct {
	_brand string
	_year  int
}

func toyota(c Car) {
	b := c.brand()
	y := c.year()

	fmt.Println("Brand is:", b, "\nYear:", y)
}

func (c CarStruct) brand() string {
	return c._brand
}

func (c CarStruct) year() int {
	return c._year
}
