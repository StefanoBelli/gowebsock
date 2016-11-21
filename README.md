## WebSocket Tester 
===

Simple CLI-based program to test WebSocket connection, written in Golang.

*Very simple usage*

### Compile
===
	
	cd src
	go build *.go

You may need to set GOPATH env var and download golang.org/x/net/websocket

### Run
===

There are two ways to connect to wsocket:

 * Using directly the cli:
 ~~~
 ./main -u ws://echo.websocket.org
 ~~~

 * Using the interactive mode:
 ~~~
> ./main 
WebSocket message exchanger                                                                                                                                             
Start setting url: "url myurl"                                                                                                                                          
Then connect using: "connect"                                                                                                                                           
 >>> url ws://echo.websocket.org                                                                                                                                        
(wsc-info) Set URL: ws://echo.websocket.org                                                                                                                             
 >>> connect                                                                                                                                                            
(wsc-info) Connecting to: ws://echo.websocket.org                                                                                                                       
(wsc-info) Connected to: ws://echo.websocket.org                                                                                                                        
 (ws) >>> /close                                                                                                                                                        
(wsc-info) Closing connection with: ws://echo.websocket.org 
 >>> 
  ~~~
Anyway, use help to get help :)

PS. */close* does not exit program, just closes connection 

PS. In websocket connection, use *//close* and *//force_close* to send */close* and */force_close* messages to the server and avoid closing connection


