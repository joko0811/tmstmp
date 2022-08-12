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

	// 追記モードで開いて書き込む
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

func FileOverwrite(file_path string, text string) error {

	// 上書きする
	f, err := os.Create(file_path)
	if err != nil {
		return fmt.Errorf("file_overwrite: %s", err)
	}

	defer f.Close()

	_, err = f.WriteString(text)
	if err != nil {
		return fmt.Errorf("file_overwrite: %s", err)
	}

	return nil
}
