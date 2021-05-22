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
- Adding much more error handling: the number of `res, _ := FailibileFunction()`
  in the code is probably a bit high for good safety, although I think I avoided
  most obvious possible malicious packets with good length checks

## License

Copyright © 2021 Scott Colby

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the “Software”), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
