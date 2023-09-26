package customerEntityStub

import (
	"github.com/conacry/go-platform/pkg/generator"
	customerEntity "otb-order/domain/entity/customer"
)

func GetProfile() *customerEntity.Profile {
	firstName := generator.RandomDefaultStr()
	lastName := generator.RandomDefaultStr()

	profile, err := customerEntity.ProfileFrom(firstName, lastName)
	if err != nil {
		return nil
	}

	return &profile
}
