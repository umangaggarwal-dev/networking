package tcp

import (
    "fmt"
    "net"
    "os"
    "strconv"
)

const (
    DEFAULT_PROTOCOL = "tcp"
)

type TcpServer struct {
    host string
    port int
    protocol string
    handler func([]byte)
}

func (t *TcpServer) StartListening() {
    // start a listener
    listener, err := net.Listen(DEFAULT_PROTOCOL, t.host + ":" + strconv.Itoa(t.port))
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    
    // close the listener before returning
    defer listener.Close()
    
    fmt.Println("Started listening on port", listener.Addr())
    for {
        // accept incoming request
        con, err := listener.Accept()
        if err != nil {
            fmt.Println("Failed to build connection:", err.Error())
        }
        t.handleRequest(con)
    }
}

func (t *TcpServer) RegisterHandler(handler func([]byte)) {
    t.handler = handler
}

func (t *TcpServer) handleRequest(conn net.Conn) {
  defer conn.Close()
  buf := make([]byte, 1024)
  // read the incoming data into buffer
  _, err := conn.Read(buf)
  if err != nil {
    fmt.Println("Error reading:", err.Error())
  }
  t.handler(buf)
}

func NewServer(host string, port int) *TcpServer {
    tcpServer := new(TcpServer)
    tcpServer.host = host
    tcpServer.port = port
    return tcpServer
}

