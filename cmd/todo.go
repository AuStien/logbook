package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "Edit the TODO file",
	Run: func(cmd *cobra.Command, args []string) {
		if err := editor.OpenFile(filepath.Join(rootDir, "TODO.md")); err != nil {
			fmt.Fprintf(os.Stderr, "failed editing TODO.md: %s\n", err.Error())
			os.Exit(1)
		}
	},
}
