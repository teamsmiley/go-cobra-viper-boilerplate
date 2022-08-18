/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package configure

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	JIRA_USERNAME string
	JIRA_TOKEN    string
)

// ConfigureCmd represents the configure command
var ConfigureCmd = &cobra.Command{
	Use:   "configure",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configure called")

		viper.Set("JIRA_USERNAME", JIRA_USERNAME)
		viper.Set("JIRA_TOKEN", JIRA_TOKEN)

		error := viper.SafeWriteConfig()
		if error != nil {
			fmt.Println(error)
		}

	},
}

func init() {
	ConfigureCmd.Flags().StringVarP(&JIRA_USERNAME, "jira-username", "u", "", "jira username")
	ConfigureCmd.Flags().StringVarP(&JIRA_TOKEN, "jira-token", "t", "", "jira token")
	if err := ConfigureCmd.MarkFlagRequired("jira-username"); err != nil {
		fmt.Println(err)
	}
	if err := ConfigureCmd.MarkFlagRequired("jira-token"); err != nil {
		fmt.Println(err)
	}
}
