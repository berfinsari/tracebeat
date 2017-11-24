package beater

import (
	"fmt"

	"github.com/aeden/traceroute"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
)

func (tb *Tracebeat) Trace(b *beat.Beat) ([]common.MapStr, error) {
	trace := make([]common.MapStr, 0)
	options := &traceroute.TracerouteOptions{}
	options.SetTimeoutMs(tb.timeoutMs)
	options.SetMaxHops(tb.maxHops)
	options.SetRetries(tb.retries)
	options.SetPacketSize(tb.packetSize)

	traceresult, err := traceroute.Traceroute(tb.host, options)
	if err != nil {
		return nil, err
	}

	for i, h := range traceresult.Hops {
		add := fmt.Sprintf("%d.%d.%d.%d", h.Address[0], h.Address[1], h.Address[2], h.Address[3])
		var elapTime float64
		elapTime = float64(h.ElapsedTime.Nanoseconds()) / 1000000

		hop := common.MapStr{
			"hopNumber":   i + 1,
			"success":     h.Success,
			"address":     add,
			"hostName":    h.Host,
			"n":           h.N,
			"ttl":         h.TTL,
			"elapsedTime": elapTime,
		}

		trace = append(trace, hop)
	}
	return trace, nil
}
