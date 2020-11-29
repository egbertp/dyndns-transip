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
		Use:     "dyndns-transip",
		Short:   "Update ip address on Transip DNS to current public ip ",
		Long:    "Use the current ip to update to a record in the TransIP dns.\nAllowing for easy updating when your ip changes.",
		Version: version,
		Run:     commands.Update,
	}
	userName string
	keyFile  string
	domain   string
	verbose  bool
)

// nolint: gochecknoglobals
var (
	version = "dev"
	commit  = ""
	date    = ""
	builtBy = ""
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

	config.Get().BindPFlag("username", rootCmd.PersistentFlags().Lookup("username")) // nolint: errcheck
	config.Get().BindPFlag("key", rootCmd.PersistentFlags().Lookup("key"))           // nolint: errcheck
	config.Get().BindPFlag("domain", rootCmd.PersistentFlags().Lookup("domain"))     // nolint: errcheck
	config.Get().BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))   // nolint: errcheck
}

func main() {
	rootCmd.Version = buildVersion(version, commit, date, builtBy)

	rootCmd.Execute() // nolint: errcheck
}

func buildVersion(version, commit, date, builtBy string) string {
	var result = version
	if commit != "" {
		result = fmt.Sprintf("%s\ncommit: %s", result, commit)
	}
	if date != "" {
		result = fmt.Sprintf("%s\nbuilt at: %s", result, date)
	}
	if builtBy != "" {
		result = fmt.Sprintf("%s\nbuilt by: %s", result, builtBy)
	}
	return result
}
