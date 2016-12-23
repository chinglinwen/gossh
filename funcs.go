package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func getenv() {
	if u := os.Getenv("USER"); u != "" {
		user = u
	}
	if p := os.Getenv("PASS"); p != "" {
		pass = p
	}
	if h := os.Getenv("HOST"); h != "" {
		host = h
	}
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func appendList(listfile string) error {
	file, err := os.Open(listfile)
	if err != nil {
		checkErr(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		h := strings.Fields(scanner.Text())
		x := hostem{user: user, pass: pass}
		switch len(h) {
		case 1:
			x.ip = h[0]
		case 2:
			x.ip = h[0]
			x.user = h[1]
		case 3:
			x.ip = h[0]
			x.user = h[1]
			x.pass = h[2]
		default:
		}
		hosts = append(hosts, x)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
