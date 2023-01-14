package log

import "fmt"

func Wait(a string) {
	fmt.Printf("[wait] %s\n", a)
}

func Info(a string) {
	fmt.Printf("[info] %s\n", a)
}

func Error(a string) {
	fmt.Printf("[error] %s\n", a)
}
