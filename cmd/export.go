/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"os"

	"github.com/joko0811/tmstmp/configs"
	"github.com/joko0811/tmstmp/pkg"
	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		c, err := configs.ReadConfigFile()
		if err != nil {
			log.Fatal(err)
		}

		log_path := pkg.GenerateFilePath([]string{configs.Project_Folder_Name, c.Log_Folder_Name}, c.Log_File_Name, true)
		b, err := os.ReadFile(log_path)

		s := "\n" + string(b)

		err = pkg.FileAppend(c.Output_File_Path, s)
		if err != nil {
			log.Fatal(err)
		}

		// ファイルを空にする
		err = pkg.FileOverwrite(log_path, "")
		if err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(exportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
