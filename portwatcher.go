package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s portrange\n\na portrange is defined like: 1,2,5-8,15\n\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

func fatal(message string) {
	fmt.Fprintf(os.Stderr, "error: %s\n\n", message)
	usage()
}

/* Expand a port range into a slice of ports to bind to */
func expand(portrange string) ([]int, error) {
	var result []int

	parts := strings.Split(portrange, ",")
	for _, fragment := range parts {
		fragment = strings.TrimSpace(fragment)
		if strings.Contains(fragment, "-") {
			fromAndTo := strings.Split(fragment, "-")
			from := strings.TrimSpace(fromAndTo[0])
			to := strings.TrimSpace(fromAndTo[1])
			fromI, err := strconv.Atoi(from)
			if err != nil {
				return nil, fmt.Errorf("'%s' could not be converted to a number", from)
			}
			toI, err := strconv.Atoi(to)
			if err != nil {
				return nil, fmt.Errorf("'%s' could not be converted to a number", to)
			}

			for i := fromI; i <= toI; i++ {
				result = append(result, i)
			}
		} else {
			port, err := strconv.Atoi(fragment)
			if err != nil {
				return nil, fmt.Errorf("'%s' could not be converted to a number", fragment)
			}
			result = append(result, port)
		}
	}

	return result, nil
}

func acceptAndPrint(ln net.Listener) {
	conn, err := ln.Accept()
	if err != nil {
		fatal(err.Error())
	}

	addr := conn.LocalAddr().String()

	fmt.Printf("Received connection on %s\n", addr)
}

func main() {
	flag.Usage = usage

	var udp = flag.Bool("udp", false, "Listen on UDP instead of TCP")

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "Port range description missing.")
		usage()
		os.Exit(1)
	}

	ports, err := expand(args[0])
	if err != nil {
		fatal(err.Error())
	}

	if *udp {
	}

	for _, port := range ports {
		var ln net.Listener
		/*
			if *udp {
				addr, err := net.ResolveUDPAddr("udp", ":"+string(port))
				if err != nil {
					fatal(err.Error())
				}
				ln, err = net.ListenUDP("udp", addr)
			} else {
				ln, err = net.Listen("tcp", ":" + string(port))
			}
		*/
		ln, err = net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			fatal(err.Error())
		}

		go acceptAndPrint(ln)
	}

	for {
		time.Sleep(time.Second)
	}
}
