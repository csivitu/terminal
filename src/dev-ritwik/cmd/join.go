package cmd

import(
	"os"
	"fmt"
	"net"
	"log"
	"strings"
)

func Joiner(name string) {

	addr := "/tmp/"+name+".sock"

	if _, err := os.Stat(addr); os.IsNotExist(err) {
		fmt.Println("Session with name"+name+"does not exist.")
		return
	  }

	c, err := net.Dial("unix", addr)

	if err!=nil {
		log.Fatal("Error while joining session")
		return
	}

	fmt.Println("Successfully joined session: "+name)

	for {
		var wr string 
		fmt.Scanln(&wr)
		wr = wr+"\n"
		b := []byte(wr)

		_, e := c.Write(b)
		
		if e!=nil {
			log.Fatal("Write error: ",e)
		}

		buf :=make([]byte, 1024)
		n, err := c.Read(buf)
		
		if err!=nil {
			log.Fatal("Error occurred while recieving data: ", err)
		}

		data := string(buf[0:n])
		data = strings.TrimSpace(data)

		fmt.Println(data)

	}

}