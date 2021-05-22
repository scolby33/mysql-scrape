# mysql-scrape

Connect to a host and port and print information about a running MySQL server
there.

## Usage

```console
$ git clone https://github.com/scolby33/mysql-scrape
$ cd mysql-scrape
$ go build
$ ./mysql-scrape localhost:3306
```

Output is written as key-value pairs separated by a single tab character, except
for the Status Flags and Capability Flags, which have a header line ending with
a single `:`, followed by the set of flags, one per line, with a single leading
tab character.

If nothing is listening, or something other than a MySQL server is listening,
the program will exit with a non-zero code.

### Example Output

```console
$ docker run --rm -d -e MYSQL_RANDOM_ROOT_PASSWORD=1 -p 127.0.0.1:3306:3306/tcp -- mysql:latest
$ ./mysql-scrape localhost:3306
Server Version  8.0.25
Connection ID   0x8
Character Set   utf8mb4
Status Flags:
        SERVER_STATUS_AUTOCOMMIT
Capability Flags:
        CLIENT_LONG_PASSWORD
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
        CapabilityFlag(33554432)
        CapabilityFlag(67108864)
        CapabilityFlag(134217728)
        CapabilityFlag(1073741824)
        CapabilityFlag(2147483648)
Auth Plugin Data        [74 10 119 86 51 38 107 47 25 64 95 14 23 34 42 79 88 113 121 42 0]
Auth Plugin Name        caching_sha2_password
```

If there is nothing running at the provided port:

```console
$ ./mysql-scrape localhost:1234
dial tcp [::1]:1234: connect: connection refused
$ echo $?
68
```

## Possible Improvements

This program was written with the minimum necessary functionality to demonstrate
the idea of scraping a MySQL server from the Initial Handshake Packet. It is
also my first non-trivial Go program. As a result, it is not very flexible and
may contain bugs or egregious style violations.

Some possible improvements would be:

- Improving package layout by splitting functions out of the `main.go` file
- Adding tests for the three implemented packet parsers
- Parsing the packets to a data structure instead of just outputting information
  to stdout
- Implementing multiple output formats, such as JSON, to allow downstream
  automated consumption of the data
- Allowing multiple arguments to be passed to scrape multiple possible servers
  in parallel, or even a flexible target specification, perhaps similar to nmap,
  for scraping a large number of servers at a time
- Adding much more error handling: the number of `res, _ := FallibleFunction()`
  in the code is probably a bit high for good safety, although I think I avoided
  most obvious possible malicious packets with good length checks

## Reference Material

The documentation I consulted while building this was mainly:

- https://dev.mysql.com/doc/internals/en/mysql-packet.html
- https://dev.mysql.com/doc/internals/en/describing-packets.html
- https://dev.mysql.com/doc/internals/en/connection-phase.html
- https://dev.mysql.com/doc/internals/en/connection-phase-packets.html#packet-Protocol::Handshake
- https://dev.mysql.com/doc/internals/en/packet-ERR_Packet.html
