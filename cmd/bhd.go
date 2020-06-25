package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/xslook/carl/pkg/bhd"
)

const (
	bhdFlagNameAll      = "all"
	bhdFlagNameFmtTable = "table"
	bhdFlagNameFmtLine  = "line"
)

var (
	bhdFlagTable = &cli.BoolFlag{
		Name:    bhdFlagNameFmtTable,
		Usage:   "Display all details",
		Aliases: []string{"a"},
	}

	bhdFlagFmtTable = &cli.BoolFlag{
		Name:    bhdFlagNameFmtTable,
		Usage:   "Display table format",
		Aliases: []string{"t"},
	}

	bhdFlagFmtLine = &cli.BoolFlag{
		Name:    bhdFlagNameFmtLine,
		Usage:   "Display line format",
		Aliases: []string{"l"},
	}
)

func bhdTableOutput(results []*bhd.Result) error {
	fmt.Printf("Binary\tOctal\tDecimal\tHexadecimal\n")
	for _, res := range results {
		fmt.Printf("%s\t%s\t%s\t%s\n", res.Bin, res.Oct, res.Dec, res.Hex)
	}
	return nil
}

func bhdLineOutput(results []*bhd.Result) error {
	for _, res := range results {
		fmt.Printf("Binary:\t%s\n", res.Bin)
		fmt.Printf("Octal:\t%s\n", res.Oct)
		fmt.Printf("Decimal:\t%s\n", res.Dec)
		fmt.Printf("Hexadecimal:\t%s\n\n", res.Hex)
	}
	return nil
}

func bhdHandler(c *cli.Context) error {
	args := c.Args()
	results := make([]*bhd.Result, 0, args.Len())
	for _, arg := range args.Slice() {
		res, err := bhd.Convert(arg)
		if err != nil {
			return err
		}
		results = append(results, res)
	}

	var err error
	tableFormat := c.Bool(bhdFlagNameFmtTable)
	lineFormat := c.Bool(bhdFlagNameFmtLine)
	if tableFormat {
		err = bhdTableOutput(results)
	} else if lineFormat {
		err = bhdLineOutput(results)
	}
	return err
}

var bhdCommand = &cli.Command{
	Name:  "bhd",
	Usage: "Number conversion",
	Flags: []cli.Flag{
		bhdFlagFmtTable,
		bhdFlagFmtLine,
	},
	Action: bhdHandler,
}
