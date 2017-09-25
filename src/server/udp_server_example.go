package main

import (
  "fmt"
  "log"
  "net"
  "time"
  "strconv"
)

var (
  msg = "0"
)


func main() {

  go server()

  i := 0
  for {
    msg = strconv.Itoa(i)
    i++
    time.Sleep(time.Second * 1)
  }

}

func server() error {
  ServerAddr, err := net.ResolveUDPAddr("udp",":10001")
  if err != nil { log.Fatal(err) }

  //simple read

  for {
    ServerConn, err := net.ListenUDP("udp", ServerAddr)
    if err != nil { log.Fatal(err) }
    defer ServerConn.Close()

    fmt.Println("lll")

    client(ServerConn)


  }

  // client(address)

  fmt.Println("DuNZO")
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

func client(conn *net.UDPConn) {
  fmt.Println("here")
  var buf [512]byte
  _, addr, err := conn.ReadFromUDP(buf[0:])
  if err != nil { return }

  for {
    _, err := conn.WriteToUDP([]byte(msg), addr)
    if err != nil { log.Fatal(err) }
    time.Sleep(time.Second * 1)
  }
}

// func client(addr string) {
//   ServerAddr,err := net.ResolveUDPAddr("udp",addr)
//   if err != nil { log.Fatal(err) }
//
//   LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:10001")
//   if err != nil { log.Fatal(err) }
//
//   Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
//   if err != nil { log.Fatal(err) }
//
//   defer Conn.Close()
//   i := 0
//   for {
//     msg := strconv.Itoa(i)
//     i++
//     buf := []byte(msg)
//     _,err := Conn.Write(buf)
//     if err != nil {
//         fmt.Println(msg, err)
//     }
//     time.Sleep(time.Second * 1)
//   }
//
// }
