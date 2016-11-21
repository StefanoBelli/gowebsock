package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"bufio"
	"os"
	"strings"
)

func establishAndDialog(url string) {
	origin := "http://localhost/"
	wsc,err := websocket.Dial(url,"",origin)

	if err != nil {
		perr("Cannot connect to:",url)
		return
	}

	pinfo("Connected to:",url)

	var destmsg = make([]byte,1024)
	var str string 

	buf_reader := bufio.NewReader(os.Stdin)
	buf_writer := bufio.NewWriter(os.Stdout)
	for {
		buf_writer.Flush()
		fmt.Printf(" (ws) >>> ")

		str,_ = buf_reader.ReadString('\n')
		str = strings.Replace(str,"\n","",-1)

		if len(str) == 0 {
			continue
		}

		if str == "/close" {
			break
		} else if str == "//close" {
			str = "/close"
		}
		
		if str == "/force_close" {
			pinfo("FORCE CLOSING")
			return
		} else if str == "//force_close" {
			str = "/force_close"
		}

		if _, err := wsc.Write([]byte(str)) ; err != nil {
			fmt.Printf(" (disconnected) Connection with %s done.\n",url)
			break
		}

		if bs, err := wsc.Read(destmsg) ; err == nil && bs > 0 {
			fmt.Printf(" (ws) < %s \n",destmsg)
		} else if err != nil || bs <= 0 {
			fmt.Printf(" (disconnected) Connection with %s done.\n",url)
			break
		}
	}

	defer buf_writer.Flush()

	pinfo("Closing connection with:",url)
	wsc.Close()
}

