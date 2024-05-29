package cmd

import (
	"fmt"
	"github.com/ashinsabu/harness-tool/internal/config_parser"

	"github.com/spf13/cobra"
)

var (
	configFilePath string
	parsedConfig   *config_parser.HarnessConfig
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
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.PersistentFlags().StringVarP(&configFilePath, "config", "c", "", "path to harness_config.yaml")

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
