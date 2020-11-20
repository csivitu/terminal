package main

import (
        "fmt"
        "flag"
)

func main()  {
        fmt.Println("Hello world!")
		name := flag.String("n","", "Name of the session you want to create")
		join := flag.String("j","", "Name of session you want to join")
        help := flag.Bool("h", false, "This is the help command")
        flag.Parse()

        if *help {
                fmt.Println("1. -n <name of sessiom> => Creates a new session with the given name, Leave back if you want to join The default session!")
				fmt.Println("1. -j <name of sessiom> => Joins a session with the given name if the session exists.")
				*join = ""
				*name = ""
		}
		
		if *join != "" {
			fmt.Println("Joining session with name: ",*join) // first check if session exists, then join
			*name = ""
		}

        if *name == "" && *join == "" && !*help{
                fmt.Println("Joining current session.")
		} 

		if *name != "" {
                fmt.Println("Making new session with name: ", *name)
        }
}