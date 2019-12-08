package timestamp

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

const (
	flagAllName  = "all"
	flagZoneName = "zone"
)

const (
	defaultFormat = "2006-01-02 15:04:05.000000 Z0700"
)

var (
	flagAll = &cli.BoolFlag{
		Name:    flagAllName,
		Usage:   "Display all details",
		Aliases: []string{"a"},
	}
	flagZone = &cli.IntFlag{
		Name:    flagZoneName,
		Usage:   "The timezone of output time, default local timezone",
		Aliases: []string{"z"},
	}
)

func Command() *cli.Command {
	var cmd = &cli.Command{
		Name:    "timestamp",
		Usage:   "Convert timestamp to time",
		Aliases: []string{"ts"},
		Flags: []cli.Flag{
			flagAll,
			flagZone,
		},
		Action: handle,
	}
	return cmd
}

// 1575681996,294,000,000
func padTimestamp(ts string) string {
	if len(ts) < 19 {
		ts = ts + strings.Repeat("0", 19-len(ts))
	}
	return ts[:19]
}

func parseTimestamp(ts string) (t time.Time, err error) {
	var sec, nao int64
	if len(ts) > 10 {
		ts = padTimestamp(ts)
		seconds := ts[:10]
		nanoseconds := ts[10:]
		sec, err = strconv.ParseInt(seconds, 10, 64)
		if err != nil {
			return
		}
		nao, err = strconv.ParseInt(nanoseconds, 10, 64)
		if err != nil {
			return
		}
	} else {
		sec, err = strconv.ParseInt(ts, 10, 64)
		if err != nil {
			return
		}
	}
	t = time.Unix(sec, nao)
	return
}

func handle(c *cli.Context) error {
	args := c.Args()
	if args.Len() < 1 {
		return nil
	}
	var zone *time.Location
	if c.IsSet(flagZoneName) {
		offset := c.Int(flagZoneName)
		var label string
		if offset > 0 {
			label = fmt.Sprintf("UTC+%d", offset)
		} else {
			label = fmt.Sprintf("UTC%d", offset)
		}
		zone = time.FixedZone(label, offset*60*60)
	} else {
		zone = time.Local
	}
	for _, arg := range args.Slice() {
		t, err := parseTimestamp(arg)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Parse timestamp: %s failed", arg))
		}
		if zone != time.Local {
			t = t.In(zone)
		}
		fmt.Println(t.Format(defaultFormat))
	}
	return nil
}
