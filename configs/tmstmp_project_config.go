package configs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

const config_file_path string = "tmstmp_config.json"

type Config struct {
	Log_Directory    string `json:"log_directory"`
	Output_File_Path string `json:"output_file_path"`
}

func PutDefaultConfigFile() error {
	// デフォルト設定ファイルの配置

	c := Config{
		Log_Directory:    "./.tmstmp/log/",
		Output_File_Path: "./.tmstmp/timestamp.txt",
	}

	b, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("putDefaultConfigFile: %s", err)
	}

	var content bytes.Buffer
	err = json.Indent(&content, b, "", "  ")
	if err != nil {
		return fmt.Errorf("putDefaultConfigFile: %s", err)
	}

	fp, err := os.Create(config_file_path)
	if err != nil {
		return fmt.Errorf("putDefaultConfigFile: %s", err)
	}
	defer fp.Close()
	fp.WriteString(content.String())

	return nil
}

func ReadConfigFile() (Config, error) {
	var config Config

	b, err := os.ReadFile(config_file_path)
	if err != nil {
		return config, fmt.Errorf("readConfigFile: %s", err)
	}

	err = json.Unmarshal(b, &config)
	if err != nil {
		return config, fmt.Errorf("readConfigFile: %s", err)
	}

	return config, err
}
