package beater

import (
	"fmt"
	"strconv"

	"github.com/aeden/traceroute"
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/logp"
)

func (tb *Tracebeat) HopsOptions(hop traceroute.TracerouteHop) ([]string, error) {
	var hopArray []string
	addr := fmt.Sprintf("%v.%v.%v.%v", hop.Address[0], hop.Address[1], hop.Address[2], hop.Address[3])
	hostOrAddr := addr

	if hop.Host != "" {
		hostOrAddr = hop.Host
	}
	if !hop.Success {
		logp.Err("Error reading hop")
		return nil, fmt.Errorf("Error reading hop")
	}
	ttlcon := strconv.Itoa(hop.TTL)
	elaptime := (hop.ElapsedTime).String()
	hopArray = append(hopArray, ttlcon, hostOrAddr, addr, elaptime)
	return hopArray, nil
}

func (tb *Tracebeat) TraceHops(b *beat.Beat) ([][]string, error) {
	options := traceroute.TracerouteOptions{}
	options.SetMaxHops(tb.maxHops)
	ttl := 0
	c := make(chan traceroute.TracerouteHop, 0)
	go func() {
		for {
			hop, ok := <-c
			if !ok {
				return
			}

			array, err := tb.HopsOptions(hop)
			if err != nil {
				return
			}
			tb.hops = append(tb.hops, array)
			ttl++
		}
	}()
	_, err := traceroute.Traceroute(tb.host, &options, c)
	if err != nil {
		logp.Err("%v", err)
		return nil, fmt.Errorf("%v", err)
	}
	return tb.hops, nil
}
