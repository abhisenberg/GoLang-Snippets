package chapter1

import (
	"fmt"
	"os"
	"strings"
)

/**
Accessing command line args using os.Args
Args is a slice of strings, where Args[0] = the name of the command itself (e.g. run)
and the user specified args are in the range os.Args[1:len(Args)]
**/
func CommandLineArgs() {
	var s string
	for i := 1; i < len(os.Args); i++ {
		s += os.Args[i] + " ~ "
	}
	fmt.Println(s)
}

/*
Using for-range
*/
func CommandLineArgs2() {
	var s string
	for _, value := range os.Args[1:] {
		s += value + " ~ "
	}
	fmt.Println(s)
}

/*
Using join
*/
func CommandLineArgs3() {
	var s string
	s = strings.Join(os.Args[1:], " ~ ")
	fmt.Println(s)
}
