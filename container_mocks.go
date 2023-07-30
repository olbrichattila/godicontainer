package godicontainer

type resolvableInterface interface {
	Test() string
}

type otherresolvableInterface interface {
	Test() string
}

type resolvableConcrete struct {
	Resolvable otherresolvableInterface `di:"autowire"`
	tst        int
}

func (r *resolvableConcrete) Test() string {
	return "It works"
}

func newResolvableConrete() (interface{}, error) {
	return &resolvableConcrete{}, nil
}

type otherresolvableConcrete struct {
}

func (r *otherresolvableConcrete) Test() string {
	return "It works as other concrete implementation"
}

func (r *otherresolvableConcrete) Construct() {

}

func newOtherResolvableConrete() (interface{}, error) {
	return &otherresolvableConcrete{}, nil
}
