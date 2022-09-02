package app

import (
	"github.com/spf13/cobra"

	"ginkgo/cmd/app/app/options"
)

func NewAPIServerCommand() *cobra.Command {
	o := options.NewAppOptions()
	cmd := &cobra.Command{
		Use:  "app",
		Long: `Long describe.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(o)
		},
	}
	cmd.Flags().StringVar(&o.ConfFile, "c", "", "Config file path.")
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version of KubeSphere ks-apiserver",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("abc")
		},
	}

	cmd.AddCommand(versionCmd)

	return cmd
}

func run(o *options.AppOptions) (err error) {
	server, err := o.NewServer()
	server.Run()
	return
}
