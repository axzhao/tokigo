package tools

import (
	"log"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func testCmd() *cobra.Command {
	var env string
	var showBsql bool
	cmd := &cobra.Command{
		Use:   `test`,
		Short: ``,
		RunE: func(c *cobra.Command, args []string) error {
			return Test(showBsql, env, args)
		},
	}
	cmd.Flags().BoolVarP(&showBsql, "bsql", "b", false, "show sqls.")
	cmd.Flags().StringVar(&env, "env", "dev", "GOENV")
	return cmd
}

func Test(showBsql bool, env string, targets []string) error {
	log.Println(color.GreenString(`Run Test Case.`))

	if cwd, err := os.Getwd(); err != nil {
		panic(err)
	} else if err := os.Chdir(cwd); err != nil {
		panic(err)
	}

	os.Setenv(`GOENV`, env)
	if showBsql {
		os.Setenv(`DebugBsql`, `1`)
	}
	c := exec.Command(`go`, append([]string{`test`, `-v`, `-test.run`}, targets...)...)
	f, err := c.Output()
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(f))

	log.Println(``)

	return nil
}
