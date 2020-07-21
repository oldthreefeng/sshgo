package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"time"
)

type SshConfig struct {
	Host       string
	Port       int64
	User       string
	Cmd        string
	Password   string
	PkPath     string
	PkPassword string
}

type Conn struct {
	user string
	auth []ssh.AuthMethod
	addr string
}

var (
	Host       string
	Port       int64
	User       string
	Cmd        string
	Password   string
	PkPath     string
	PkPassword string
)

func (c *SshConfig) InitConfig() error {
	var path = "./config.toml"
	_, err := toml.DecodeFile(path, c)
	if Host != "" {
		c.Host = Host
	}
	if User != "" {
		c.User = User
	}
	if Port != c.Port {
		c.Port = Port
	}
	if PkPath != "" {
		c.PkPath = PkPath
	}
	if PkPassword != "" {
		c.PkPassword = PkPassword
	}
	return err
}

func (c *Conn) SetConf() (err error) {
	var cfg SshConfig
	err = cfg.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	c.addr = fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	c.auth = make([]ssh.AuthMethod, 0)
	c.user = cfg.User
	var method ssh.AuthMethod
	// use privatakey to login.
	if cfg.PkPath != "" {
		c.auth = make([]ssh.AuthMethod, 0)
		method, err = PublicFile(cfg.PkPath, cfg.PkPassword)
		if err != nil {
			return err
		}
		// use password & user to login
	} else {
		method = ssh.Password(cfg.Password)
	}
	c.auth = append(c.auth, method)
	return nil
}

func (c *Conn) SetSession() (session *ssh.Session, err error) {
	Client, err := ssh.Dial("tcp", c.addr, &ssh.ClientConfig{
		User: c.user,
		Auth: c.auth,
		//需要验证服务端，不做验证返回nil就可以，点击HostKeyCallback看源码就知道了
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: time.Second * 2,
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// create session
	if session, err = Client.NewSession(); err != nil {
		return nil, err
	}

	return session, nil
}
