package main

func walk(x interface{}, f func(input string)) {
	f("hello")
}
