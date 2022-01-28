package myquote

import (
	"fmt"

	"rsc.io/quote"
)

func PrintQuotes() {
	fmt.Println(GetGlass())
	fmt.Println(GetGo())
	fmt.Println(GetHello())
	fmt.Println(GetOpt())
}

func GetGlass() string {
	return quote.Glass()
}
func GetGo() string {
	return quote.Go()
}
func GetHello() string {
	return quote.Hello()
}
func GetOpt() string {
	return quote.Opt()
}
