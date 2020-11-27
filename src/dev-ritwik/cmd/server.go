package cmd

import(
	"log"
    "net"
    "fmt"
    //"strings"
)

func Server(c net.Conn) {
    for {
        buf := make([]byte, 512)
        nr, err := c.Read(buf)
        if err != nil {
            return
        }
        data := buf[0:nr]
        //data = strings.TrimSpace(data)
        println(string(data))
        _, err = c.Write(data)
        if err != nil {
            log.Fatal("Write: ", err)
        }
        var wr string 
		fmt.Scanln(&wr)
		wr = wr+"\n"
		b := []byte(wr)

		_, e := c.Write(b)

		if e!=nil {
			log.Fatal("Write error: ",e)
		}

    }
    c.Close()
}