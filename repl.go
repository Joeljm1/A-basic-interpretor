package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	sc := bufio.NewScanner(os.Stdin)
	for {
		func() {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println("Error:", err)
				}
			}()
			fmt.Print("> ")
			if !sc.Scan() {
				return
			}
			p := NewParser(sc.Text())
			pgm := p.ParseProgram()
			fmt.Printf("%v\n", pgm.Value())
			fmt.Printf("%v\n", pgm.String())
		}()
	}
}
