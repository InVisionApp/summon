package command

import (
	"fmt"
	"github.com/codegangsta/cli"
	"io/ioutil"
)

var ProviderCommand = cli.Command{
	Name:  "provider",
	Usage: "List, install and uninstall providers",
	Subcommands: []cli.Command{
		{
			Name:  "list",
			Usage: "List installed providers",
			Action: func(c *cli.Context) {
				files, _ := ioutil.ReadDir("/usr/libexec/summon")
				for _, f := range files {
					fmt.Println(f.Name())
				}
			},
		},
		{
			Name:  "install",
			Usage: "Install a new provider",
		},
		{
			Name:  "uninstall",
			Usage: "Uninstall a provider",
		},
	},
}
