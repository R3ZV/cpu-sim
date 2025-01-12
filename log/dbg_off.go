//go:build !debug
// +build !debug

package log

func Debug(msg string, args ...interface{}) {}

func Assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}
