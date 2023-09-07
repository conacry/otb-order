package productEntity

import "github.com/conacry/go-platform/pkg/errors"

type Builder struct {
	productID   *ProductID
	title       *Title
	description *Description
	cost        *Cost

	errors *errors.Errors
}

func NewBuilder() *Builder {
	return &Builder{
		errors: errors.NewErrors(),
	}
}

func (b *Builder) ProductID(productID *ProductID) *Builder {
	b.productID = productID
	return b
}

func (b *Builder) Title(title *Title) *Builder {
	b.title = title
	return b
}

func (b *Builder) Description(description *Description) *Builder {
	b.description = description
	return b
}

func (b *Builder) Cost(cost *Cost) *Builder {
	b.cost = cost
	return b
}

func (b *Builder) Build() (*Product, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	return b.createFromBuilder(), nil
}

func (b *Builder) checkRequiredFields() {
	if b.productID == nil {
		b.errors.AddError(ErrProductIDIsRequired)
	}
	if b.title == nil {
		b.errors.AddError(ErrTitleIsRequired)
	}
	if b.description == nil {
		b.errors.AddError(ErrDescriptionIsRequired)
	}
	if b.cost == nil {
		b.errors.AddError(ErrCostIsRequired)
	}
}

func (b *Builder) createFromBuilder() *Product {
	return &Product{
		productID:   b.productID,
		title:       b.title,
		description: b.description,
		cost:        b.cost,
	}
}
