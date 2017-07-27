package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	defaultFilters = []string{
		"instance-state-name=running",
	}
	defaultColumns = []string{
		"tag:Name",
		"instance-id",
		"private-ip",
	}
	emptyString      = "-"
	cacheBasePath    = "~/.cache/ec2ls-cache/"
	defaultCacheName = "out"
)

// Ec2Info is aws ec2 instance information.
type Ec2Info struct {
	Name      string `json:"name"`
	ID        string `json:"id"`
	PrivateIP string `json:"private_ip"`
}

func main() {

	var profile, region, filters, columns, sortcolumn, cachename string
	var updateCache bool

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "profile",
			Usage:       "aws iam profile name",
			Destination: &profile,
		},
		cli.StringFlag{
			Name:        "region",
			Usage:       "aws region name",
			Destination: &region,
		},
		cli.BoolFlag{
			Name:        "update, u",
			Usage:       "update cache",
			Destination: &updateCache,
		},
		cli.StringFlag{
			Name:        "filters",
			Usage:       "filters",
			Destination: &filters,
		},
		cli.StringFlag{
			Name:        "columns",
			Usage:       "display columns",
			Destination: &columns,
		},
		cli.StringFlag{
			Name:        "sortcolumn",
			Usage:       "column name for sort order",
			Destination: &sortcolumn,
		},
		cli.StringFlag{
			Name:        "cachename",
			Usage:       "cache file name",
			Destination: &cachename,
		},
	}

	app.Name = "ec2ls-cache"
	app.Usage = ""
	app.Action = func(c *cli.Context) error {

		// TODO: validate
		if cachename == "" {
			cachename = defaultCacheName
		}

		return nil
	}

	app.Run(os.Args)

	ec2s, err := ec2list(profile, region, updateCache, cachename)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	output(ec2s)
}
