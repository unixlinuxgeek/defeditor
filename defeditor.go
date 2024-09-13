// File utilities for linux systems

package defeditor

import (
	"fmt"
	"os"
	"os/exec"
)

// Open creates a temporary file and opens your preferred graphical file editor.
// (in ubuntu it is "Image viewer".)
//
// Open создаёт временный файл и открывает предпочтительный файловый редактор
// (в ubuntu это "Image viewer".)
func Open(name string) (string, error) {
	tmpFile, err := os.Create(os.TempDir() + "/" + name)
	if err != nil {
		os.Exit(1)
	}
	path, err := exec.LookPath("xdg-open")
	fmt.Printf("%s\n", path)

	wd, _ := os.Getwd()
	cmd := exec.Command(path, tmpFile.Name())
	err = cmd.Run()
	if err != nil {
		fmt.Printf("%s\n", err)
		return "", err
	}
	return wd + tmpFile.Name(), nil
}
