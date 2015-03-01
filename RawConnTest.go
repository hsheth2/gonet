package main

import (
    "fmt"
//    "net/ipv4"
    "net"
    "golang.org/x/net/ipv4"
    "os"
)

func main() {
    fmt.Println("Hello, World!")

    c, err := openUDP(5049, 5050)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    readUDP(c, 1024)
    writeUDP(c, make([]byte, 0))
    closeUDP(c)
}

type UDP struct {
    open bool
	conn *ipv4.RawConn

    src, dest uint16

    pl net.PacketConn
}
func openUDP(src, dest uint16) (*UDP, error) {
    p, err := net.ListenPacket(fmt.Sprintf("ip4:%d", dest), "127.0.0.1")
    if err != nil {
        return nil, err;
    }

    r, err := ipv4.NewRawConn(p)
    if  err != nil {
        return nil, err;
    }

    return &UDP{open: true, conn: r, src: src, dest: dest, pl: p}, nil
}
func readUDP(c *UDP, size int) ([]byte, error) {
    return make([]byte, 0), nil
}
func writeUDP(c *UDP, x []byte) error {
    return nil
}
func closeUDP(c *UDP) error {
    return nil
}