// Package cmd
// DateTime: 2023-03-29 17:39
// Author: CN
// Mail: Nanshao@n-s.fun
// Description:

package cmd

import (
	"os"

	"github.com/nanshaobit/nsmail/internal"

	"github.com/spf13/cobra"
)

var Addr string
var Domain string
var ReadTimeout int
var WriteTimeout int
var MaxMessageBytes int
var MaxRecipients int
var AllowInsecureAuth bool

var rootCmd = &cobra.Command{
	Use:   "nsmail",
	Short: "nsmail start script",
	Long:  `nsmail scripts, use --help for more.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		internal.RunServer(Addr, Domain, ReadTimeout, WriteTimeout, MaxMessageBytes, MaxRecipients, AllowInsecureAuth)
		return nil
	},
}

func init() {
	rootCmd.Flags().StringVar(&Addr, "addr", ":25", "addr")
	rootCmd.Flags().StringVar(&Domain, "domain", "localhost", "domain")
	rootCmd.Flags().IntVar(&ReadTimeout, "readTimeout", 10, "read timeout")
	rootCmd.Flags().IntVar(&WriteTimeout, "writeTimeout", 10, "write timeout")
	rootCmd.Flags().IntVar(&MaxMessageBytes, "maxMessage", 1024*1024, "MaxMessageBytes")
	rootCmd.Flags().IntVar(&MaxRecipients, "maxRecipients", 50, "MaxRecipients")
	rootCmd.Flags().BoolVar(&AllowInsecureAuth, "auth", true, "AllowInsecureAuth")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
