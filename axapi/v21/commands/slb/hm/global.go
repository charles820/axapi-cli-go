package hm

import (
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"

	. "github.com/K-DEN/axapi-cli-go/axapi/util"
	. "github.com/K-DEN/axapi-cli-go/axapi/v21/auth"
	. "github.com/K-DEN/axapi-cli-go/axapi/v21/util"
)

type Global struct{}

func (g *Global) Help() string {
	return `
$ ... slb.hm.global <subcommand>

To see options for each <subcommand>, try
$ ... slb.hm.global <subdommand> --help

Available Subcommands:
	get              : "slb.hm.global.get method."
	`
}

func (g *Global) Synopsis() string {
	return "slb.hm.global method."
}

func (g *Global) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(g.Help())
		os.Exit(0)
	}
	var res string
	switch args[0] {
	case "get":
		res, _ = g.Get(args)
	}
	fmt.Println(res)
	return 0
}

/**
 * AXAPI(2.1) Method: Slb.Hm.Global.Get
 */

func (g *Global) Get(args []string) (string, error) {
	// options
	var opts GlobalFlags

	// parse
	old_args := os.Args
	os.Args[0] = strings.Join(os.Args[:(len(os.Args)-len(args)+1)], " ")
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}
	os.Args = old_args
	if opts.SessionId == "" {
		opts.SessionId = IssueSessionId()
	}

	// dispatch
	item := CreateUrlValues(map[string]string{
		"method":     "slb.hm.global.get",
		"session_id": opts.SessionId,
	})
	req := CreateRequest()
	req.QueryString = item
	res, err := Dispatch(req)
	return res, err
}
