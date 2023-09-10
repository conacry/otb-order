package entityStub

import (
	"github.com/conacry/go-platform/pkg/generator"
	customerEntity "online-shop-order/internal/domain/entity/customer"
)

func Profile() *customerEntity.Profile {
	firstName := generator.RandomDefaultStr()
	lastName := generator.RandomDefaultStr()

	profile, err := customerEntity.ProfileFrom(firstName, lastName)
	if err != nil {
		return nil
	}

	return &profile
}
