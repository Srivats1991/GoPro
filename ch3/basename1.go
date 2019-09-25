//basename removes directory components and .suffix
//eg., a=>a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
package main

import ( 
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		fmt.Println(basename(input.Text()))
	}
}

func basename(s string) string {
	//discard last '/' and everything before.
	for i:=len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	//preserve everything before last '.'.
	for i:=len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
