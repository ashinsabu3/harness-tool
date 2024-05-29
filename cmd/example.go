package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type exampleOptions struct {
	one bool
	two bool
}

func defaultExampleOptions() *exampleOptions {
	return &exampleOptions{}
}

func newExampleCmd() *cobra.Command {
	o := defaultExampleOptions()

	cmd := &cobra.Command{
		Use:          "helloworld",
		Short:        "helloworld subcommand",
		SilenceUsage: true,
		Args:         cobra.NoArgs,
		RunE:         o.run,
	}

	cmd.Flags().BoolVarP(&o.one, "one", "o", o.one, "one")
	cmd.Flags().BoolVarP(&o.two, "two", "t", o.two, "two")

	return cmd
}

func (o *exampleOptions) run(cmd *cobra.Command, args []string) error {
	if o.one {
		_, err := fmt.Fprintf(cmd.OutOrStdout(), "hello world\n")
		if err != nil {
			return err
		}
	}

	if o.two {
		_, err := fmt.Fprintf(cmd.OutOrStdout(), "hello world 2\n")
		if err != nil {
			return err
		}
	}

	_, err := fmt.Fprintf(cmd.OutOrStdout(), "harness-tool finished.")
	if err != nil {
		return err
	}

	return nil
}

//func (o *exampleOptions) parseArgs(args []string) ([]int, error) {
//	values := make([]int, 2) //nolint: gomnd
//
//	for i, a := range args {
//		v, err := convert.ToInteger(a)
//		if err != nil {
//			return nil, fmt.Errorf("error converting to integer: %w", err)
//		}
//
//		values[i] = v
//	}
//
//	return values, nil
//}
