package main

import (
        "fmt"
		"flag"
		"os"
		"net"
		"log"
		"strings"
)

func main()  {
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
				createServer(*name)
        }
}

func createServer(name string) {

	addr := "/tmp/"+name+".sock"
	if fileExists(addr){
		fmt.Println("Session with name"+name+"already active.")
		return
	}

	l, err := net.Listen("unix",addr)

	if err!=nil {
		log.Fatal("Error occurred :",err)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()

		if err!= nil {
			log.Fatal("Error occurred :",err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {

	for {

		buf :=make([]byte, 1024)
		n, err := c.Read(buf)
		
		if err!=nil {
			log.Fatal("Error occurred while recieving data: ", err)
		}

		data := string(buf[0:n])
		data = strings.TrimSpace(data)

		fmt.Println(data)

		var wr string 
		fmt.Scanln(&wr)
		wr = wr+"\n"
		b := []byte(wr)

		_, e := c.Write(b)

		if e!=nil {
			log.Fatal("Write error: ",err)
		}

	}
	
	
}


func fileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}