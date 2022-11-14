package director

import (
	"builder_method/builder"
	"builder_method/product"
)

type Director struct {
	builder builder.IBuilder
}

func NewDirector(b builder.IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) SetBuilder(b builder.IBuilder) {
	d.builder = b
}

func (d *Director) BuildHouse() product.House {
	d.builder.SetDoorType()
	d.builder.SetWindowType()
	d.builder.SetNumFloor()
	return d.builder.GetHouse()
}
