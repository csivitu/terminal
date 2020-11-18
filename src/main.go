package main

import(
	"fmt"
	"flag"
)
func main(){
	seshFlag := flag.String("s", "", "Flag for connecting to an existing session s")
	flag.Parse()
	fmt.Println(*seshFlag)
}