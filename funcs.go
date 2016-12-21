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
		hosts = append(hosts, strings.TrimSpace(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
