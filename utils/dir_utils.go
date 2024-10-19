package utils

import (
	"fmt"
	"os"
)

func CreateViteStructure() {
	directories := []string{
		"public/images",
		"public/icons",
		"src/components",
		"src/components/common",
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

	for _, dir := range directories {
		fullPath := "" + dir
		err := os.MkdirAll(fullPath, os.ModePerm)
		if err != nil {
			fmt.Printf("Erro ao criar diretório %s: %v\n", fullPath, err)
		} else {
			fmt.Printf("Diretório criado: %s\n", fullPath)
		}
	}

	fmt.Println("Criando diretório com TypeScript...")
	CreateFile("src/types/global-types.ts", "tsx")

	fmt.Println("Estrutura de projeto Vite criada com sucesso!")
}

func CreateGoStructure() {
	directories := []string{
		"cmd",
		"pkg",
		"internal",
		"api",
		"configs",
		"scripts",
		"build",
		"docs",
	}

	for _, dir := range directories {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			fmt.Printf("Erro ao criar diretório %s: %v\n", dir, err)
		} else {
			fmt.Printf("Diretório criado: %s\n", dir)
		}
	}

	CreateFile("main.go", "go")

	fmt.Println("Estrutura de projeto Go criada com sucesso!")
}
