package commands

import (
	"github.com/egbertp/dyndns-transip/internal/config"
	"github.com/egbertp/dyndns-transip/internal/gipify"
	"github.com/egbertp/dyndns-transip/internal/logger"
	"github.com/egbertp/dyndns-transip/internal/tld"
	"github.com/spf13/cobra"
)

// Create record with public IP
func Create(cmd *cobra.Command, args []string) {
	logger.SetVerbose(config.Get().GetBool("verbose"))
	IP, err := gipify.GetIP()
	if err != nil {
		logger.Get().Fatalf("Error getting IP address. (%s)", err.Error())
	}
	err = tld.InitTLD(config.Get().GetString("username"), config.Get().GetString("private-key"))
	if err != nil {
		logger.Get().Fatalf("Error accessing the API. please verify configuration (%s)", err.Error())
	}
	tld.SetRecordInformation(
		config.Get().GetString("domain"),
		config.Get().GetString("domain-entry"),
		config.Get().GetInt("domain-ttl"),
	)
	_, err = tld.FindRecord()
	if err == nil {
		logger.Get().Fatalf("Record already exists. Use update from now on.")
	}
	err = tld.CreateRecord(IP)
	if err != nil {
		logger.Get().Fatalf("Unable to create record. (%s)", err.Error())
	} else {
		logger.Get().Infof("Create record for %s.%s with ip %s.", config.Get().GetString("domain-entry"), config.Get().GetString("domain"), IP.IP)
	}
}
