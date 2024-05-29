package cmd

import (
	"fmt"
	"github.com/ashinsabu/harness-tool/utils"

	"github.com/spf13/cobra"
)

func newRunCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("harness-tool invoked with config file path: %s\n", configFilePath)
			fmt.Println("Config parsed from YAML")
			fmt.Println("---------------------------------------")
			//spew.Dump(parsedConfig)
			fmt.Println(utils.PrettyPrint(parsedConfig))
		},
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return cmd
}
