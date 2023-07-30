package godicontainer

type ResolvableInterface interface {
	Test() string
}

type OtherResolvableInterface interface {
	Test() string
}

type ResolvableConcrete struct {
	Resolvable OtherResolvableInterface `di:"autowire"`
	tst        int
}

func (r *ResolvableConcrete) Test() string {
	return "It works"
}

func newResolvableConrete() (interface{}, error) {
	return &ResolvableConcrete{}, nil
}

type OtherResolvableConcrete struct {
}

func (r *OtherResolvableConcrete) Test() string {
	return "It works as other concrete implementation"
}

func (r *OtherResolvableConcrete) Construct() {

}

func newOtherResolvableConrete() (interface{}, error) {
	return &OtherResolvableConcrete{}, nil
}
