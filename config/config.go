package config

import (
	"os"
	"time"

	"github.com/evalphobia/minfraud-api-go/client"
)

const (
	defaultEnvAccountID  = "MINFRAUD_ACCOUNT_ID"
	defaultEnvLicenseKey = "MINFRAUD_LICENSE_KEY"
)

var (
	envAccountID  string
	envLicenseKey string
)

func init() {
	envAccountID = os.Getenv(defaultEnvAccountID)
	envLicenseKey = os.Getenv(defaultEnvLicenseKey)
}

// Config contains parameters for minFraud.
type Config struct {
	AccountID  string
	LicenseKey string

	Debug   bool
	Timeout time.Duration
}

// Client generates client.Client from Config data.
func (c Config) Client() (*client.Client, error) {
	cli := client.New()
	cli.SetOption(client.Option{
		Debug:   c.Debug,
		Timeout: c.Timeout,
	})
	cli.SetAuthData(c.getAuthData())
	return cli, nil
}

// getAuthData returns MaxMind's AccountID and LicenseKey from Config data or Environment variables.
func (c Config) getAuthData() (accountID, licenseKey string) {
	accountID = c.AccountID
	if accountID == "" {
		accountID = envAccountID
	}

	licenseKey = c.LicenseKey
	if licenseKey == "" {
		licenseKey = envLicenseKey
	}
	return
}
