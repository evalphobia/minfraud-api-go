package client

import (
	"fmt"

	"github.com/evalphobia/httpwrapper/request"
)

// Client is http client for minFraud API.
type Client struct {
	Option
	MinFraudAccountID  string
	MinFraudLicenseKey string
}

func New() *Client {
	return &Client{}
}

func (c *Client) SetAuthData(accountID, licenseKey string) {
	c.MinFraudAccountID = accountID
	c.MinFraudLicenseKey = licenseKey
}

func (c *Client) SetOption(opt Option) {
	c.Option = opt
}

// CallPOST sends POST request to `url` with `params` and set reqponse to `result`
func (c *Client) CallPOST(path string, params, result interface{}) (err error) {
	opt := c.Option
	url := fmt.Sprintf("%s%s", opt.getBaseURL(), path)

	resp, err := request.POST(url, request.Option{
		Payload:     params,
		PayloadType: request.PayloadTypeJSON,
		User:        c.MinFraudAccountID,
		Pass:        c.MinFraudLicenseKey,
		Retry:       opt.Retry,
		Debug:       opt.Debug,
		UserAgent:   opt.getUserAgent(),
		Timeout:     opt.getTimeout(),
	})
	if err != nil {
		return err
	}
	err = resp.JSON(result)
	return err
}
