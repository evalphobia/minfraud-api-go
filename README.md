minfraud-api-go
----

[![License: MIT][401]][402] [![GoDoc][101]][102] [![Release][103]][104] [![Build Status][201]][202] [![Coveralls Coverage][203]][204] [![Codecov Coverage][205]][206]
[![Go Report Card][301]][302] [![Code Climate][303]][304] [![BCH compliance][305]][306] [![CodeFactor][307]][308] [![codebeat][309]][310] [![Scrutinizer Code Quality][311]][312] [![FOSSA Status][403]][404]


<!-- Basic -->

[101]: https://godoc.org/github.com/evalphobia/minfraud-api-go?status.svg
[102]: https://godoc.org/github.com/evalphobia/minfraud-api-go
[103]: https://img.shields.io/github/release/evalphobia/minfraud-api-go.svg
[104]: https://github.com/evalphobia/minfraud-api-go/releases/latest
[105]: https://img.shields.io/github/downloads/evalphobia/minfraud-api-go/total.svg?maxAge=1800
[106]: https://github.com/evalphobia/minfraud-api-go/releases
[107]: https://img.shields.io/github/stars/evalphobia/minfraud-api-go.svg
[108]: https://github.com/evalphobia/minfraud-api-go/stargazers


<!-- Testing -->

[201]: https://github.com/evalphobia/minfraud-api-go/workflows/test/badge.svg
[202]: https://github.com/evalphobia/minfraud-api-go/actions?query=workflow%3Atest
[203]: https://coveralls.io/repos/evalphobia/minfraud-api-go/badge.svg?branch=master&service=github
[204]: https://coveralls.io/github/evalphobia/minfraud-api-go?branch=master
[205]: https://codecov.io/gh/evalphobia/minfraud-api-go/branch/master/graph/badge.svg
[206]: https://codecov.io/gh/evalphobia/minfraud-api-go


<!-- Code Quality -->

[301]: https://goreportcard.com/badge/github.com/evalphobia/minfraud-api-go
[302]: https://goreportcard.com/report/github.com/evalphobia/minfraud-api-go
[303]: https://codeclimate.com/github/evalphobia/minfraud-api-go/badges/gpa.svg
[304]: https://codeclimate.com/github/evalphobia/minfraud-api-go
[305]: https://bettercodehub.com/edge/badge/evalphobia/minfraud-api-go?branch=master
[306]: https://bettercodehub.com/
[307]: https://www.codefactor.io/repository/github/evalphobia/minfraud-api-go/badge
[308]: https://www.codefactor.io/repository/github/evalphobia/minfraud-api-go
[309]: https://codebeat.co/badges/142f5ca7-da37-474f-9264-f708ade08b5c
[310]: https://codebeat.co/projects/github-com-evalphobia-minfraud-api-go-master
[311]: https://scrutinizer-ci.com/g/evalphobia/minfraud-api-go/badges/quality-score.png?b=master
[312]: https://scrutinizer-ci.com/g/evalphobia/minfraud-api-go/?branch=master

<!-- License -->
[401]: https://img.shields.io/badge/License-MIT-blue.svg
[402]: LICENSE.md
[403]: https://app.fossa.com/api/projects/git%2Bgithub.com%2Fevalphobia%2Fminfraud-api-go.svg?type=shield
[404]: https://app.fossa.com/projects/git%2Bgithub.com%2Fevalphobia%2Fminfraud-api-go?ref=badge_shield


