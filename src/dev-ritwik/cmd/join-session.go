package cmd

import(
	"os"
	"log"
	"fmt"
	"net"
)

func JoinSession(name string){
	addr := "/tmp/"+name+".sock"
	if _, err := os.Stat(addr); os.IsNotExist(err) {
		fmt.Println("Session with name"+name+"does not exist.")
		return
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