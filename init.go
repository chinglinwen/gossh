package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/ssh"
)

var (
	user      string
	pass      string
	port      string
	host      string
	hosts     []string
	listfile  string
	scpfile   string
	scptarget string

	command string
)

type machine struct {
	hostname string
	port     string
	config   *ssh.ClientConfig
}

func init() {
	getenv()
	flag.Usage = func() {
		fmt.Println("Usage of gossh:")
		flag.PrintDefaults()
		fmt.Printf(example)
	}
	version := flag.Bool("v", false, "show version.")

	flag.StringVar(&user, "u", "", "user name")
	flag.StringVar(&pass, "p", "", "password")
	flag.StringVar(&port, "port", "22", "ssh port")
	flag.StringVar(&listfile, "l", "", "list file of hosts")
	flag.StringVar(&scpfile, "c", "", "scp file to copy")
	flag.Parse()

	if *version {
		fmt.Println("version=1.0.1, 2016-12-21")
		os.Exit(1)
	}

	if listfile != "" {
		checkErr(appendList(listfile))
	}

	args := flag.Args()
	if host == "" && len(hosts) == 0 {
		if len(args) < 1 {
			fmt.Println("no host or cmd been specified")
			os.Exit(1)
		}
	}

	//exmaple: gossh ip command
	//so gossh -l ip.list command, will not append host
	if listfile == "" {
		host = args[0]
		hosts = append(hosts, host)
	}

	fi, err := os.Stdin.Stat()
	checkErr(err)
	if listfile == "" && len(args) < 2 && (fi.Mode()&os.ModeCharDevice) != 0 {
		fmt.Println("no commands provided")
		os.Exit(1)
	}

	//in case ip is hide into list file
	if listfile == "" && len(args) == 2 {
		command = args[1]
	} else if len(args) == 1 {
		command = args[0]
	}

	if (fi.Mode() & os.ModeCharDevice) == 0 {
		b, err := ioutil.ReadAll(os.Stdin)
		checkErr(err)
		command = string(b)
	}

	if listfile == "" {
		scptarget = args[1]
	} else {
		scptarget = args[0]
	}
}
