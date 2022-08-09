package main

import (
	"flag"
	"fmt"
	"sentinel/sentinel"
	"strings"
)

func main() {
	file := flag.String("file", "", "")
	flag.Parse()
	arrCmd := strings.Split(string(*file), ",")
	sentinel.NewMaster(arrCmd)
	fmt.Scanln()
}
