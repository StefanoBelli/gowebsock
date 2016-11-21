package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func loopPrompt(helpmap map[string]string) {
	const NIL_CMD = "null-string-cmd"
	url := NIL_CMD
	
	buf_reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf(" >>> ")
		cmd,_ := buf_reader.ReadString('\n')
		cmd = strings.Replace(cmd,"\n","",-1)

		if len(cmd) == 0 {
			perr("Empty command")
			continue
		}
		
		if cmd == "exit" || cmd == "quit" {
			pinfo("Exiting... Bye bye!")
			os.Exit(0)
		} else if cmd == "connect" && url != NIL_CMD {
			pinfo("Connecting to:",url)
			establishAndDialog(url)
			continue
		} else if cmd == "connect" && url == NIL_CMD {
			perr("You need to set an URL before connecting")
			perr("See \"help\"")
			continue
		} else if cmd == "help" {
			printHelp(helpmap)
			continue
		}

		argcmd := strings.Split(cmd," ")

		if argcmd[0] == "url" && len(argcmd[1]) > 0 {
			url = argcmd[1]
			pinfo("Set URL:",url)
		} else if argcmd[0] == "url" && len(argcmd[1]) == 0 {
			perr("url argument cannot be empty")
		}
	}
}

func main() {
	url,err := getUrl()

	fmt.Println("WebSocket message exchanger")
	if err == nil {
		pinfo("Connecting to:",url);
		establishAndDialog(url)
	} else {
		fmt.Println("Start setting url: \"url myurl\"")
		fmt.Println("Then connect using: \"connect\"")
	}

	loopPrompt(setOptMap())
}
