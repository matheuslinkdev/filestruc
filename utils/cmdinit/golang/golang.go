package golang

import (
	"fmt"
	"time"

	"github.com/matheuslinkdev/filestruc/utils/dir_utils/golang"
	"github.com/spf13/cobra"
)

func NewGolangCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "golang",
		Short: "Creates a folder structure in your golang project, including a entry point",
		Run: func(cmd *cobra.Command, args []string) {
			time.Sleep(2 * time.Second)
			fmt.Println("Creating golang folder structure...")
			time.Sleep(2 * time.Second)
			golang.CreateGoStructure()
		},
	}
}
