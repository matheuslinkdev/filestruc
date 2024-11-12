package nodejs

import (
	"fmt"

	"github.com/matheuslinkdev/filestruc/utils/dir_utils"
	fileutils "github.com/matheuslinkdev/filestruc/utils/file_utils"
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

	fileutils.CreateFile("src/index.js", "js")

	fmt.Println("Node js folder structure created!")
}
