package cmd

import (
	"fmt"
	"github.com/ashinsabu/harness-tool/internal/clients"
	"github.com/spf13/cobra"
	"log"
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

func (p *pipeline) run(cmd *cobra.Command, args []string) {
	uuid, err := clients.ExecutePipeline(parsedConfig, p.pipelineId, p.localInputSetId)
	if err != nil {
		log.Fatalf("failed to execute pipeline: %v", err)
	}

	fmt.Printf("Pipeline executed successfully. Plan Execution UUID: %s\n", uuid)

}
