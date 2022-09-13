package main

import "fmt"

// main class as an interface
type IProduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

// New type computer
type Computer struct {
	name  string
	stock int
}

// Implement the interface to be considered an IProduct too
func (c *Computer) setStock(stock int) {
	c.stock = stock
}

func (c *Computer) setName(name string) {
	c.name = name
}

func (c *Computer) getStock() int {
	return c.stock
}

func (c *Computer) getName() string {
	return c.name
}

type Laptop struct {
	Computer
}

// Create Laptop based on Computer that Implements IProduct Interface
func newLaptop() IProduct {
	return &Laptop{
		Computer: Computer{
			name:  "Laptop Computer",
			stock: 25,
		},
	}
}

type Desktop struct {
	Computer
}

// Instanciate from particular type
func NewDesktop() IProduct {
	return &Desktop{
		Computer: Computer{
			name:  "Desktop Computer",
			stock: 35,
		},
	}
}

func GetComputerFactory(computerType string) (IProduct, error) {
	if computerType == "laptop" {
		return newLaptop(), nil
	}
	if computerType == "desktop" {
		return NewDesktop(), nil
	}

	return nil, fmt.Errorf("Invalid computer type")
}

// Test it
func printNameAndStock(p IProduct) {
	fmt.Printf("Product name: %s, with stock %d\n", p.getName(), p.getStock())
}

func main() {
	laptop, _ := GetComputerFactory("laptop")
	desktop, _ := GetComputerFactory("desktop")

	printNameAndStock(laptop)
	printNameAndStock(desktop)
}
