package main

import (
	"fmt"
	"os"

	"github.com/matheuslinkdev/filestruc/utils/cmdinit/golang"
	"github.com/matheuslinkdev/filestruc/utils/cmdinit/nodejs"
	"github.com/matheuslinkdev/filestruc/utils/cmdinit/react_ts"
	"github.com/spf13/cobra"
)

func main() {

	var rootCmd = &cobra.Command{
		Use:   "filestruc",
		Short: "Helps you to create a folder structure in your project",
		Long: "Helps you to create a complete folder structure in your project, using the language that you've selected",
	}

	rootCmd.AddCommand(golang.NewGolangCmd())
	rootCmd.AddCommand(vite.NewReactTSCmd())
	rootCmd.AddCommand(nodejs.NewNodeCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
