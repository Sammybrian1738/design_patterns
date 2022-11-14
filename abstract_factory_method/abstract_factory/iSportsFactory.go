// abstract factory interface
package abstractfactory

import (
	abstractproduct "abstract_factory_method/abstract_product"
	concretefactory "abstract_factory_method/concrete_factory"

	"fmt"
)

type ISportsFactory interface {
	MakeShoe() abstractproduct.IShoe
	MakeShirt() abstractproduct.IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &concretefactory.Adidas{}, nil
	}

	if brand == "nike" {
		return &concretefactory.Nike{}, nil
	}

	return nil, fmt.Errorf("wrong brand type passed")
}
