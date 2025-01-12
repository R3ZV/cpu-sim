//go:build debug
// +build debug

package log

import "fmt"

func Debug(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
}
