package customerEntity

import "github.com/conacry/go-platform/pkg/errors"

type Profile struct {
	firstName string
	lastName  string

	errors *errors.Errors
}

func ProfileFrom(firstName string, lastName string) (Profile, error) {
	if firstName == "" {
		return Profile{}, ErrFirstNameProfileIsEmpty
	}
	if lastName == "" {
		return Profile{}, ErrLastNameProfileIsEmpty
	}

	return Profile{
		firstName: firstName,
		lastName:  lastName,
	}, nil
}

func (p Profile) String() string {
	return p.firstName + p.lastName
}
