package cmd

import (
	"fmt"
    "github.com/matheuslinkdev/filestruc/utils"
	"github.com/spf13/cobra"
)

func NewViteTSCmd() *cobra.Command{
    return &cobra.Command{
        Use:   "vite-ts",
		Short: "Cria a estrutura básica de um projeto Vite com TypeScript",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("Criando estrutura para Vite c/ TypeScript")
            utils.CreateViteStructure()
        },
    }
}