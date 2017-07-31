package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	defaultColumns   = "InstanceId"
	emptyString      = "-"
	cacheBasePath    = "~/.cache/ec2ls-cache/"
	defaultCacheName = "out"
)

func main() {

	var profile, region, columns, sortcolumn, cachename string
	var filters []string
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
		cli.StringSliceFlag{
			Name:  "filters",
			Usage: "filters",
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
		filters = c.StringSlice("filters")
		if columns == "" {
			columns = defaultColumns
		}

		return nil
	}

	app.Run(os.Args)

	cacheinfo, err := ec2list(profile, region, updateCache, cachename, filters, columns)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	output(cacheinfo)
}
