package godicontainer

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ContainerTestSuite struct {
	suite.Suite
	container *Container
}

func TestContainerRunner(t *testing.T) {
	suite.Run(t, new(ContainerTestSuite))
}

func (t *ContainerTestSuite) SetupTest() {
	t.container = NewContainer()
}

func (t *ContainerTestSuite) TestSetAndGetWorskWithTheSameInterfaceAndDIfferentImplementation() {
	resolveInterfaceName := "ResolvableInterface"
	t.container.Set(resolveInterfaceName, newResolvableConrete)
	resolved, err := t.container.Get(resolveInterfaceName)

	t.Nil(err)

	asInterface, ok := resolved.(ResolvableInterface)

	t.True(ok)

	text := asInterface.Test()

	t.Equal("It works", text)

	t.container.Set(resolveInterfaceName, newOtherResolvableConrete)
	resolved, err = t.container.Get(resolveInterfaceName)

	t.Nil(err)

	asInterface, ok = resolved.(ResolvableInterface)

	t.True(ok)

	text = asInterface.Test()

	t.Equal("It works as other concrete implementation", text)
}

func (t *ContainerTestSuite) TestResolveDepencencies() {
	s := ResolvableConcrete{}
	resolveInterfaceName := "OtherResolvableInterface"
	t.container.Set(resolveInterfaceName, newResolvableConrete)

	t.container.ResolvDependencies(s, &s)
	s.Resolvable.Test()
	t.Equal("It works", s.Resolvable.Test())

	t.container.Set(resolveInterfaceName, newOtherResolvableConrete)
	t.container.ResolvDependencies(s, &s)

	t.Equal("It works as other concrete implementation", s.Resolvable.Test())
}

func (t *ContainerTestSuite) TestSetDefinitions() {

	definitions := CallbackDefinitions{
		"ResolvableInterface":      newResolvableConrete,
		"OtherResolvableInterface": newOtherResolvableConrete,
	}

	t.container.SetDefinitions(definitions)

	resolvableInstance, err := t.container.Get("ResolvableInterface")
	t.Nil(err)

	resolved, ok := resolvableInstance.(ResolvableInterface)

	t.True(ok)

	t.Equal("It works", resolved.Test())

	resolvableInstance, err = t.container.Get("OtherResolvableInterface")
	t.Nil(err)

	resolved, ok = resolvableInstance.(ResolvableInterface)

	t.True(ok)

	t.Equal("It works as other concrete implementation", resolved.Test())
}
