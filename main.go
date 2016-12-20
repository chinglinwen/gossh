package main

import (
        "bufio"
        "flag"
        "fmt"
        "io/ioutil"
        "log"
        "os"
        "time"

        "github.com/tmc/scp"
        "golang.org/x/crypto/ssh"
)

var example = `
Examples: 
   ./gossh ip command
   echo date | ./gossh ip
   ./gossh -l ip.list command

Use as scp
  ./gossh -c srcfile host targetfile
  ./gossh -l ip.list -c srcfile targetfile

Environment variables:
  USER,PASS,PORT,HOST

`

var (
        user     string
        pass     string
        port     string
        host     string
        hosts    []string
        listfile string
        scpfile  string
)

type machine struct {
        hostname string
        port     string
        config   *ssh.ClientConfig
}

func main() {
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
        getenv()

        if *version {
                fmt.Println("version=1.0.0, 2016-12-20")
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
        var command string
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

        results := make(chan string, len(hosts))
        timeout := time.After(10 * time.Second)

        config := &ssh.ClientConfig{
                User: user,
                Auth: []ssh.AuthMethod{ssh.Password(pass)},
        }

        for _, hostname := range hosts {
                m := machine{hostname, port, config}
                if scpfile != "" {
                        var scptarget string
                        if listfile == "" {
                                scptarget = args[1]
                        } else {
                                scptarget = args[0]
                        }
                        go doscp(m, scpfile, scptarget, results)
                } else {
                        go dossh(m, command, results)
                }
        }

        for i := 0; i < len(hosts); i++ {
                select {
                case res := <-results:
                        fmt.Print(res)
                case <-timeout:
                        fmt.Println("Timed out!")
                        return
                }
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
                hosts = append(hosts, scanner.Text())
        }
        if err := scanner.Err(); err != nil {
                return err
        }
        return nil
}

func doscp(m machine, src, dest string, results chan string) {
        session, err := getSession(m)
        checkErr(err)
        err = scp.CopyPath(src, dest, session)
        if err != nil {
                results <- "scp" + err.Error() + "\n"
                return
        }
        results <- "scp ok\n"
        return
}

func dossh(m machine, cmd string, results chan string) {
        session, err := getSession(m)
        checkErr(err)
        defer session.Close()

        var result string
        out, err := session.CombinedOutput(cmd)
        if err != nil {
                result = m.hostname + " error: " + string(out)
        } else {
                result = string(out)
        }
        results <- result
        return
}

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

func getSession(m machine) (*ssh.Session, error) {
        conn, err := ssh.Dial("tcp", m.hostname+":"+m.port, m.config)
        checkErr(err)
        return conn.NewSession()
}

func checkErr(err error) {
        if err != nil {
                log.Fatal(err)
        }
}