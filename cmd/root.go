package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/austien/logbook/editors"
	"github.com/austien/logbook/journal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	editor  editors.Editor
	rootDir string
)

func init() {
	rootCmd.PersistentFlags().String("home", "$HOME/.logbook", "home for logs (default is $HOME/.logbook)")
	rootCmd.PersistentFlags().String("editor", "vi", "which editor to use (default is vi")

	viper.SetEnvPrefix("LOGBOOK")
	viper.BindPFlag("home", rootCmd.LocalFlags().Lookup("home"))
	viper.BindEnv("home")
	viper.BindPFlag("editor", rootCmd.LocalFlags().Lookup("editor"))
	viper.BindEnv("editor", "EDITOR")

	rootCmd.AddCommand(todoCmd)
	rootCmd.AddCommand(viewCmd)
	rootCmd.AddCommand(binderCmd)
}

var rootCmd = &cobra.Command{
	Use:   "log",
	Short: "Jot jot",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error
		editor, err = editors.GetEditor(viper.GetString("editor"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed getting editor setup for %s: %s\n", viper.GetString("editor"), err.Error())
			os.Exit(1)
		}

		rootDir, err = filepath.Abs(viper.GetString("home"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed getting absolute path for %s: %s\n", viper.GetString("home"), err.Error())
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()

		j := journal.New(rootDir, editor)

		if err := j.CreateEntry(now); err != nil {
			fmt.Fprintf(os.Stderr, "upserDayFile: %s\n", err.Error())
			os.Exit(1)
		}
	},
}

func Execute() error {
	if len(os.Args[1:]) == 1 {
		var cmdFound bool
		cmds := rootCmd.Commands()

		for _, cmd := range cmds {
			if cmd.Name() == os.Args[1] {
				cmdFound = true
				break
			}
		}
		if !cmdFound {
			args := append([]string{"binder"}, os.Args[1])
			rootCmd.SetArgs(args)
		}
	}

	return rootCmd.Execute()
}
