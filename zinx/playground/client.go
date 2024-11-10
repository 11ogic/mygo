package main

import (
    "fmt"
    "net"
    "time"
)

func main() {
    fmt.Println("the client trying to connect to the server...")

    conn, err := net.Dial("tcp", "0.0.0.0:8099")
    if err != nil {
        fmt.Println("client connect back error", err)
        return
    }

    for {
        _, err := conn.Write([]byte("Hello v0.1"))
        if err != nil {
            fmt.Println("write conn error", err)
            return
        }
        buf := make([]byte, 512)
        cnt, err := conn.Read(buf)
        if err != nil {
            fmt.Println("read conn error", err)
        }
        fmt.Println(string(buf[:cnt]))

        time.Sleep(1 * time.Second)
    }
}
