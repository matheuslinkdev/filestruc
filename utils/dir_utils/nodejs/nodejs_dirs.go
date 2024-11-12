package nodejs

import (
	"fmt"

	"github.com/matheuslinkdev/filestruc/utils/dir_utils"
)

func CreateNodeStructure() {
	directories := []string{
		"public/images",
		"public/icons",
		"src/controllers",
		"src/middlewares",
		"src/routes",
		"src/services",
		"src/models",
		"src/utils",
		"src/config",
		"src/tests/unit",
		"src/tests/integration",
		"src/logs",
		"src/scripts",
	}

	dir_utils.CreateDirectories(directories)

	fmt.Println("Node js folder structure created!")
}
