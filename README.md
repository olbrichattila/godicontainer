# Depencency injection container for golang

!!! Not this is the very first version of the container, not yet production ready, it is in a prototype stage !!!


## Usage:

Mapping, resolving a struct via interface


```
package main

import (
	"fmt"

	"github.com/olbrichattila/godicontainer"
)

// App to centralize the container
type App struct {
	container *godicontainer.Container
}

type AnimalInterface interface {
	Eats() string
}

type Dog struct {
}

// Animal also can be a cat
type Cat struct {
}

func (a *Dog) Eats() string {
	return "Dog eats cat"
}

func (a *Cat) Eats() string {
	return "Cat eats mouse if not eaten by the dog"
}

func NewCat() (interface{}, error) {
	return &Cat{}, nil
}

func NewDog() (interface{}, error) {
	return &Dog{}, nil
}

func newApp() *App {
	return &App{
		container: godicontainer.NewContainer(),
	}
}

func main() {
	app := newApp()
	app.container.Set("AnimalInterface", NewCat)

	animal, err := app.container.Get("AnimalInterface")
	if err != nil {
		fmt.Println(err)
		return
	}

	eated := animal.(AnimalInterface).Eats()

	fmt.Println(eated)

	app.container.Set("AnimalInterface", NewDog)

	animal, err = app.container.Get("AnimalInterface")
	if err != nil {
		fmt.Println(err)
		return
	}

	eated = animal.(AnimalInterface).Eats()

	fmt.Println(eated)
}
```

## Autowire dependencies to a struct

Create a struct and add your properites with interface type hint, 
Add `di:"autowire"` annotation to your struct to be auto wired


```
package main

import (
	"fmt"

	"github.com/olbrichattila/godicontainer"
)

type App struct {
	container *godicontainer.Container
}

type AnimalInterface interface {
	Eats() string
}

type Dog struct {
}

type Cat struct {
}

type HumanInterface interface {
	Say() string
}

type Human struct {
}

// Add your annotations to be autowired
type Creatures struct {
	Animal AnimalInterface `di:"autowire"`
	Human  HumanInterface `di:"autowire"`
}

func (a *Dog) Eats() string {
	return "Dog eats cat"
}

func (a *Cat) Eats() string {
	return "Cat eats mouse if not eaten by the dog"
}

func (h *Human) Say() string {
	return "Human says hello world"
}

func NewCat() (interface{}, error) {
	return &Cat{}, nil
}

func NewDog() (interface{}, error) {
	return &Dog{}, nil
}

func NewHuman() (interface{}, error) {
	return &Human{}, nil
}

func newApp() *App {
	return &App{
		container: godicontainer.NewContainer(),
	}
}

func main() {
	app := newApp()
	app.container.Set("AnimalInterface", NewCat)
	app.container.Set("HumanInterface", NewHuman)

	// You cannot use a ponter here
	creatures := Creatures{}

	// You have to pass the stucture instance and it's pointer as well
	// The original strut will contain your changes
	// (this may change in the future)
	app.container.ResolvDependencies(creatures, &creatures)

	fmt.Println(creatures.Animal.Eats())
	fmt.Println(creatures.Human.Say())
}
```



