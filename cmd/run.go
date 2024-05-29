package cmd

import (
	"github.com/spf13/cobra"
)

type pipeline struct {
	pipelineId      string
	localInputSetId string
}

func newRunCmd() *cobra.Command {

	pipelineToBeRun := &pipeline{}

	cmd := &cobra.Command{
		Use:   "run",
		Short: "",
		Long:  ``,
		Run:   pipelineToBeRun.run,
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	cmd.Flags().StringVar(&pipelineToBeRun.pipelineId, "id", "", "Pipeline Identifier")
	cmd.Flags().StringVar(&pipelineToBeRun.localInputSetId, "inputsetid", "", "Local inputset id from config yaml provided")
	return cmd
}

func (*pipeline) run(cmd *cobra.Command, args []string) {

}
