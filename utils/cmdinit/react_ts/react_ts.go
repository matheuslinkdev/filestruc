package vite

import (
	"fmt"
	"time"

	"github.com/matheuslinkdev/filestruc/utils/dir_utils/react"
	"github.com/spf13/cobra"
)

func NewReactTSCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "react-ts",
		Short: "Creates a folder structure for React + TS",
		Run: func(cmd *cobra.Command, args []string) {
			time.Sleep(2 * time.Second)
			fmt.Println("Creating folder structure for React w/TS")
			time.Sleep(2 * time.Second)
			react.CreateReactStructure()
		},
	}
}
