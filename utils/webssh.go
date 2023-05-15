package utils

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/ssh"
)

var SSHConnPool = make(map[string]*ssh.Client)     // 用来存储连接
var SSHSessionPool = make(map[string]*ssh.Session) // 用来存储会话

type RemoteHost struct {
	Host     string
	Port     string
	UserName string
	Password string
}

func (r *RemoteHost) SSHHandler() (*ssh.Client, error) {
	sshConfig := &ssh.ClientConfig{
		User: r.UserName,
		Auth: []ssh.AuthMethod{
			ssh.Password(r.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return ssh.Dial("tcp", fmt.Sprintf("%s:%s", r.Host, r.Port), sshConfig)
}

func (r *RemoteHost) NewSSHSession(key string) (*ssh.Session, error) {
	newKey := r.Host + "_" + r.UserName
	if client, ok := SSHConnPool[newKey]; ok {
		session, err := client.NewSession()
		if err != nil {
			return nil, err
		}
		SSHSessionPool[key] = session
		return session, nil
	}
	client, err := r.SSHHandler()
	if err != nil {
		return nil, err
	}
	SSHConnPool[newKey] = client
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	SSHSessionPool[key] = session
	return session, nil
}

// 执行命令
// 返回: 输出、错误
func SessionDial(session *ssh.Session, command string) (string, string) {
	var stdoutBuf bytes.Buffer
	var stderrBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Stderr = &stderrBuf
	err := session.Run(command)
	if err != nil {
		return "", fmt.Sprintf("错误:%v", err)
	}
	return stdoutBuf.String(), stderrBuf.String()
}
