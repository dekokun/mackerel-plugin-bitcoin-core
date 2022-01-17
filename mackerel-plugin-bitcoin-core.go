package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/btcsuite/btcd/rpcclient"
	mp "github.com/mackerelio/go-mackerel-plugin"
)

type BitcoinCorePlugin struct {
	Prefix   string
	Dest     string
	User     string
	Password string
}

func (b BitcoinCorePlugin) GraphDefinition() map[string]mp.Graphs {
	labelPrefix := strings.Title(b.MetricKeyPrefix())
	return map[string]mp.Graphs{
		"block": {
			Label: labelPrefix + " block height",
			Unit:  mp.UnitFloat,
			Metrics: []mp.Metrics{
				{Name: "height", Label: "height", Diff: false, Stacked: false},
			},
		},
		"network": {
			Label: labelPrefix + " network score",
			Unit:  mp.UnitFloat,
			Metrics: []mp.Metrics{
				{Name: "score.*", Label: "score.*", Diff: false, Stacked: false},
			},
		},
		"peer": {
			Label: labelPrefix + " peer connections",
			Unit:  mp.UnitFloat,
			Metrics: []mp.Metrics{
				{Name: "in", Label: "in", Diff: false, Stacked: true},
				{Name: "out", Label: "out", Diff: false, Stacked: true},
				{Name: "total", Label: "total", Diff: false, Stacked: false},
			},
		},
	}
}

func (b BitcoinCorePlugin) FetchMetrics() (map[string]float64, error) {
	connCfg := &rpcclient.ConnConfig{
		Host:                b.Dest,
		User:                b.User,
		Pass:                b.Password,
		DisableConnectOnNew: true,
		HTTPPostMode:        true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:          true, // Bitcoin core does not provide TLS by default
	}
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()
	stat := map[string]float64{}
	// Get the current block count.
	blockCount, err := client.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	stat["height"] = float64(blockCount)
	info, err := client.GetNetworkInfo()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v", info.LocalAddresses[0].Score)
	for _, v := range info.LocalAddresses {
		stat["network.score."+v.Address] = float64(v.Score)
	}
	peerInfo, err := client.GetPeerInfo()
	total := 0
	in := 0
	out := 0
	for _, v := range peerInfo {
		if v.Inbound {
			in += 1
		} else {
			out += 1
		}
		total += 1
	}
	stat["in"] = float64(in)
	stat["out"] = float64(out)
	stat["total"] = float64(total)

	if err != nil {
		log.Fatal(err)
	}

	return stat, nil
}

func (b BitcoinCorePlugin) MetricKeyPrefix() string {
	if b.Prefix == "" {
		b.Prefix = "bitcoin"
	}
	return b.Prefix
}

var version string
var commit string

func main() {
	log.SetFlags(0)
	err := Run(context.Background(), os.Args[1:], os.Stdout, os.Stderr)
	if err != nil && err != flag.ErrHelp {
		log.Println(err)
		exitCode := 1
		if ecoder, ok := err.(interface{ ExitCode() int }); ok {
			exitCode = ecoder.ExitCode()
		}
		os.Exit(exitCode)
	}
}

const cmdName = "mackerel-plugin-bitcoin-core"

// Run the mackerel-plugin-bitcoin-core
func Run(ctx context.Context, argv []string, outStream, errStream io.Writer) error {
	log.SetOutput(errStream)
	fs := flag.NewFlagSet(
		fmt.Sprintf("%s (v%s rev:%s)", cmdName, version, commit), flag.ContinueOnError)
	fs.SetOutput(errStream)
	ver := fs.Bool("version", false, "display version")
	optHost := fs.String("host", "localhost", "host to connect")
	optPort := fs.String("port", "8332", "port to connect")
	optUser := fs.String("user", "", "rpc user")
	optPassword := fs.String("password", "", "rpc password")
	optPrefix := fs.String("metric-key-prefix", "bitcoin", "Metric key prefix")
	if err := fs.Parse(argv); err != nil {
		return err
	}
	if *ver {
		return printVersion(outStream)
	}
	b := BitcoinCorePlugin{
		Prefix:   *optPrefix,
		Dest:     *optHost + ":" + *optPort,
		User:     *optUser,
		Password: *optPassword,
	}
	plugin := mp.NewMackerelPlugin(b)
	plugin.Run()
	return nil
}

func printVersion(out io.Writer) error {
	_, err := fmt.Fprintf(out, "%s v%s (rev:%s)\n", cmdName, version, commit)
	return err
}
