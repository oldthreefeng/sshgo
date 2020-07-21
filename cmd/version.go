package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show sshgo version",
	Long:  `show sshgo version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(VersionStr)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

}

var (
	Version    = "latest"
	Build      = ""
	BuildTime  = ""
	VersionStr = fmt.Sprintf("sshgo version %v, build %v %v, Build Time : %v", Version, Build, runtime.Version(), BuildTime)
)
