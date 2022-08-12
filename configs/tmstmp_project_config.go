package configs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/joko0811/tmstmp/pkg"
)

const (
	Project_Folder_Name = ".tmstmp"
	Config_File_Name    = "tmstmp_config.json"
)

func GetConfigFilePath() string {
	return pkg.GenerateFilePath([]string{Project_Folder_Name}, Config_File_Name, true)
}

type Config struct {
	Log_Folder_Name  string `json:"log_directory"`
	Log_File_Name    string `json:"log_file_name"`
	Output_File_Path string `json:"output_file_path"`
}

func InitConfig() Config {
	return Config{
		Log_Folder_Name:  "log",
		Log_File_Name:    "timelog.txt",
		Output_File_Path: "./timestamp.txt",
	}
}

func PutDefaultConfigFile(c Config) error {
	// デフォルト設定ファイルの配置

	config_file_path := GetConfigFilePath()

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
	config_file_path := GetConfigFilePath()

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

func GetLogFilePath(c Config) string {
	log_path := pkg.GenerateFilePath([]string{Project_Folder_Name, c.Log_Folder_Name}, c.Log_File_Name, true)

	return log_path
}
