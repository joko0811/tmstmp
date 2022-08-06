/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/joko0811/tmstmp/configs"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		c := configs.InitConfig()

		project_folder_path := "./" + configs.Project_Folder_Name
		if _, err := os.Stat(project_folder_path); os.IsNotExist(err) {
			os.Mkdir(project_folder_path, 0777)
		}

		err := configs.PutDefaultConfigFile(c)
		if err != nil {
			log.Fatal(err)
		}

		if _, err := os.Stat(c.Log_Folder_Name); os.IsNotExist(err) {
			os.Mkdir(project_folder_path+"/"+c.Log_Folder_Name, 0777)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
