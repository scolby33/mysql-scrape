package main

import (
    "bytes"
    "fmt"
    "io"
    "net"
    "os"
    "time"
)

// see sysexits(3)
const EX_OK int = 0
const EX_USAGE int = 64
const EX_NOHOST int = 68
const EX_IOERR int = 74
const EX_PROTOCOL int = 76

//go:generate stringer -type=CapabilityFlag
type CapabilityFlag uint64

const (
    CLIENT_LONG_PASSWORD CapabilityFlag = 1 << iota
    CLIENT_FOUND_ROWS
    CLIENT_LONG_FLAG
    CLIENT_CONNECT_WITH_DB
    CLIENT_NO_SCHEMA
    CLIENT_COMPRESS
    CLIENT_ODBC
    CLIENT_LOCAL_FILES
    CLIENT_IGNORE_SPACE
    CLIENT_PROTOCOL_41
    CLIENT_INTERACTIVE
    CLIENT_SSL
    CLIENT_IGNORE_SIGPIPE
    CLIENT_TRANSACTIONS
    CLIENT_RESERVED
    CLIENT_SECURE_CONNECTION
    CLIENT_MULTI_STATEMENTS
    CLIENT_MULTI_RESULTS
    CLIENT_PS_MULTI_RESULTS
    CLIENT_PLUGIN_AUTH
    CLIENT_CONNECT_ATTRS
    CLIENT_PLUGIN_AUTH_LENENC_CLIENT_DATA
    CLIENT_CAN_HANDLE_EXPIRED_PASSWORDS
    CLIENT_SESSION_TRACK
    CLIENT_DEPRECATE_EOF
)

func printCapabilityFlags(flags CapabilityFlag) {
    for i := CapabilityFlag(1); i < (1 << 63); i = i << 1 {
        flag := flags & i
        if flag != 0 {
            fmt.Printf("\t%s\n", flag)
        }
    }
}

//go:generate stringer -type=StatusFlag
type StatusFlag uint16

const (
    SERVER_STATUS_IN_TRANS StatusFlag = 1 << iota
    SERVER_STATUS_AUTOCOMMIT
    SERVER_MORE_RESULTS_EXISTS
    SERVER_STATUS_NO_GOOD_INDEX_USED
    SERVER_STATUS_NO_INDEX_USED
    SERVER_STATUS_CURSOR_EXISTS
    SERVER_STATUS_LAST_ROW_SENT
    SERVER_STATUS_DB_DROPPED
    SERVER_STATUS_NO_BACKSLASH_ESCAPES
    SERVER_STATUS_METADATA_CHANGED
    SERVER_QUERY_WAS_SLOW
    SERVER_PS_OUT_PARAMS
    SERVER_STATUS_IN_TRANS_READONLY
    SERVER_SESSION_STATE_CHANGED
)

func printStatusFlags(flags StatusFlag) {
    for i := StatusFlag(1); i < (1 << 7); i = i << 1 {
        flag := flags & i
        if flag != 0 {
            fmt.Printf("\t%s\n", flag)
        }
    }
}

//go:generate stringer -type=CharacterSet
type CharacterSet byte

const (
    big5     CharacterSet = 1
    dec8     CharacterSet = 3
    cp850    CharacterSet = 4
    hp8      CharacterSet = 6
    koi8r    CharacterSet = 7
    latin1   CharacterSet = 8
    latin2   CharacterSet = 9
    swe7     CharacterSet = 10
    ascii    CharacterSet = 11
    ujis     CharacterSet = 12
    sjis     CharacterSet = 13
    hebrew   CharacterSet = 16
    tis620   CharacterSet = 18
    euckr    CharacterSet = 19
    koi8u    CharacterSet = 22
    gb2312   CharacterSet = 24
    greek    CharacterSet = 25
    cp1250   CharacterSet = 26
    gbk      CharacterSet = 28
    latin5   CharacterSet = 30
    armscii8 CharacterSet = 32
    utf8     CharacterSet = 33
    ucs2     CharacterSet = 35
    cp866    CharacterSet = 36
    keybcs2  CharacterSet = 37
    macce    CharacterSet = 38
    macroman CharacterSet = 39
    cp852    CharacterSet = 40
    latin7   CharacterSet = 41
    cp1251   CharacterSet = 51
    utf16    CharacterSet = 54
    utf16le  CharacterSet = 56
    cp1256   CharacterSet = 57
    cp1257   CharacterSet = 59
    utf32    CharacterSet = 60
    binary   CharacterSet = 63
    geostd8  CharacterSet = 92
    cp932    CharacterSet = 95
    eucjpms  CharacterSet = 97
    gb18030  CharacterSet = 248
    utf8mb4  CharacterSet = 255
)

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

