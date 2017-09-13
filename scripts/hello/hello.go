package main

import ("fmt"
        "github.com/golang/example/stringutil"
)

func main() {
	fmt.Printf("Hello, world.\n")
        fmt.Printf(stringutil.Reverse("\n Hello, world.\n"))
}
