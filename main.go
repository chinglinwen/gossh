package main

import (
	"fmt"
	"time"

	"golang.org/x/crypto/ssh"
)

func main() {
	results := make(chan string, len(hosts))

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.Password(pass)},
	}

	for _, hostname := range hosts {
		m := machine{hostname, port, config}
		go do(m, results)
	}

	for i := 0; i < len(hosts); i++ {
		select {
		case res := <-results:
			fmt.Printf("%v", res)
		}
	}
}

func do(m machine, results chan string) {
	result := make(chan string)
	if scpfile != "" {
		go doscp(m, result)
	} else {
		go dossh(m, result)
	}

	select {
	case res := <-result:
		results <- fmt.Sprintf("%v: %v", m.hostname, res)
	case <-time.After(time.Duration(timeout) * time.Second):
		results <- fmt.Sprintf("%v: timed out\n", m.hostname)
	}
	return
}
