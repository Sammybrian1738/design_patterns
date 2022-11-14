package builder

import "builder_method/product"

type IBuilder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor()
	GetHouse() product.House
}

func GetBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return NewNormalBuilder()
	}

	if builderType == "igloo" {
		return NewIglooBuilder()
	}
	return nil
}
