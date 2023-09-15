package productEntity

import "unicode/utf8"

const DescriptionMaxLength = 2000

type Description string

func DescriptionFrom(text string) (Description, error) {
	if text == "" {
		return "", ErrDescriptionIsEmpty
	}

	if utf8.RuneCountInString(text) > DescriptionMaxLength {
		return "", ErrDescriptionTooLong
	}

	return Description(text), nil
}

func (t Description) String() string {
	return string(t)
}