func parseHandshakeV10(data []byte) {
    buf := bytes.NewBuffer(data)

    server_version, _ := buf.ReadString(0x00)
    fmt.Printf("Server Version\t%s\n", server_version)

    connection_id_bytes := buf.Next(4)
    connection_id, _ := decodeFixedLengthInteger(connection_id_bytes)
    fmt.Printf("Connection ID\t0x%x\n", connection_id)

    var auth_plugin_data [21]byte
    _, _ = buf.Read(auth_plugin_data[:8])

    _, _ = buf.ReadByte() // filler byte

    if buf.Len() == 0 {
        // remaining fields are optional, so might end here
        return
    }

    var capability_flags_bytes [4]byte

    _, _ = buf.Read(capability_flags_bytes[:2])

    character_set, _ := decodeFixedLengthInteger(buf.Next(1))
    fmt.Printf("Character Set\t%s\n", CharacterSet(character_set))

    var status_flags_bytes [2]byte
    _, _ = buf.Read(status_flags_bytes[:])
    status_flags, _ := decodeFixedLengthInteger(status_flags_bytes[:])
    fmt.Println("Status Flags:")
    printStatusFlags(StatusFlag(status_flags))

    _, _ = buf.Read(capability_flags_bytes[2:])
    capability_flags, _ := decodeFixedLengthInteger(capability_flags_bytes[:])

    fmt.Println("Capability Flags:")
    printCapabilityFlags(CapabilityFlag(capability_flags))

    auth_plugin_data_length_byte, _ := buf.ReadByte()
    auth_plugin_data_length := int(auth_plugin_data_length_byte)

    _ = buf.Next(10) // reserved, all nulls

    auth_plugin_data_part_2_length := auth_plugin_data_length - 8
    if auth_plugin_data_part_2_length > 13 {
        auth_plugin_data_part_2_length = 13
    }
    auth_plugin_data_length = auth_plugin_data_part_2_length + 8

    _, _ = buf.Read(auth_plugin_data[8:auth_plugin_data_length])
    fmt.Printf("Auth Plugin Data\t%v\n", auth_plugin_data)

    auth_plugin_name, _ := buf.ReadString(0x00)
    fmt.Printf("Auth Plugin Name\t%s\n", auth_plugin_name)
}

func parseHandshakeV9(data []byte) {
    buf := bytes.NewBuffer(data)

    server_version, _ := buf.ReadString(0x00)
    fmt.Printf("Server Version\t%s\n", server_version)

    connection_id_bytes := buf.Next(4)
    connection_id, _ := decodeFixedLengthInteger(connection_id_bytes)
    fmt.Printf("Connection ID\t0x%x\n", connection_id)

    scramble, _ := buf.ReadString(0x00)
    fmt.Printf("Scramble\t0x%x\n", scramble)
}

func parseError(data []byte) {
    buf := bytes.NewBuffer(data)

    error_code := buf.Next(2)
    fmt.Printf("Error Code\t0x%x\n", error_code)

    sql_state_marker, _ := buf.ReadByte()
    if sql_state_marker == 0x23 { // ASCII '#'
        sql_state := buf.Next(5)
        fmt.Printf("SQL State\t0x%x\n", sql_state)
    }

    error_message := buf.String()
    fmt.Printf("Error Message\t%s\n", error_message)
}

// Use run function for two reasons:
// 1. Allow an argv to be passed in by a possible external caller
// 2. Return error code ints instead of directly using os.Exit(), ensuring that
//    deferred statements are run
func run(argv []string) int {
    var n int
    var err error

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
    conn.SetDeadline(time.Now().Add(time.Second * 2))

    var header_buf [4]byte
    n, err = conn.Read(header_buf[:])
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return EX_IOERR
    }
    if n != 4 {
        fmt.Fprintf(os.Stderr, "Too few bytes in packet header: %d", n)
        return EX_PROTOCOL
    }
    payload_length, _ := decodeFixedLengthInteger(header_buf[:3])
    //sequence_number, _ := decodeFixedLengthInteger(header_buf[3:])

    payload_buf := make([]byte, payload_length)
    _, err = io.ReadFull(conn, payload_buf)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return EX_IOERR
    }

    switch magic := payload_buf[0]; magic {
    case 0x0a:
        parseHandshakeV10(payload_buf[1:])
    case 0x09:
        parseHandshakeV9(payload_buf[1:])
    case 0xff:
        parseError(payload_buf[1:])
    default:
        fmt.Fprintf(os.Stderr, "Unknown packet type: %v\n", magic)
        return EX_PROTOCOL
    }

    return EX_OK
}

func main() {
    os.Exit(run(os.Args))
}
