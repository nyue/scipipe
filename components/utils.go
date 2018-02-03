package components

import (
	"github.com/scipipe/scipipe"
)

// Check checks the error err, prints an error message and exits the program
func Check(err error) {
	scipipe.CheckErr(err)
}
