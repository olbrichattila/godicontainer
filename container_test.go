package godicontainer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type containerTestSuite struct {
	suite.Suite
	container *Container
}

func TestContainerRunner(t *testing.T) {
	suite.Run(t, new(containerTestSuite))
}

func (t *containerTestSuite) SetupTest() {
	t.container = NewContainer()
	fmt.Println("Running test", t.T().Name())
}

func (t *containerTestSuite) TestSetAndGetWorskWithTheSameInterfaceAndDIfferentImplementation() {
	resolveInterfaceName := "resolvableInterface"
	t.container.Set(resolveInterfaceName, newResolvableConrete)
	resolved, err := t.container.Get(resolveInterfaceName)

	t.Nil(err)

	asInterface, ok := resolved.(resolvableInterface)

	t.True(ok)

	text := asInterface.Test()

	t.Equal("It works", text)

	t.container.Set(resolveInterfaceName, newOtherResolvableConrete)
	resolved, err = t.container.Get(resolveInterfaceName)

	t.Nil(err)

	asInterface, ok = resolved.(resolvableInterface)

	t.True(ok)

	text = asInterface.Test()

	t.Equal("It works as other concrete implementation", text)
}

func (t *containerTestSuite) TestResolveDepencencies() {
	s := resolvableConcrete{}
	resolveInterfaceName := "otherresolvableInterface"
	t.container.Set(resolveInterfaceName, newResolvableConrete)

	t.container.ResolvDependencies(s, &s)
	s.Resolvable.Test()
	t.Equal("It works", s.Resolvable.Test())

	t.container.Set(resolveInterfaceName, newOtherResolvableConrete)
	t.container.ResolvDependencies(s, &s)

	t.Equal("It works as other concrete implementation", s.Resolvable.Test())
}

func (t *containerTestSuite) TestSetDefinitions() {

	definitions := CallbackDefinitions{
		"resolvableInterface":      newResolvableConrete,
		"otherresolvableInterface": newOtherResolvableConrete,
	}

	t.container.SetDefinitions(definitions)

	resolvableInstance, err := t.container.Get("resolvableInterface")
	t.Nil(err)

	resolved, ok := resolvableInstance.(resolvableInterface)

	t.True(ok)

	t.Equal("It works", resolved.Test())

	resolvableInstance, err = t.container.Get("otherresolvableInterface")
	t.Nil(err)

	resolved, ok = resolvableInstance.(resolvableInterface)

	t.True(ok)

	t.Equal("It works as other concrete implementation", resolved.Test())
}
