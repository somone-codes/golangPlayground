package hello

import (
	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3" // using two versions of same lib
)

func Hello() string {
	return quote.Hello()
}

func Proverb() string {
	return quoteV3.Concurrency()
}
