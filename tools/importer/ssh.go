package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

type SshConnection struct {
	client *ssh.Client
}

func NewSsh(host, username, password string) (*SshConnection, error) {
	// var hostKey ssh.PublicKey
	// An SSH client is represented with a ClientConn.
	//
	// To authenticate with the remote server you must pass at least one
	// implementation of AuthMethod via the Auth field in ClientConfig,
	// and provide a HostKeyCallback.
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //ssh.FixedHostKey(hostKey),
		Timeout:         10 * time.Second,
	}
	client, err := ssh.Dial("tcp", host, config)
	if err != nil {
		return nil, fmt.Errorf("failed to dial, %v", err)
	}

	conn := &SshConnection{
		client: client,
	}

	return conn, nil
}

func (c *SshConnection) Close() error {
	return c.client.Close()
}

func (c *SshConnection) Run(cmd string) (string, error) {
	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := c.client.NewSession()
	if err != nil {
		return "", fmt.Errorf("failed to create session, %v", err)
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(cmd); err != nil {
		return "", fmt.Errorf("failed to run, %v", err.Error())
	}
	return b.String(), nil
}

func GetMikrotikConfig(conn *SshConnection) (string, error) {
	return conn.Run("/export terse")
}

func GetResourceId(conn *SshConnection, path string, requiredFields []string) string {
	var id string
	var filter = strings.Join(requiredFields, " ")

	cmd := fmt.Sprintf(":put [%v get [ find %v ]]", path, filter)
	log.Debug("Searching id in ", path, " with command: ", cmd)
	res, err := conn.Run(cmd)
	if err != nil {
		log.Error("Error running command: ", err)
		return "?"
	}

	ss := reId.FindStringSubmatch(res)
	log.Debug("ss is ", ss)
	if len(ss) == 2 {
		id = ss[1]
		log.Info("Found id ", id, " for ", filter)
		return id
	}

	// Let's try with dynamic=no
	cmd = fmt.Sprintf(":put [%v get [ find %v dynamic=no ]]", path, filter)
	log.Debug("Searching id in ", path, " with command: ", cmd)
	res, err = conn.Run(cmd)
	if err != nil {
		log.Error("Error running command: ", err)
		return "?"
	}

	ss = reId.FindStringSubmatch(res)
	log.Debug("ss is ", ss)
	if len(ss) == 2 {
		id = ss[1]
		log.Info("Found id ", id, " for ", filter)
		return id
	}

	// print
	var proplist []string
	for _, v := range requiredFields {
		ss := strings.SplitN(v, "=", 2)
		proplist = append(proplist, ss[0])
	}
	cmd = fmt.Sprintf("%v print show-ids where %v dynamic=no", path, filter)
	log.Debug("Searching id in ", path, " with command: ", cmd)
	res, err = conn.Run(cmd)
	if err != nil {
		log.Error("Error running command: ", err)
		return "?"
	}

	ss = rePrintId.FindStringSubmatch(res)
	log.Debug("ss is ", ss)
	if len(ss) == 2 {
		id = ss[1]
		log.Info("Found id ", id, " for ", filter)
		return id
	}

	if id == "" {
		log.Warnf("Id not found for %v (filter: %v)", requiredFields, filter)
		return "?"
	}

	return id
}
