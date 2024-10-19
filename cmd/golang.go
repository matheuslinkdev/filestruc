package cmd

import (
	"fmt"

	"github.com/matheuslinkdev/filestruc/utils"
	"github.com/spf13/cobra"
)

func NewGolangCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "golang",
		Short: "Cria a estrutura básica de um projeto Go",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Criando estrutura para Go...")
			utils.CreateGoStructure()
		},
	}
}
