package main

import "flag"

var consumerKeyPtr = flag.String("ckey", "", "Your 500px consumer key")

func init() {
	flag.StringVar(consumerKeyPtr, "c", "bar", "a string var")
}

func main() {
	flag.Parse()
	consumerKey := *consumerKeyPtr

	println(consumerKey)
	println("done")
}
