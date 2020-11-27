package cmd

import(
	"os"
	"fmt"
	"net"
	"log"
)

func CreateSession(name string){
	
	addr := "/tmp/"+name+".sock"
	if _, err := os.Stat(addr); err == nil {
		if _, err := os.Stat("/path/to/whatever"); err == nil {
  		fmt.Println("Session with name"+name+"already active.")
		return
}
	  
	  }

	l, err := net.Listen("unix", addr)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer l.Close()

	for {
		fd, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		go Server(fd)
		//go Joiner(addr)
	}
}