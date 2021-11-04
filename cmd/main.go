package main

import (
	SearchPath "SearchPath/internal"
	"fmt"
	"os"
)

func main() {
	projectName := "SearchPath"

	dir, _ := os.Getwd()
	fmt.Println(dir)

	fmt.Println(SearchPath.AnyProjectPath(projectName))
}
