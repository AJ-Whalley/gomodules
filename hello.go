package hello

import (
	"rsc.io/quote/v3"
)

//Hello returns "Hello World"
func Hello() string {
	return quote.HelloV3()
}

//Proverb returns a Go proverb
func Proverb() string {
	return quote.Concurrency()
}
