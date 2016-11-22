package main

import (
	"os"
	"fmt"
)

const (
	ARGS_NOURL string = "null-url-from-args"
	ARGS_NOPROTO string = "null-proto-from-args"
)

type args struct {
	url string
	proto string
}

func getArgs() args {
	arg := os.Args[1:]
	_args := args{ARGS_NOURL,ARGS_NOPROTO}

	for i,e := range arg {
		if ( e == "-u" ) {
			_args.url=arg[i+1]
		}

		if ( e == "-p" ) {
			_args.proto=arg[i+1]
		}
	}

	return _args
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
	m["proto"] = "set protocol: \"proto <protocol>\", use \"proto default\" to use the default one"
	
	return m
}

func printHelp(helpm map[string]string) {
	for i,v := range helpm {
		fmt.Printf(" %s : %s \n",i,v)
	}
}

