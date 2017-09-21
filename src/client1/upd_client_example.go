package main

import (
  "fmt"
  "log"
  "net"
  "time"
  "strconv"
)

func main() {
  client()
}

func client() {
  ServerAddr,err := net.ResolveUDPAddr("udp","127.0.0.1:10001")
  if err != nil { log.Fatal(err) }

  LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
  if err != nil { log.Fatal(err) }

  Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
  if err != nil { log.Fatal(err) }

  defer Conn.Close()
  i := 0
  for {
    msg := strconv.Itoa(i)
    i++
    buf := []byte(msg)
    _,err := Conn.Write(buf)
    if err != nil {
        fmt.Println(msg, err)
    }
    time.Sleep(time.Second * 1)
  }

}
