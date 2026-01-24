package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">")
		if !sc.Scan() {
			return
		}
		p := NewParser(sc.Text())
		pgm := p.ParseProgram()
		println(pgm.String())
	}
}
