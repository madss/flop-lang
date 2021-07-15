package interpreter

import "fmt"

type Builtin func(args ...interface{}) error

func print(args ...interface{}) error {
	fmt.Print(args...)
	return nil
}
