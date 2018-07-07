package practice

// This practice file is implemented following the article
// "Why Go’s structs are superior to class-based inheritance" on medium.com
//https://medium.com/@simplyianm/why-gos-structs-are-superior-to-class-based-inheritance-b661ba897c67

import "fmt"

// Animal is an interface for all animal implementations
type Animal interface {
	Name() string
}

// PartyAnimal is a composed interface extending the Animal interface
type PartyAnimal interface {
	Animal // extending an existing interface
	Party()
}

//******************************************************************************
//* Animal implementation "Dog"
//******************************************************************************

// Dog is an Animal implementation
type Dog struct{}

// Bark makes the dog say something
func (d *Dog) Bark() string {
	return "Woof!"
}

// Name returns the name of the animal, in our case "Dog"
func (d *Dog) Name() string {
	return "Dog"
}

//******************************************************************************
//* Animal implementation "Cat"
//******************************************************************************

// Cat is an Animal implementation
type Cat struct{}

// Speak makes the cat say something
func (c *Cat) Speak() string {
	return "Miao!"
}

// Name returns the name of the animal, in our case "Dog"
func (c *Cat) Name() string {
	return "Cat"
}

//******************************************************************************
//* PartyAnimal implementation "Raver"
//******************************************************************************

// Raver is a PartyAnimal implementation
type Raver struct{}

// Name returns the name of the PartyAnimal, in our case "Raver"
func (r *Raver) Name() string {
	return "Raver"
}

// Party makes the PartyAnimal go wild
func (r *Raver) Party() {
	fmt.Println("emtz..emtz..emtz!!")
}

//******************************************************************************
//* Specialiced version of Dog
//******************************************************************************

// GuideDog is a special kind of Dog
type GuideDog struct {
	*Dog // embedding a struct
}

// Help will give instructions in what way the GuideDog helps
func (gd *GuideDog) Help() {
	fmt.Printf("I'm a %s and I will guide you\n", gd.Name())
}

//******************************************************************************
//* Specialized version of Cat and Dog
//******************************************************************************

// CatDog is a special kind of Cat and Dog
type CatDog struct {
	*Cat // embedding struct Cat
	*Dog // embedding struct Dog
}

// Name returns the name of the animal, in our case "CatDog"
func (cd *CatDog) Name() string {
	// referencing the Name() methods of both embedded structures
	return fmt.Sprintf("%s%s", cd.Cat.Name(), cd.Dog.Name())
}

//******************************************************************************
//* Main practice function of package
//******************************************************************************

func practiceInterface() {
	var animal1 Animal
	animal1 = &Dog{}
	fmt.Println(animal1.Name())

	var animal2 Animal
	animal2 = &Cat{}
	fmt.Println(animal2.Name())

	var animal3 Animal
	animal3 = &CatDog{}
	fmt.Println(animal3.Name())

	var partyAnimal PartyAnimal
	partyAnimal = &Raver{}
	partyAnimal.Party()

	gd := &GuideDog{}
	gd.Help()
}
