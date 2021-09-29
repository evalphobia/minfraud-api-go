package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/mkideal/cli"

	"github.com/evalphobia/minfraud-api-go/config"
	"github.com/evalphobia/minfraud-api-go/minfraud"
)

var outputHeader = []string{
	"ip_address",
	"email",
	"risk_score",
	"ip_risk",
	"ip_is_anonymous",
	"ip_is_anonymous_vpn",
	"ip_is_hosting_provider",
	"ip_is_public_proxy",
	"ip_is_residential_proxy",
	"ip_is_tor_exit_node",
	"ip_organization",
	"ip_user_count",
	"ip_user_type",
	"ip_country",
	"ip_city",
	"ip_registered_country",
	"ip_represented_country",
	"email_domain_first_seen",
	"email_first_seen",
	"email_is_disposable",
	"email_is_free",
	"email_is_high_risk",
}

// parameters of 'list' command.
type listT struct {
	cli.Helper
	Command  string `cli:"*c,command" usage:"set type of api [score, insights, factors] --command='score'"`
	InputCSV string `cli:"*i,input" usage:"input csv/tsv file path --input='./input.csv'"`
	Output   string `cli:"*o,output" usage:"output tsv file path --output='./output.tsv'"`
	Debug    bool   `cli:"debug" usage:"set if you use HTTP debug feature --debug"`
}

func (a *listT) Validate(ctx *cli.Context) error {
	if _, ok := validCommands[a.Command]; !ok {
		keys := make([]string, 0, len(validCommands))
		for k := range validCommands {
			keys = append(keys, k)
		}
		return fmt.Errorf("command should be one of the %v", keys)
	}
	return nil
}

var listC = &cli.Command{
	Name: "list",
	Desc: "Exec api call for minFraud API from csv list file",
	Argv: func() interface{} { return new(listT) },
	Fn:   execList,
}

func execList(ctx *cli.Context) error {
	argv := ctx.Argv().(*listT)

	r := newListRunner(*argv)
	return r.Run()
}

type ListRunner struct {
	// parameters
	Command  string
	InputCSV string
	Output   string
	Debug    bool
}

func newListRunner(p listT) ListRunner {
	return ListRunner{
		Command:  p.Command,
		InputCSV: p.InputCSV,
		Output:   p.Output,
		Debug:    p.Debug,
	}
}

func (r *ListRunner) Run() error {
	f, err := NewCSVHandler(r.InputCSV)
	if err != nil {
		return err
	}

	w, err := NewFileHandler(r.Output)
	if err != nil {
		return err
	}

	lines, err := f.ReadAll()
	if err != nil {
		return err
	}

	maxReq := make(chan struct{}, 2)

	conf := config.Config{
		Debug: r.Debug,
	}

	svc, err := minfraud.New(conf)
	if err != nil {
		return err
	}

	result := make([]string, len(lines))
	var wg sync.WaitGroup
	for i, line := range lines {
		wg.Add(1)
		go func(i int, line map[string]string) {
			maxReq <- struct{}{}
			defer func() {
				<-maxReq
				wg.Done()
			}()

			fmt.Printf("exec #: [%d]\n", i)
			row, err := r.execAPI(svc, line)
			if err != nil {
				fmt.Printf("[ERROR] #: [%d]; err=[%v]\n", i, err)
				return
			}
			result[i] = strings.Join(row, "\t")
		}(i, line)
	}
	wg.Wait()

	result = append([]string{strings.Join(outputHeader, "\t")}, result...)
	return w.WriteAll(result)
}

func (r *ListRunner) execAPI(svc *minfraud.MinFraud, param map[string]string) ([]string, error) {
	params := minfraud.BaseRequest{
		Device: &minfraud.DeviceData{},
		Email:  &minfraud.EmailData{},
	}
	if v, ok := param["ip_address"]; ok {
		params.Device.IPAddress = v
	}
	if v, ok := param["email"]; ok {
		params.Email.Address = v
	}
	resp, err := execAPI(svc, params, r.Command)
	if err != nil {
		return nil, err
	}
	if resp.HasError() {
		return nil, fmt.Errorf("API Error: [%s] [%s]", resp.ErrData.Code, resp.ErrData.Error)
	}

	row := make([]string, 0, len(outputHeader))
	for _, v := range outputHeader {
		row = append(row, getValue(param, resp, v))
	}
	return row, nil
}

func getValue(param map[string]string, resp minfraud.APIResponse, name string) string {
	switch name {
	case "ip_address":
		return param["ip_address"]
	case "email":
		return param["email"]
	case "risk_score":
		return strconv.FormatFloat(resp.RiskScore, 'f', 5, 64)
	case "ip_risk":
		return strconv.FormatFloat(resp.IPAddress.Risk, 'f', 5, 64)
	case "ip_is_anonymous":
		return strconv.FormatBool(resp.IPAddress.Traits.IsAnonymous)
	case "ip_is_anonymous_vpn":
		return strconv.FormatBool(resp.IPAddress.Traits.IsAnonymousVPN)
	case "ip_is_hosting_provider":
		return strconv.FormatBool(resp.IPAddress.Traits.IsHostingProvider)
	case "ip_is_public_proxy":
		return strconv.FormatBool(resp.IPAddress.Traits.IsPublicProxy)
	case "ip_is_residential_proxy":
		return strconv.FormatBool(resp.IPAddress.Traits.IsResidentialProxy)
	case "ip_is_tor_exit_node":
		return strconv.FormatBool(resp.IPAddress.Traits.IsTorExitNode)
	case "ip_organization":
		return resp.IPAddress.Traits.Organization
	case "ip_user_count":
		return strconv.FormatInt(resp.IPAddress.Traits.UserCount, 10)
	case "ip_user_type":
		return resp.IPAddress.Traits.UserType
	case "ip_country":
		return resp.IPAddress.Country.ISOCode
	case "ip_city":
		return resp.IPAddress.City.Names.EN
	case "ip_registered_country":
		return resp.IPAddress.RegisteredCountry.ISOCode
	case "ip_represented_country":
		return resp.IPAddress.RepresentedCountry.ISOCode
	case "email_domain_first_seen":
		return resp.Email.Domain.FirstSeen
	case "email_first_seen":
		return resp.Email.FirstSeen
	case "email_is_disposable":
		return strconv.FormatBool(resp.Email.IsDisposable)
	case "email_is_free":
		return strconv.FormatBool(resp.Email.IsFree)
	case "email_is_high_risk":
		return strconv.FormatBool(resp.Email.IsHighRisk)
	}
	return ""
}
