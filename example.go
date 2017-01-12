package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/olebedev/when"
	"github.com/when/rules/common"
	"github.com/when/rules/en"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	fmt.Print("Enter in a time command, like: drop me a line in next wednesday at 2:25 p.m\n")
	text, _ := reader.ReadString('\n')

	fmt.Scanln(text)

	r, err := w.Parse(text, time.Now())
	if err != nil {
		// an error has occurred
	}
	if r == nil {
		// no matches found
	}

	fmt.Println(
		"time was [",
		r.Time.String(),
		"]mentioned in",
		"\""+text+"\"",
	)

}
