package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/urfave/cli/v2"
)

const (
	flagFileName   = "file"
	flagOutputName = "output"
)

var (
	flagFile = &cli.StringFlag{
		Name:    flagFileName,
		Aliases: []string{"f"},
		Usage:   "Input file, default: stdin",
	}
	flagOutput = &cli.StringFlag{
		Name:    flagOutputName,
		Aliases: []string{"o"},
		Usage:   "Output file, default: stdout",
	}
)

func isValidJSON(bts []byte) bool {
	if len(bts) < 2 {
		return false
	}
	return true
}

func cmdAction(c *cli.Context) error {
	return nil
}

func encodeAction(c *cli.Context) error {
	return nil
}

func getFormatJSON(bts []byte) ([]byte, error) {
	m := make(map[string]interface{}, 0)
	err := json.Unmarshal(bts, &m)
	if err != nil {
		return nil, err
	}
	bytes, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func decodeAction(c *cli.Context) error {
	var file string
	if c.IsSet(flagFileName) {
		file = c.String(flagFileName)
	}

	outputs := make([][]byte, 0)
	if file != "" {
		bts, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}

		out, err := getFormatJSON(bts)
		if err != nil {
			return err
		}
		outputs = append(outputs, out)
	}

	args := c.Args()
	for _, arg := range args.Slice() {
		out, err := getFormatJSON([]byte(arg))
		if err != nil {
			return err
		}
		outputs = append(outputs, out)
	}

	for _, output := range outputs {
		fmt.Printf("%s\n", output)
	}
	return nil
}

func verifyAction(c *cli.Context) error {
	return nil
}

var (
	encodeCmd = &cli.Command{
		Name:    "encode",
		Usage:   "Encode a JSON object to a string",
		Aliases: []string{"marshal", "stringify"},
		Flags: []cli.Flag{
			flagFile,
			flagOutput,
		},
		Action: encodeAction,
	}
	decodeCmd = &cli.Command{
		Name:    "decode",
		Usage:   "Decode a string to a JSON object",
		Aliases: []string{"unmarshal", "parse"},
		Flags: []cli.Flag{
			flagFile,
			flagOutput,
		},
		Action: decodeAction,
	}
	verifyCmd = &cli.Command{
		Name:  "verify",
		Usage: "Verify string is a valid JSON or not",
		Flags: []cli.Flag{
			flagFile,
		},
		Action: verifyAction,
	}

	jsonCommand = &cli.Command{
		Name:  "json",
		Usage: "JSON toolbox",
		Flags: []cli.Flag{},
		Subcommands: []*cli.Command{
			encodeCmd,
			decodeCmd,
			verifyCmd,
		},
		Action: cmdAction,
	}
)