Unofficial golang library for [MaxMind minFraud](https://dev.maxmind.com/minfraud).


# Quick Usage for binary

## install

Download binary from release page, or build from source:

```bash
$ git clone --depth 1 https://github.com/evalphobia/minfraud-api-go.git
$ cd ./minfraud-api-go/cmd
$ go build -o ./minfraud-api-go .
```

## Subcommands

### root command

```bash
$ minfraud-api-go
Commands:

  help     show help
  single   Exec single api call for minFraud API
  list     Exec api call for minFraud API from csv list file
```

### single command

`single` command is used to execute single minFraud API call.

```bash
./minfraud-api-go single -h
Exec single api call for minFraud API

Options:

  -h, --help      display help information
  -c, --command  *set type of api [score, insights, factors] --command='score'
  -i, --ipaddr    input ip address --ipaddr='8.8.8.8'
  -e, --email     input email address --email='example@example.com'
      --debug     set if you use HTTP debug feature --debug
```

For example, you can check ip address or email address, or both.

```bash
# set auth data
$ export MINFRAUD_ACCOUNT_ID=xxx
$ export MINFRAUD_LICENSE_KEY=yyy

# check ip address
$ ./minfraud-api-go single -c score -i 8.8.8.8

# check email address
$ ./minfraud-api-go single -c insights -e example@example.com

# check combination of ip address and email address
$ ./minfraud-api-go single -c insights -i 8.8.8.8 -e example@example.com
```

### list command

`list` command is used to execute multiple minFraud API call from list and save risk scores to output file.

```bash
./minfraud-api-go list -h
Exec api call for minFraud API from csv list file

Options:

  -h, --help      display help information
  -c, --command  *set type of api [score, insights, factors] --command='score'
  -i, --input    *input csv/tsv file path --input='./input.csv'
  -o, --output   *output tsv file path --output='./output.tsv'
      --debug     set if you use HTTP debug feature --debug
```

For example, you can check the scores

```bash
# set auth data
$ export MINFRAUD_ACCOUNT_ID=xxx
$ export MINFRAUD_LICENSE_KEY=yyy

# prepare CSV/TSV file
$ cat input.tsv
ip_address	email
8.8.8.8	example@example.com
8.8.4.4	example@example.com
1.1.1.1	example@example.com


# check risk from the TSV file
$ ./minfraud-api-go list -c insights -i ./input.tsv -o ./output.tsv
exec #: [2]
exec #: [0]
exec #: [1]

$ cat output.tsv
ip_address	risk_score	ip_risk	ip_is_anonymous	ip_is_anonymous_vpn	ip_is_hosting_provider	ip_is_public_proxyip_is_residential_proxy	ip_is_tor_exit_node	ip_organization	ip_user_count	ip_user_type	ip_country	ip_city	ip_registered_country	ip_represented_country	email_domain_first_seen	email_first_seen	email_is_disposable	email_is_free	email_is_high_risk
8.8.8.8	0.82000	0.01000	false	false	false	false	false	false	Google	25	business	US		USfalse	false	false
8.8.4.4	0.82000	0.01000	false	false	false	false	false	false	Google	5	business	US		USfalse	false	false
1.1.1.1	0.10000	0.01000	false	false	false	false	false	false	Mountain View Communications	20	content_delivery_network	AU		AU				false	false	false
```

# Quick Usage for API

```go
package main

import (
	"fmt"

	"github.com/evalphobia/minfraud-api-go/config"
	"github.com/evalphobia/minfraud-api-go/minfraud"
)

func main() {
	conf := config.Config{
        // you can set auth values to config directly, otherwise used from environment variables.
		AccountID:  "<your MaxMind account id>",
		LicenseKey: "<your MaxMind license key>",
		Debug:      false,
	}

	svc, err := minfraud.New(conf)
	if err != nil {
		panic(err)
	}

	// execute score API
	resp, err := svc.Score(minfraud.BaseRequest{
		Device: &DeviceData{
			IPAddress: "8.8.8.8",
		},
    })
	if err != nil {
		panic(err)
	}
	if resp.HasError() {
		panic(fmt.Errorf("code=[%s] error=[%s]", resp.ErrData.Code, resp.ErrData.Error))
	}

	// just print response in json format
	b, _ := json.Marshal(resp)
	fmt.Printf("%s", string(b))
}
```

see example dir for more examples, and see [official API document](https://dev.maxmind.com/minfraud/api-documentation) for more details (especially request/response).


# Environment variables

| Name | Description |
|:--|:--|
| `MINFRAUD_ACCOUNT_ID` | [MaxMind Account ID](https://support.maxmind.com/account-faq/license-keys/how-do-i-generate-a-license-key/). |
| `MINFRAUD_LICENSE_KEY` | [MaxMind License Key](https://support.maxmind.com/account-faq/license-keys/how-do-i-generate-a-license-key/). |
