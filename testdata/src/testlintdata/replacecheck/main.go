// want "replace github.com/pentops/return42 => ../return42\n"
package main

import (
	"fmt"
	"github.com/pentops/return42"
)

func main() {
	fmt.Println(return42.DoIt())
}
