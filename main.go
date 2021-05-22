package main

import (
    "fmt"
    "net"
    "os"
)

// see sysexits(3)
const EX_USAGE int = 64
const EX_NOHOST int = 68

func usage() {
}

// Use run function for two reasons:
// 1. Allow an argv to be passed in by a possible external caller
// 2. Return error code ints instead of directly using os.Exit(), ensuring that
//    deferred statements are run
func run(argv []string) int {
    if len(argv) != 2 {
        fmt.Fprint(os.Stderr, `mysql-scrape

Usage:
  mysql-scrape <IP address>:<port>
`)
        return EX_USAGE
    }
    addr := argv[1]

    conn, err := net.Dial("tcp", addr)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return EX_NOHOST
    }
    defer conn.Close()

    fmt.Printf("%s", conn)

    return 0
}

func main() {
    os.Exit(run(os.Args))
}
