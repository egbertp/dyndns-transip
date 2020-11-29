package main

import (
	"fmt"

	"github.com/egbertp/dyndns-transip/internal/commands"
	"github.com/egbertp/dyndns-transip/internal/config"
	"github.com/egbertp/dyndns-transip/internal/logger"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "dyndns",
		Short: "Update ip address on Transip DNS to current public ip ",
		Long:  "Use the current ip to update to a record in the TransIP dns.\nAllowing for easy updating when your ip changes.",
		Run:   commands.Update,
	}
	userName string
	keyFile  string
	domain   string
	verbose  bool
)

// Variables to identify the build
var (
	CommitHash         string
	ApplicationVersion string
	BuildTime          string
)

func init() {
	logger.Init()
	config.Init()
	rootCmd.PersistentFlags().StringVarP(&userName, "username", "u", "", "Transip username")
	rootCmd.PersistentFlags().StringVarP(&keyFile, "key", "k", "", "Transip password key file")
	rootCmd.PersistentFlags().StringVarP(&domain, "domain", "d", "", "The domain (A|AAAA record) for which the ip must be set. (including optional subdomain)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Log level verbose")

	rootCmd.AddCommand(&cobra.Command{
		Use:     "validate",
		Short:   "Validate the the setup",
		Long:    "Run validation to verify the setup is correct",
		Example: "",
		Run:     commands.Validate,
	})
	rootCmd.AddCommand(&cobra.Command{
		Use:     "create",
		Short:   "One time create record for updating",
		Long:    "Create a record for this configuration",
		Example: "",
		Run:     commands.Create,
	})
	rootCmd.AddCommand(&cobra.Command{
		Use:     "update",
		Short:   "Update ip address on Transip DNS to current public ip (default command)",
		Long:    "Use the current ip to update to a record in the TransIP dns.\nAllowing for easy updating when your ip changes.",
		Example: "",
		Run:     commands.Update,
	})
	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print the version number of dyndns-transip",
		// Long:  `All software has versions. This is dyndns-transip's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("dyndns-transip version: %s, commit: %s, build time: %s\n", ApplicationVersion, CommitHash, BuildTime)
		},
	})

	config.Get().BindPFlag("username", rootCmd.PersistentFlags().Lookup("username")) // nolint: errcheck
	config.Get().BindPFlag("key", rootCmd.PersistentFlags().Lookup("key"))           // nolint: errcheck
	config.Get().BindPFlag("domain", rootCmd.PersistentFlags().Lookup("domain"))     // nolint: errcheck
	config.Get().BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))   // nolint: errcheck
}

func main() {
	rootCmd.Execute() // nolint: errcheck
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}
