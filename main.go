package main

import (
    "fmt"
    "net"
    "os"
)

// see sysexits(3)
const EX_USAGE int = 64
const EX_NOHOST int = 68
func decodeFixedLengthInteger(data []byte) (uint64, error) {
    l := len(data)
    if !(l == 1 || l == 2 || l == 3 || l == 4 || l == 6 || l == 8) {
        return 0, fmt.Errorf("Invalid length for FixedLengthInteger: %d", l)
    }
    var acc uint64
    for index, value := range data {
        acc += uint64(value) * (1 << (index * 8))
    }
    return acc, nil
}


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
