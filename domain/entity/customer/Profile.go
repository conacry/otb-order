package customerEntity

type Profile struct {
	firstName string
	lastName  string
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

func (p *Profile) FirstName() string {
	return p.firstName
}

func (p *Profile) LastName() string {
	return p.lastName
}

func (p *Profile) String() string {
	return p.firstName + " " + p.lastName
}
