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

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "timestamp",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		message, _ := cmd.Flags().GetString("message")
		if message != "" {

			err := log_append(message)
			if err != nil {
				log.Fatal(err)
			}

		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.timestamp.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("message", "m", "", "打刻するメッセージを指定してください")
}

func log_append(message string) error {

	message = pkg.Generate_logfmt(message)

	c, err := configs.ReadConfigFile()
	if err != nil {
		return err
	}
	log_dir := c.Log_Directory + "time.log"

	err = pkg.File_append(log_dir, message+"\n")
	if err != nil {
		return err
	}

	return nil
}
