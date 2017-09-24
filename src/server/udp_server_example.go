package main

import (
  "fmt"
  "log"
  "net"
  "time"
  "strconv"
)

func main() {

  server()

}

func server() error {
  ServerAddr, err := net.ResolveUDPAddr("udp",":10001")
  if err != nil {
  	log.Fatal(err)
  }

  //simple read
  ServerConn, err := net.ListenUDP("udp", ServerAddr)
  if err != nil {
  	log.Fatal(err)
  }
  defer ServerConn.Close()

  buf := make([]byte, 1024)

  n,addr,err := ServerConn.ReadFromUDP(buf)

  fmt.Println(n, "->", addr)

  if err != nil { log.Fatal(err) }

  address := string(addr.IP) + ":" + string(addr.Port)


  client(address)


  return nil
  // for {
  //
  //   fmt.Println("Received ",string(buf[0:n]), " from ",addr)
  //   ServerConn.Write([]byte("0"))
  //   if err != nil {
  //       fmt.Println("Error: ",err)
  //   }
  // }
}

func client(addr string) {
  ServerAddr,err := net.ResolveUDPAddr("udp",addr)
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
