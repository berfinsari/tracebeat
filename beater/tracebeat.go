package beater

import (
	"fmt"
	"time"

	"github.com/aeden/traceroute"
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/cfgfile"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"

	"github.com/berfinsari/tracebeat/config"
)

type Tracebeat struct {
	done     chan struct{}
	config   config.TracebeatConfig
	client   publisher.Client
	period   time.Duration
	TbConfig config.ConfigSettings
	host     string
	maxHops  int
	hops     [][]string
}

const selector = "tracebeat"

func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	tb := &Tracebeat{
		done: make(chan struct{}),
	}
	err := cfgfile.Read(&tb.TbConfig, "")
	if err != nil {
		logp.Err("Error reading configuration file: %v", err)
		return nil, fmt.Errorf("Error reading configuration file: %v", err)
	}

	return tb, nil
}

func (tb *Tracebeat) Run(b *beat.Beat) error {
	logp.Info("tracebeat is running! Hit CTRL-C to stop it.")

	tb.client = b.Publisher.Connect()
	tb.CheckConfig(b)
	traceroute, traceerr := tb.TraceHops(b) //return değeri geliyor
	if traceerr != nil {
		return fmt.Errorf("%v", traceerr)
	}

	ticker := time.NewTicker(tb.period)
	counter := 1
	for {
		select {
		case <-tb.done:
			return nil
		case <-ticker.C:
		}

		event := common.MapStr{
			"@timestamp": common.Time(time.Now()),
			"type":       b.Name,
			"counter":    counter,
			"traceroute": traceroute,
		}
		tb.client.PublishEvent(event)
		logp.Info("Event sent")
		counter++
	}
}

func (tb *Tracebeat) CheckConfig(b *beat.Beat) {
	if tb.TbConfig.Input.Period != nil {
		tb.period = time.Duration(*tb.TbConfig.Input.Period) * time.Second
	} else {
		tb.period = 30 * time.Second
	}

	tb.host = *tb.TbConfig.Input.Host

	if tb.TbConfig.Input.MaxHops != nil {
		tb.maxHops = *tb.TbConfig.Input.MaxHops
	} else {
		tb.maxHops = 64
	}

	options := traceroute.TracerouteOptions{}
	options.SetMaxHops(tb.maxHops)

	logp.Debug(selector, "Host %v", tb.host)
	logp.Debug(selector, "Period %v", tb.period)
	logp.Debug(selector, "MaxHops %v", tb.maxHops)
}

func (tb *Tracebeat) Stop() {
	tb.client.Close()
	close(tb.done)
}
