package main

/*
	Abstract Factory is a creational design pattern that lets you produce families of related objects without specifying their concrete classes.
*/

/* How to Implement

Map out a matrix of distinct product types versus variants of these products.

Declare abstract product interfaces for all product types. Then make all concrete product classes implement these interfaces.

Declare the abstract factory interface with a set of creation methods for all abstract products.

Implement a set of concrete factory classes, one for each product variant.

Create factory initialization code somewhere in the app. It should instantiate one of the concrete factory classes, depending on the application configuration or the current environment. Pass this factory object to all classes that construct products.

Scan through the code and find all direct calls to product constructors. Replace them with calls to the appropriate creation method on the factory object.
*/

import (
	abstractfactory "abstract_factory_method/abstract_factory"
	abstractproduct "abstract_factory_method/abstract_product"
	"fmt"
)

func main() {
	adidasFactory, _ := abstractfactory.GetSportsFactory("adidas")
	nikeFactory, _ := abstractfactory.GetSportsFactory("nike")

	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}

func printShoeDetails(s abstractproduct.IShoe) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}

func printShirtDetails(s abstractproduct.IShirt) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}
