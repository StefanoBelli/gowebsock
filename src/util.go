package main

import (
	"os"
	"errors"
	"fmt"
)

func getUrl() (string,error) {
	arg := os.Args[1:]

	for i,e := range arg {
		if ( e == "-u" ) {
			return arg[i+1],nil
		}
	}

	return "",errors.New("nourl")
}

func pinfo(msg ... string) {
	var final string 
	const INFO_PREFIX = "(wsc-info) "

	final = INFO_PREFIX
	for _,m := range msg {
		final += m + " "
	}

	fmt.Printf("%s\n",final)
}

func perr(msg ... string) {
	var final string
	const ERROR_PREFIX = "(wsc-err) "

	final = ERROR_PREFIX
	for _,m := range msg {
		final += m + " "
	}

	fmt.Fprintf(os.Stderr,"%s\n",final)
}

func setOptMap() map[string]string {
	const CMDS_CNT int = 5
	
	m := make(map[string]string,CMDS_CNT)
	m["help"] = "prints this help"
	m["url"] = "set url: \"url <wsurl>\""
	m["connect"] = "connects to the chosen URL"
	m["quit"] = "exits wsclient, same as exit"
	m["exit"] = "exits wsclient, same as quit"
	
	return m
}

func printHelp(helpm map[string]string) {
	for i,v := range helpm {
		fmt.Printf(" %s : %s \n",i,v)
	}
}

