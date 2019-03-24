package model

// unexported struct (only accessible inside package model)
type Address struct {
	houseNo, stree, city, state, country string
	zipCode                              int
}
