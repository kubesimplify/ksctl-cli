package cmd

// authors Dipankar <dipankar@dipankar-das.com>

import (
	"os"

	control_pkg "github.com/kubesimplify/ksctl/pkg/controllers"
	"github.com/kubesimplify/ksctl/pkg/utils/consts"
	"github.com/spf13/cobra"
)

var switchCluster = &cobra.Command{
	Use:     "switch-cluster",
	Aliases: []string{"switch"},
	Short:   "Use to switch between clusters",
	Long: `It is used to switch cluster with the given ClusterName from user. For example:

ksctl switch-context -p <civo,local,civo-ha,azure-ha,azure>  -n <clustername> -r <region> <arguments to civo cloud provider>
`,
	Run: func(cmd *cobra.Command, args []string) {
		verbosity, _ := cmd.Flags().GetInt("verbose")
		if err := control_pkg.InitializeStorageFactory(&cli.Client); err != nil {
			log.Error("Inialize Storage Driver", "Reason", err)
		}
		SetRequiredFeatureFlags(cmd)
		cli.Client.Metadata.ClusterName = clusterName
		cli.Client.Metadata.Region = region
		cli.Client.Metadata.LogVerbosity = verbosity
		cli.Client.Metadata.LogWritter = os.Stdout

		switch provider {
		case string(consts.CloudLocal):
			cli.Client.Metadata.Provider = consts.CloudLocal

		case string(consts.ClusterTypeHa) + "-" + string(consts.CloudCivo):
			cli.Client.Metadata.Provider = consts.CloudCivo
			cli.Client.Metadata.IsHA = true

		case string(consts.CloudCivo):
			cli.Client.Metadata.Provider = consts.CloudCivo

		case string(consts.ClusterTypeHa) + "-" + string(consts.CloudAzure):
			cli.Client.Metadata.Provider = consts.CloudAzure
			cli.Client.Metadata.IsHA = true

		case string(consts.CloudAzure):
			cli.Client.Metadata.Provider = consts.CloudAzure
		}

		err := controller.SwitchCluster(&cli.Client)
		if err != nil {
			log.Error("Switch cluster failed", "Reason", err)
			os.Exit(1)
		}
		log.Success("Switch cluster Successful")
	},
}

func init() {
	rootCmd.AddCommand(switchCluster)
	clusterNameFlag(switchCluster)
	regionFlag(switchCluster)

	switchCluster.Flags().StringVarP(&provider, "provider", "p", "", "Provider")
	switchCluster.MarkFlagRequired("name")
	switchCluster.MarkFlagRequired("provider")
}
