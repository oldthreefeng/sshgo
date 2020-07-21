package cmd

import (
	"github.com/oldthreefeng/sshgo/config"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "sshgo",
		Short: "cmd is run ssh cmd",
		Long:  `cmd is run ssh cmd `,
		Run: func(cmd *cobra.Command, args []string) {
			if config.Cmd != "" {
				sshCmd(config.Cmd)
			} else {
				sshTune()
			}
		},
	}
)

// 初始化, 设置 flag 等
func init() {
	rootCmd.PersistentFlags().StringVarP(&config.Cmd, "cmd", "c","", "ssh cmd string")
	rootCmd.PersistentFlags().StringVarP(&config.Host, "host", "H","", "ssh remote host addr ip")
	rootCmd.PersistentFlags().StringVarP(&config.Password, "passwd", "p","", "ssh remote password ")
	rootCmd.PersistentFlags().StringVarP(&config.User, "user", "u","", "ssh remote host user")
	rootCmd.PersistentFlags().StringVarP(&config.PkPath, "pkFile", "","", "private key path")
	rootCmd.PersistentFlags().StringVarP(&config.PkPassword, "pk-passwd", "","", "ssh private key password")
	rootCmd.PersistentFlags().Int64VarP(&config.Port, "port", "P",22, "ssh remote port")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func sshTune() {
	var c config.Conn
	err := c.SetConf()
	if err != nil {
		log.Fatalln(err)
	}
	session, err := c.SetSession()
	if err != nil {
		log.Fatalln(err)
	}
	defer session.Close()
	//当ssh连接建立过后, 我们就可以通过这个连接建立一个回话, 在回话上和远程主机通信。
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.ECHOCTL:       0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	termFd := int(os.Stdin.Fd())
	w, h, _ := terminal.GetSize(termFd)
	termState, _ := terminal.MakeRaw(termFd)
	defer terminal.Restore(termFd, termState)

	err = session.RequestPty("xterm-256color", h, w, modes)
	if err != nil {
		log.Fatalln(err)
	}
	err = session.Shell()
	if err != nil {
		log.Fatalln(err)
	}
	err = session.Wait()
	if err != nil {
		log.Fatalln(err)
	}
}

func sshCmd(cmd string)  {
	var c config.Conn
	err := c.SetConf()
	if err != nil {
		log.Fatalln(err)
	}
	session, err := c.SetSession()
	if err != nil {
		log.Fatalln(err)
	}
	defer session.Close()
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	_ = session.Run(cmd)
}
