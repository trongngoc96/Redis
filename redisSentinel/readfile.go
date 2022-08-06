package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func ChangeConfig(addr) {
	data, err := ioutil.ReadFile("config.js")
	if err != nil {
		fmt.Println(err)
	}
	//REDIS_HOST: '10.22.7.107',
	fmt.Print(string(data))

	f, err := os.OpenFile("config.js", os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	re := regexp.MustCompile(`REDIS_HOST..............'`)
	if _, err = f.WriteString(

		re.ReplaceAllString(string(data), "REDIS_HOST:" + " '"+ addr + "' ")); err != nil {
		panic(err)
	}

	data, err = ioutil.ReadFile("config.js")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(data))
}
