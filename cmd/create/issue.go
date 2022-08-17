package create

import (
	"fmt"

	"github.com/spf13/cobra"
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
		fmt.Println("edit issue called")
	},
}

var (
	Summary     string
	Description string
)

func init() {
	CreateCmd.AddCommand(issueCmd)
	CreateCmd.Flags().StringVarP(&Summary, "summary", "s", "", "a summary of ticket")
	CreateCmd.Flags().StringVarP(&Description, "description", "d", "", "a description of ticket")
	if err := CreateCmd.MarkFlagRequired("summary"); err != nil {
		fmt.Println(err)
	}
	if err := CreateCmd.MarkFlagRequired("description"); err != nil {
		fmt.Println(err)
	}
}
