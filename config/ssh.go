package config

import (
	"errors"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"os"
)

func IsFile(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

//采用公钥验证,这里封装了一下,使用秘钥+密码验证
func PublicFile(privateKeyPath, password string) (method ssh.AuthMethod, err error) {
	if !IsFile(privateKeyPath) {
		return nil, errors.New("file not exist")
	}
	bufKey, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return nil, err
	}
	var key ssh.Signer
	if password == "" {
		key, err = ssh.ParsePrivateKey(bufKey)
		if err != nil {
			return nil, err
		}
	} else {
		bufPwd := []byte(password)
		key, err = ssh.ParsePrivateKeyWithPassphrase(bufKey, bufPwd)
		if err != nil {
			return nil, err
		}
	}
	return ssh.PublicKeys(key), nil
}

