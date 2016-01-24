package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func Intro() {
	//data, _ := ioutil.ReadFile("./intro.txt")

	//fmt.Println(string(data))

	file, _ := os.Open("./intro.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		sleep := time.Duration(rand.Intn(2)) * time.Second
		time.Sleep(sleep)
	}
}

func Divider() {
	data, _ := ioutil.ReadFile("./divider.txt")

	fmt.Println(string(data))
}
