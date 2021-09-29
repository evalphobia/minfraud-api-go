package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mkideal/cli"

	"github.com/evalphobia/minfraud-api-go/config"
	"github.com/evalphobia/minfraud-api-go/minfraud"
)

const (
	commandScore    = "score"
	commandInsights = "insights"
	commandFactors  = "factors"
)

var validCommands = map[string]struct{}{
	commandScore:    {},
	commandInsights: {},
	commandFactors:  {},
}

// parameters of 'single' command.
type singleT struct {
	cli.Helper
	Command   string `cli:"*c,command" usage:"set type of api [score, insights, factors] --command='score'"`
	IPAddress string `cli:"i,ipaddr" usage:"input ip address --ipaddr='8.8.8.8'"`
	Email     string `cli:"e,email" usage:"input email address --email='example@example.com'"`
	Debug     bool   `cli:"debug" usage:"set if you use HTTP debug feature --debug"`
}

func (a *singleT) Validate(ctx *cli.Context) error {
	if _, ok := validCommands[a.Command]; !ok {
		keys := make([]string, 0, len(validCommands))
		for k := range validCommands {
			keys = append(keys, k)
		}
		return fmt.Errorf("command should be one of the %v", keys)
	}
	if a.IPAddress == "" && a.Email == "" {
		return errors.New("you must set --ipaddr or --email")
	}
	return nil
}

var singleC = &cli.Command{
	Name: "single",
	Desc: "Exec single api call for minFraud API",
	Argv: func() interface{} { return new(singleT) },
	Fn:   execSingle,
}

func execSingle(ctx *cli.Context) error {
	argv := ctx.Argv().(*singleT)

	r := newSingleRunner(*argv)
	return r.Run()
}

type SingleRunner struct {
	// parameters
	Command   string
	IPAddress string
	Email     string
	Debug     bool
}

func newSingleRunner(p singleT) SingleRunner {
	return SingleRunner{
		Command:   p.Command,
		IPAddress: p.IPAddress,
		Email:     p.Email,
		Debug:     p.Debug,
	}
}

func (r *SingleRunner) Run() error {
	conf := config.Config{
		Debug: r.Debug,
	}

	svc, err := minfraud.New(conf)
	if err != nil {
		return err
	}

	resp, err := r.execAPI(svc)
	if err != nil {
		return err
	}

	b, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	// just print response in json format
	fmt.Printf("%s\n", string(b))
	return nil
}

func (r *SingleRunner) execAPI(svc *minfraud.MinFraud) (minfraud.APIResponse, error) {
	params := minfraud.BaseRequest{
		Device: &minfraud.DeviceData{
			IPAddress: r.IPAddress,
		},
		Email: &minfraud.EmailData{
			Address: r.Email,
		},
	}
	return execAPI(svc, params, r.Command)
}

func execAPI(svc *minfraud.MinFraud, req minfraud.BaseRequest, command string) (resp minfraud.APIResponse, err error) {
	switch command {
	case commandScore:
		r, err := svc.Score(req)
		if err != nil {
			return resp, err
		}
		resp = r.APIResponse()
	case commandInsights:
		r, err := svc.Insights(req)
		if err != nil {
			return resp, err
		}
		resp = r.APIResponse()
	case commandFactors:
		r, err := svc.Factors(req)
		if err != nil {
			return resp, err
		}
		resp = r.APIResponse()
	}
	return resp, nil
}
