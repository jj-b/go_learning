package main

import (
	"fmt"
	"os"
	"strings"

	/*import lib to handle commanddline flags
	  We also give it an alias 'flag'*/

	flag "github.com/ogier/pflag"
)

//Declare variables
//We can declare them outside a function, as case denotes scope
var (	user string)
//main function
func main() {
	//Call a method to parse the flags
	flag.Parse()
	// if user does not supply flags, print usage
	// we can clean this up later by putting this into its own function
	if flag.NFlag() == 0 {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}
	users := strings.Split(user, ",")
	fmt.Printf("Searching user(s): %s\n", users)
}

//A function to define our flags
func init() {
	/*We are evaluating expecting a string
	  We want a posix complient flag
	  We want to bind a variable to this flag*/
	flag.StringVarP(&user, "user", "u", "", "Search Users")
}
