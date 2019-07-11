package main

import (
	"fmt"
	"os"
	"tokigo/toki/tools"

	"github.com/lovego/cmd"
	"github.com/spf13/cobra"
)

const version = `19.07.10`

func main() {

	root := &cobra.Command{
		Use:          `toki`,
		Short:        `tools.`,
		SilenceUsage: true,
	}

	root.AddCommand(tools.Cmds(root)...)
	root.AddCommand(versionCmd(), updateCmd())

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   `version`,
		Short: `show toki version.`,
		RunE: func(c *cobra.Command, args []string) error {
			fmt.Println(`toki version ` + version)
			return nil
		},
	}
}

func updateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   `update`,
		Short: `update to lastest version.`,
		RunE: func(c *cobra.Command, args []string) error {
			fmt.Println(`current version ` + version)
			if _, err := cmd.Run(cmd.O{},
				`go`, `get`, `-u`, `-v`, `github.com/axzhao/toki`,
			); err != nil {
				return err
			}
			_, err := cmd.Run(cmd.O{}, `toki`, `version`)
			return err
		},
	}
}
