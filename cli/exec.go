// SPDX-License-Identifier: MIT-0
// LICENSE: https://spdx.org/licenses/MIT-0.html

package cli

import (
	"os"
)

// Exec .
func Exec() {
	args := os.Args
	if len(args) > 1 {
		switch args[1] {
		case "version", "-version", "--version":
			{
				version()
			}	
		default:
			{
				// fmt.Println("print help section here")
			}
		}
	} else {
		run()
	}
}
