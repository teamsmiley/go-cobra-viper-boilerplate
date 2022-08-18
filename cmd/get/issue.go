package get

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// issueCmd represents the issue command
var issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		fmt.Println("get issue called")
		// add
		fmt.Println(viper.GetString("JIRA_USERNAME"))
		fmt.Println(viper.GetString("JIRA_TOKEN"))
	},
}

func init() {
	GetCmd.AddCommand(issueCmd)

}
