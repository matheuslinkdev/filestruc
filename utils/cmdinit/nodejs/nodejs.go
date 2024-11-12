package nodejs

import (
	"fmt"
	"time"

	"github.com/matheuslinkdev/filestruc/utils/dir_utils/nodejs"
	"github.com/matheuslinkdev/filestruc/utils/runcmd"
	"github.com/spf13/cobra"
)

func NewNodeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "node-js",
		Short: "Creates a folder structure for Node JS",
		Run: func(cmd *cobra.Command, args []string) {
			runcmd.RunCommand("npm", "init", "-y")
			time.Sleep(2 * time.Second)
			fmt.Println("Creating Node JS folder structure...")
			time.Sleep(2 * time.Second)
			nodejs.CreateNodeStructure()
		},
	}
}
