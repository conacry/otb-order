package customerEntity

import "github.com/conacry/go-platform/pkg/errors"

type Builder struct {
	customerID *CustomerID
	profile    *Profile

	errors *errors.Errors
}

func NewBuilder() *Builder {
	return &Builder{
		errors: errors.NewErrors(),
	}
}

func (b *Builder) CustomerID(customerID *CustomerID) *Builder {
	b.customerID = customerID
	return b
}

func (b *Builder) Profile(profile *Profile) *Builder {
	b.profile = profile
	return b
}

func (b *Builder) Build() (*Customer, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	return b.createFromBuilder(), nil
}

func (b *Builder) checkRequiredFields() {
	if b.customerID == nil {
		b.errors.AddError(ErrCustomerIDIsRequired)
	}
	if b.profile == nil {
		b.errors.AddError(ErrProfileIsRequired)
	}
}
func (b *Builder) createFromBuilder() *Customer {
	return &Customer{
		customerID: b.customerID,
		profile:    b.profile,
	}
}
