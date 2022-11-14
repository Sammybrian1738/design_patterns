package concretefactory

import (
	abstractproduct "abstract_factory_method/abstract_product"
	concreteproduct "abstract_factory_method/concrete_product"
)

type Nike struct {
}

func (n *Nike) MakeShoe() abstractproduct.IShoe {
	return &concreteproduct.NikeShoe{
		Shoe: abstractproduct.Shoe{
			Logo: "nike",
			Size: 14,
		},
	}
}

func (n *Nike) MakeShirt() abstractproduct.IShirt {
	return &concreteproduct.NikeShirt{
		Shirt: abstractproduct.Shirt{
			Logo: "nike",
			Size: 14,
		},
	}
}
