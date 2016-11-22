package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func loopPrompt(helpmap map[string]string) {
	const NIL_CMD = "null-string-cmd"

	proto := ""
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
			if proto != "" {
				pinfo("Connecting to:",url,"using protocol:",proto)
			} else {
				pinfo("Connecting to:",url)
			}

			establishAndDialog(url,proto)
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

		if len(argcmd) == 1 {
			pinfo("Argument required for",argcmd[0])
			continue
		}

		if argcmd[0] == "url" {
			url = argcmd[1]
			pinfo("Set URL:",url)
		} else if argcmd[0] == "proto" {
			if argcmd[1] == "default" {
				proto = ""
				pinfo("Set default proto")
			} else {
				proto = argcmd[1]
				pinfo("Set Proto:",proto)
			}
		}
	}
}

func main() {
	arg := getArgs()
	fmt.Println("WebSocket message exchanger")
	if arg.url != ARGS_NOURL && arg.proto == ARGS_NOPROTO {
		pinfo("Connecting to:",arg.url)
		establishAndDialog(arg.url,"")
	} else if arg.url != ARGS_NOURL && arg.proto != ARGS_NOPROTO {
		pinfo("Connecting to:",arg.url,"using protocol:",arg.proto)
		establishAndDialog(arg.url,arg.proto)
	} else {
		fmt.Println("Start setting url: \"url myurl\"")
		fmt.Println("Then connect using: \"connect\"")
	}

	loopPrompt(setOptMap())
}
