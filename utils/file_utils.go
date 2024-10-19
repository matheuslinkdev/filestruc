package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFile(path, ext string) {
	filePath := filepath.Join(path)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Erro ao criar arquivo %s: %v\n", filePath, err)
		return
	}
	defer file.Close()

	content := fmt.Sprintf("// Este é um arquivo %s criado pelo filestruc cli \n", ext)
	file.WriteString(content)

	fmt.Printf("Arquivo criado: %s\n", filePath)
}
