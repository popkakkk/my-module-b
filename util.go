package b_package

import (
	"fmt"

	c "github.com/popkakkk/my-module-c" // Adjust the import path as necessary
)

func UtilFunction() string {
	fmt.Println("use package c  --> ", c.UtilFunction())
	return "This is a utility function from package b"
}
