package configs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Log_Directory    string `json:"log_directory"`
	Output_File_Path string `json:"output_file_path"`
}

func PutDefaultConfigFile() error {
	// デフォルト設定ファイルの配置

	c := Config{
		Log_Directory:    "./log/",
		Output_File_Path: "./timestamp.txt",
	}

	b, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("putDefaultConfigFile: %s", err)
	}

	var content bytes.Buffer
	err = json.Indent(&content, b, "", "  ")

	fp, err := os.Create("config.json")
	if err != nil {
		return fmt.Errorf("putDefaultConfigFile: %s", err)
	}
	defer fp.Close()
	fp.WriteString(content.String())

	return nil
}

func ReadConfigFile() (Config, error) {
	var config Config

	b, err := os.ReadFile("./config.json")
	if err != nil {
		return config, fmt.Errorf("readConfigFile: %s", err)
	}

	err = json.Unmarshal(b, &config)
	if err != nil {
		return config, fmt.Errorf("readConfigFile: %s", err)
	}

	return config, err
}
