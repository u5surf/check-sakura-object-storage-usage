package checkusage

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jessevdk/go-flags"
	"github.com/mackerelio/checkers"
)

var opts struct {
	Warning  *string `short:"w" long:"warning" value-name:"N%" description:"Exit with WARNING status if less than N% of bucket are free"`
	Critical *string `short:"c" long:"critical" value-name:"N%" description:"Exit with CRITICAL status if less than N% of bucket are free"`
	Site     *string `short:"s" long:"site" value-name:"SITE" description:"Choose a site where monitored bucket"`
	Bucket   *string `short:"b" long:"bucket" value-name:"BUCKET" description:"Choose a monitoring bucket"`
}

// Do the plugin
func Do() {
	ckr := run(os.Args[1:])
	ckr.Name = "Sakura Object Storage Usage"
	ckr.Exit()
}

func checkUsage(current checkers.Status, threshold string, usage *Usage, status checkers.Status) (checkers.Status, error) {
	thresholdPct, err := strconv.ParseFloat(strings.TrimRight(threshold, "%"), 64)
	if err != nil {
		return checkers.UNKNOWN, err
	}
	freePct := getFreePct(usage)
	if thresholdPct > freePct {
		current = status
	}
	return current, nil
}

func getFreePct(usage *Usage) float64 {
	return float64(100) - usage.amount*100/usage.quota
}

func run(args []string) *checkers.Checker {

	apiKey, ok := os.LookupEnv("SAKURA_API_ACCESS_TOKEN")
	if !ok {
		return checkers.Unknown(fmt.Sprintln("SAKURA_API_ACCESS_TOKEN is not set"))
	}
	apiSecret, ok := os.LookupEnv("SAKURA_API_ACCESS_TOKEN_SECRET")
	if !ok {
		return checkers.Unknown(fmt.Sprintln("SAKURA_API_ACCESS_TOKEN_SECRET is not set"))
	}

	_, err := flags.ParseArgs(&opts, args)
	if err != nil {
		os.Exit(1)
	}
	if opts.Site == nil {
		return checkers.Unknown(fmt.Sprintf("site is required"))
	}

	usage, err := GetUsage(*opts.Site, *opts.Bucket, apiKey, apiSecret)
	if err != nil {
		return checkers.Unknown(fmt.Sprintf("%w", err))
	}
	current := checkers.OK
	if current != checkers.CRITICAL && opts.Critical != nil {
		current, err = checkUsage(current, *opts.Critical, usage, checkers.CRITICAL)
		if err != nil {
			return checkers.Unknown(fmt.Sprintf("site:%s, bucket:%s: Failed to check usage status: %s", *opts.Site, *opts.Bucket, err))
		}
	}
	if current == checkers.OK && opts.Warning != nil {
		current, err = checkUsage(current, *opts.Warning, usage, checkers.WARNING)
		if err != nil {
			return checkers.Unknown(fmt.Sprintf("site:%s, bucket:%s: Failed to check usage status: %s", *opts.Site, *opts.Bucket, err))
		}
	}
	msg := fmt.Sprintf("usage: site:%s, bucket:%s, current free: %f%%", *opts.Site, *opts.Bucket, getFreePct(usage))
	return checkers.NewChecker(current, msg)
}
