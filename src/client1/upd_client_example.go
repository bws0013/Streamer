package main

import (
  "fmt"
  "log"
  "net"
)

func main() {
  client()
}

func client() {
  ServerAddr,err := net.ResolveUDPAddr("udp","127.0.0.1:10001")
  if err != nil { log.Fatal(err) }

  // LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
  // if err != nil { log.Fatal(err) }

  Conn, err := net.DialUDP("udp", nil, ServerAddr)
  if err != nil { log.Fatal(err) }

  defer Conn.Close()
  // i := 0
  // for {
  //   msg := strconv.Itoa(i)
  //   i++
  //   buf := []byte(msg)
  //   _,err := Conn.Write(buf)
  //   if err != nil {
  //       fmt.Println(msg, err)
  //   }
  //   time.Sleep(time.Second * 1)
  // }

  _, err = Conn.Write([]byte("anything"))
  if err != nil { log.Fatal(err) }

  var buf [512]byte
  n, err := Conn.Read(buf[0:])
  if err != nil { log.Fatal(err) }

  fmt.Println(string(buf[0:n]))

}
