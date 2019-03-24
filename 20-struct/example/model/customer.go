package model

// Customer struct
type Customer struct { // exported struct type
	Id      int     // exported field
	Name    string  // exported field
	Address Address // unexported field, only accessible inside package `model`
	Married bool    // unexported field, only accessible inside package `model`
}
