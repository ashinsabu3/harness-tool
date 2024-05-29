package cmd

import (
	"fmt"
	"github.com/ashinsabu/harness-tool/internal/config_parser"
	"github.com/ashinsabu/harness-tool/utils"

	"github.com/spf13/cobra"
)

var (
	configFilePath string
	parsedConfig   *config_parser.HarnessConfig
	verbose        bool
)

func newRootCmd(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "harness-tool",
		Short: "harness-tool cli tool",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if configFilePath == "" {
				return fmt.Errorf("config file path must be specified")
			}

			var err error
			parsedConfig, err = config_parser.ParseConfig(configFilePath)
			if err != nil {
				return fmt.Errorf("failed to parse config file: %w", err)
			}

			if verbose {
				fmt.Printf("harness-tool invoked with config file path: %s\n", configFilePath)
				fmt.Println("---------------------------------------")
				fmt.Println("Config parsed from YAML")
				fmt.Println("---------------------------------------")
				//spew.Dump(parsedConfig)
				fmt.Println(utils.PrettyPrint(parsedConfig))
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.PersistentFlags().StringVarP(&configFilePath, "config", "c", "", "path to harness_config.yaml")
	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "")

	cmd.AddCommand(newVersionCmd(version)) // version subcommand
	cmd.AddCommand(newExampleCmd())        // example subcommand
	cmd.AddCommand(newRunCmd())

	return cmd
}

// Execute invokes the command.
func Execute(version string) error {
	if err := newRootCmd(version).Execute(); err != nil {
		return fmt.Errorf("error executing root command: %w", err)
	}

	return nil
}
