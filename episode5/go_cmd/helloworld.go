package main

import "flag"

var intflagvar int
var stringflagvar string
var boolvar bool
//var
func main() {
	flag.IntVar(&intflagvar, "integer flag", 9090, "help message for this flag")
	flag.StringVar(&stringflagvar, "name of flag", "default value", "help message for this flag")
	flag.StringVar(&boolvar, "name of flag", true, "help message")
	flag.Parse()
}
