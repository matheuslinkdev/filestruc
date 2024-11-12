package react

import (
	"fmt"
	"time"

	"github.com/matheuslinkdev/filestruc/utils/dir_utils"
	fileutils "github.com/matheuslinkdev/filestruc/utils/file_utils"
)

func CreateReactStructure() {
	directories := []string{
		"public/images",
		"public/icons",
		"messages",
		"docs",
		"src/components",
		"src/components/common",
		"src/components/common/NavBar",
		"src/components/common/Footer",
		"src/components/layout",
		"src/components/fragments",
		"src/components/ui",
		"src/components/feedback",
		"src/providers",
		"src/err",
		"src/utils",
		"src/api",
		"src/lib",
		"src/types",
	}

	dir_utils.CreateDirectories(directories)

	fileutils.CreateFile("src/types/global-types.tsx", "tsx")

	time.Sleep(2 * time.Second)
	fmt.Println("React folder structure created!")
}
