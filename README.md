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

Autowire dependencies to a struct

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

type Creatures struct {
	Animal AnimalInterface
	Human  HumanInterface
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

	creatures := &Creatures{}

	app.container.ResolvDependencies(Creatures{}, creatures)

	fmt.Println(creatures.Animal.Eats())
	fmt.Println(creatures.Human.Say())
}
```

