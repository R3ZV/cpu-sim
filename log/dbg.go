//go:build debug
// +build debug

package log

import "fmt"

func Debug(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
}

func Assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}
