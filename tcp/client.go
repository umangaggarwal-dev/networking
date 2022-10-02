package tcp

import (
    "net"
    "strconv"
    "time"
    "io"
    "fmt"
)

const (
    Protocol = "tcp"
)

type TCPConnection struct {
    host string
    port int
    conn net.Conn
}

func (t *TCPConnection) Connect() (bool, error) {
    conn, err := net.Dial(Protocol, t.host + ":" + strconv.Itoa(t.port))
    if err != nil {
        return false, err
    }
    t.conn = conn
    return true, nil
}

func (t *TCPConnection) Close() (bool, error) {
    return true, nil
}

func (t *TCPConnection) IsAlive() (bool) {
    var data []byte
    t.conn.SetReadDeadline(time.Now())
    _, err := t.conn.Read(data)
    if err == io.EOF {
        t.conn.Close()
        t.conn = nil
        return false
    }
    fmt.Println(string(data[:]))
    var zero time.Time
    t.conn.SetReadDeadline(zero)
    return true
}

func NewConnection(host string, port int) (*TCPConnection, error) {
    tcpConnection := new(TCPConnection)
    tcpConnection.host = host
    tcpConnection.port = port
    _, err := tcpConnection.Connect()
    if (err != nil) {
        return nil, err
    }
    return tcpConnection, nil
}

