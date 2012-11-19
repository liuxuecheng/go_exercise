/* Mimics the UNIX cat command */
package main

import(
	"io"
	"os"
	"fmt"
	"bufio"
	"flag"
	)

var numberFlag = flag.Bool("n",false,"number each line")

func cat(r *bufio.Reader){
	i := 1
	for{
		buf, e:= r.ReadBytes('\n')
		if e == io.EOF{
			break
		}
		if *numberFlag{
			fmt.Fprintf(os.Stdout, "%5d %s", i ,buf)
			i++
		}else{
			fmt.Fprintf(os.Stdout,"%s",buf)
		}
	}
	return
}

func main(){
	fname := flag.String("fname", "map.go", "a file name ")
	flag.Parse()
	f, e := os.Open(*fname)
	if e != nil {
		fmt.Fprintf(os.Stderr, "error reading from %s\n")
	}
	cat(bufio.NewReader(f))
}
