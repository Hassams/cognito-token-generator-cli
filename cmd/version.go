package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of JWT CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("JWT CLI Tool v1.0.0")
	},
}
