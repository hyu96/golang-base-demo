package cmd

import (
	grpc_server "github.com/huydq/product-service/transport/grpc"
	"github.com/spf13/cobra"
)

// httpCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		grpc_server.StartGrpcServer()
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
}
