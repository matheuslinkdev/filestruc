package main

import (
	"fmt"
	"os"

	"github.com/matheuslinkdev/filestruc/cmd"
	"github.com/spf13/cobra"
)

func main(){

	var rootCmd = &cobra.Command{
		Use: "filestruc",
		Short: "Helps you to create a file structure in projects using vite or golang",
	}

	rootCmd.AddCommand(cmd.NewGolangCmd())
	rootCmd.AddCommand(cmd.NewViteTSCmd())

	// Executar a CLI
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}