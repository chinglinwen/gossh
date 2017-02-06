package main

import (
	"bytes"

	"github.com/tmc/scp"
	"golang.org/x/crypto/ssh"
)

func doscp(m machine, result chan string) {
	session, err := getSession(m)
	if err != nil {
		result <- "session error: " + err.Error() + "\n"
		return
	}
	err = scp.CopyPath(scpfile, scptarget, session)
	if err != nil {
		result <- "scp error:" + err.Error() + "\n"
		return
	}
	result <- "scp ok\n"
	return
}

func dossh(m machine, result chan string) {
	session, err := getSession(m)
	if err != nil {
		result <- "session error: " + err.Error() + "\n"
		return
	}
	defer session.Close()

	stdout := new(bytes.Buffer)
	session.Stdout = stdout
	stdin := bytes.NewBufferString(command)
	session.Stdin = stdin

	if err := session.Shell(); err != nil {
		result <- "unable to execute command: %s" + err.Error() + "\n"
		return
	}
	if err := session.Wait(); err != nil {
		result <- "remote command did not exit cleanly: %v" + err.Error() + "\n"
		return
	}

	result <- stdout.String()
	return
}

func getSession(m machine) (*ssh.Session, error) {
	conn, err := ssh.Dial("tcp", m.hostname+":"+m.port, m.config)
	if err != nil {
		return nil, err
	}
	return conn.NewSession()
}
