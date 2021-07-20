package hello

import (
	quoteV3 "rsc.io/quote/v3" // using two versions of same lib
)

func Hello() string {
	return quoteV3.HelloV3()
}

func Proverb() string {
	return quoteV3.Concurrency()
}
