package tools

import (
	"github.com/spf13/cobra"
)

func Cmds(rootCmd *cobra.Command) []*cobra.Command {
	return []*cobra.Command{testCmd()}
}
