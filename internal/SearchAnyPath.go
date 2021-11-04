package SearchPath

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

/* *
*	the path from the top directory to an any directory name.
* 	args:
*		dirName string : The name of the directory you want to search
*	return:
*		path string : the path from the top directory to arg dirName.
* */
func AnyProjectPath(dirName string) string {
	currentDir, _ := os.Getwd()
	fmt.Println(currentDir)

	var dirNameList []string

	// Change the delimiter depending on the OS
	if runtime.GOOS != "windows" {
		dirNameList = strings.Split(currentDir, "/")
		fmt.Println(dirNameList)
	} else {
		dirNameList = strings.Split(currentDir, "\\")
		fmt.Println(dirNameList)
	}

	cnt := 0
	for i := 0; i < len(dirNameList); i++ {
		cnt++
		fmt.Println(dirNameList[i])

		if dirNameList[i] == dirName {
			break
		}
	}

	path := strings.Join(dirNameList[0:cnt], "/")
	return path
}
