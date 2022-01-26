package domain

type Location string

func NewLocation(location string) Location {
	return Location(location)
}

func (l Location) String() string {
	return string(l)
}
