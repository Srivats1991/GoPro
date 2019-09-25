//Cf converts its numeric argument to celsius and fahrenheit
package main

import (
	"fmt"
	"os"
	"strconv"
	"~/Desktop/golang/ch2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahreheit(t)
		c := tempvonv.Celisius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c,
			tempconv.CToF(c))
	}
}
