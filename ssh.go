package main

import (
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

	out, err := session.CombinedOutput(command)
	if err != nil {
		result <- "run error: " + string(out)
		return
	}
	result <- string(out)
	return
}

func getSession(m machine) (*ssh.Session, error) {
	conn, err := ssh.Dial("tcp", m.hostname+":"+m.port, m.config)
	if err != nil {
		return nil, err
	}
	return conn.NewSession()
}
