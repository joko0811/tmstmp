package pkg

import (
	"fmt"
	"os"
)

func GenerateFilePath(dir []string, filename string, isRelative bool) string {
	var path string
	if isRelative {
		path = "."
	}
	path += "/"

	for _, d := range dir {
		path += d + "/"
	}

	return path + filename
}

func FileAppend(file_path string, text string) error {

	// 上書きモードで開いて書き込む
	f, err := os.OpenFile(file_path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("file_append: %s", err)
	}

	defer f.Close()
	if _, err = f.WriteString(text); err != nil {
		return fmt.Errorf("file_append: %s", err)
	}

	return nil
}
