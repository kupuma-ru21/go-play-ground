package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// unbuffered: os.Stdout implements io.Writer
	fmt.Fprintf(os.Stdout, "%s\n", "hello world! - unbuffered")
	// buffered:
	buf := bufio.NewWriter(os.Stdout)
	// and now so does buf:
	fmt.Fprintf(buf, "%s\n", "hello world! - buffered")
	buf.Flush()
}
