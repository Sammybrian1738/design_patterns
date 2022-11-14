package concretefactory

import (
	abstractproduct "abstract_factory_method/abstract_product"
	concreteproduct "abstract_factory_method/concrete_product"
)

type Adidas struct {
}

func (a *Adidas) MakeShoe() abstractproduct.IShoe {
	return &concreteproduct.AdidasShoe{
		Shoe: abstractproduct.Shoe{
			Logo: "adidas",
			Size: 14,
		},
	}
}

func (a *Adidas) MakeShirt() abstractproduct.IShirt {
	return &concreteproduct.AdidasShirt{
		Shirt: abstractproduct.Shirt{
			Logo: "adidas",
			Size: 14,
		},
	}
}
