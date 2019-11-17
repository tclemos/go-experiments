package main

import (
	"fmt"

	// import pkg1 to be used with the qualifier pkg1
	"github.com/tclemos/dot-import/pkg1"

	// import pkg1 to be used without qualifier
	// go-lint also recommends to not use dot imports
	. "github.com/tclemos/dot-import/pkg1"

	// import pkg2 to be used with the qualifier pkg2
	"github.com/tclemos/dot-import/pkg2"
	// this import is not allowed because the pkg1 already declares a function with the same name
	// . "github.com/tclemos/dot-import/pkg2"
)

func main() {
	pkg1.Colide() // this uses the import from line 6
	pkg2.Colide() // this uses the import from line 8
	Colide()      // this uses the import from line 7
}

// Colide is discarded by the compiler, because we are importing pkg1 using .
func Colide() {
	fmt.Println("MAIN - Colide")
}
