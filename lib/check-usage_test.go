package checkusage

import (
	"errors"
	"testing"

	"github.com/mackerelio/checkers"
	"github.com/stretchr/testify/assert"
)

type ObjectStorageAPIMock struct{}

func (o *ObjectStorageAPIMock) GetUsage() (*Usage, error) {
	return &Usage{
		quota:  float64(10240),
		amount: float64(9217.024),
	}, nil
}

func TestRun(t *testing.T) {
	tests := []struct {
		desc     string
		warning  string
		critical string
		exp      checkers.Status
	}{
		{
			desc:     "90.01%% should be warning",
			warning:  "10%",
			critical: "5%",
			exp:      checkers.WARNING,
		},
		{
			desc:     "90.01%% should be critical",
			warning:  "15%",
			critical: "10%",
			exp:      checkers.CRITICAL,
		},
		{
			desc:     "90.01%% should be ok",
			warning:  "9%",
			critical: "8%",
			exp:      checkers.OK,
		},
		{
			desc:     "invalid warning",
			warning:  "foo",
			critical: "5%",
			exp:      checkers.UNKNOWN,
		},
		{
			desc:     "invalid critical",
			warning:  "15%",
			critical: "foo",
			exp:      checkers.UNKNOWN,
		},
	}

	for _, tt := range tests {
		site := "foo"
		bucket := "bar"
		opts.Warning = &tt.warning
		opts.Critical = &tt.critical
		opts.Site = &site
		opts.Bucket = &bucket
		cli := &ObjectStorageAPIMock{}
		r := &runner{
			cli: cli,
		}
		ckr := r.Run()
		assert.Equal(t, ckr.Status, tt.exp, tt.desc)
	}
}

type ObjectStorageNegAPIMock struct{}

func (o *ObjectStorageNegAPIMock) GetUsage() (*Usage, error) {
	return nil, errors.New("somthing wrong")
}

func TestRunNegative(t *testing.T) {
	opts.Warning = nil
	opts.Critical = nil
	opts.Site = nil
	opts.Bucket = nil
	cli := &ObjectStorageNegAPIMock{}
	r := &runner{
		cli: cli,
	}
	ckr := r.Run()
	assert.Equal(t, ckr.Status, checkers.UNKNOWN, "GetUsage returns some error, Run returns status.UNKNOWN")
}

func TestParseOption(t *testing.T) {
	tests := []struct {
		desc       string
		prepareEnv func(t *testing.T)
		args       []string
		exp        checkers.Status
	}{
		{
			desc: "site is not set",
			prepareEnv: func(t *testing.T) {
				opts.Warning = nil
				opts.Critical = nil
				opts.Site = nil
				opts.Bucket = nil
				t.Setenv("SAKURA_ACCESS_TOKEN", "SAKURA_ACCESS_TOKEN")
				t.Setenv("SAKURA_ACCESS_TOKEN_SECRET", "SAKURA_ACCESS_TOKEN_SECRET")
			},
			args: []string{
				"-b",
				"sample",
			},
			exp: checkers.UNKNOWN,
		},
		{
			desc: "bucket is not set",
			prepareEnv: func(t *testing.T) {
				opts.Warning = nil
				opts.Critical = nil
				opts.Site = nil
				opts.Bucket = nil
				t.Setenv("SAKURA_ACCESS_TOKEN", "SAKURA_ACCESS_TOKEN")
				t.Setenv("SAKURA_ACCESS_TOKEN_SECRET", "SAKURA_ACCESS_TOKEN_SECRET")
			},
			args: []string{
				"-s",
				"iks01",
			},
			exp: checkers.UNKNOWN,
		},
	}

	for _, tt := range tests {
		tt.prepareEnv(t)
		act, _, _ := ParseOption(tt.args)
		assert.Equal(t, act.Status, tt.exp, tt.desc)
	}
}
