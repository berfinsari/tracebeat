package beater

import (
	"fmt"
	"strconv"
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

func (tb *Tracebeat) CheckConfig(b *beat.Beat) {
	if tb.TbConfig.Input.Period != nil {
		tb.period = time.Duration(*tb.TbConfig.Input.Period) * time.Second
	} else {
		tb.period = 30 * time.Second
	}

	tb.host = *tb.TbConfig.Input.Host
	fmt.Printf("tb.host = %v", tb.host)

	if tb.TbConfig.Input.MaxHops != nil {
		tb.maxHops = *tb.TbConfig.Input.MaxHops
	} else {
		tb.maxHops = 64
	}

	options := traceroute.TracerouteOptions{}
	options.SetMaxHops(tb.maxHops)

}

func (tb *Tracebeat) HopsAyarla(hop traceroute.TracerouteHop) []string {
	var dizi []string
	addr := fmt.Sprintf("%v.%v.%v.%v", hop.Address[0], hop.Address[1], hop.Address[2], hop.Address[3])
	hostOrAddr := addr
	if hop.Host != "" {
		hostOrAddr = hop.Host
	}
	if hop.Success {
		ttlcon := strconv.Itoa(hop.TTL)
		//elaptime := strconv.Itoa(hop.ElapsedTime)
		dizi = append(dizi, ttlcon, hostOrAddr, addr)
	}
	return dizi
}

func (tb *Tracebeat) TraceHops() {
	options := traceroute.TracerouteOptions{}
	options.SetMaxHops(tb.maxHops)
	ttl_sayisi := 0
	c := make(chan traceroute.TracerouteHop, 0)
	go func() {
		for {
			hop, ok := <-c
			if !ok {
				fmt.Println("hop not ok")
				return
			}

			dizi := tb.HopsAyarla(hop)
			tb.hops = append(tb.hops, dizi)
			ttl_sayisi++
		}
	}()

	_, err := traceroute.Traceroute(tb.host, &options, c)
	if err != nil {
		logp.Err("Error", err)
	}
}
func (tb *Tracebeat) Run(b *beat.Beat) error {
	logp.Info("tracebeat is running! Hit CTRL-C to stop it.")
	tb.CheckConfig(b)
	tb.TraceHops()
	tb.client = b.Publisher.Connect()
	counter := 1
	ticker := time.NewTicker(tb.period)
	defer ticker.Stop()

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
		}
		tb.client.PublishEvent(event)
		logp.Info("Event sent")
		counter++
	}
}

func (tb *Tracebeat) Stop() {
	tb.client.Close()
	close(tb.done)
}
