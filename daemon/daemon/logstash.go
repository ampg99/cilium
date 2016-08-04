package daemon

import (
	"encoding/json"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/noironetworks/cilium-net/bpf/policymap"
)

type LogstashStat struct {
	FromID  uint32
	From    string
	ToID    string
	Bytes   uint64
	Packets uint64
	Action  string
}

func newLogstashClient(addr string) net.Conn {
	i := 3
	for {
		c, err := net.Dial("tcp", addr)
		if err != nil {
			if i >= 0 {
				log.Errorf("Error while connecting to Logstash address %s: %s", addr, err)
				if i == 0 {
					log.Info("Mutting Logstash connection errors but still retrying...")
				}
				i--
			}
		} else {
			log.Infof("Connection to Logstash %s successfully made", addr)
			return c
		}
		time.Sleep(10 * time.Second)
	}
}

func (d *Daemon) EnableLogstash(LogstashAddr string, refreshTime int) {
	readStats := func(c net.Conn) {
		defer func() {
			recover()
		}()
		for {
			timeToProcess1 := time.Now()

			allPes := map[uint16][]policymap.PolicyEntryDump{}
			d.endpointsMU.RLock()
			for _, ep := range d.endpoints {
				pes, err := ep.PolicyMap.DumpToSlice()
				if err != nil {
					continue
				}
				allPes[ep.ID] = pes
			}
			d.endpointsMU.RUnlock()
			lss := d.processStats(allPes)
			for _, ls := range lss {
				if err := json.NewEncoder(c).Encode(ls); err != nil {
					log.Errorf("Error while sending data to Logstash: %s", err)
					timeToProcess2 := time.Now()
					time.Sleep(time.Second*time.Duration(refreshTime) - timeToProcess2.Sub(timeToProcess1))
					return
				}
			}

			timeToProcess2 := time.Now()
			time.Sleep(time.Second*time.Duration(refreshTime) - timeToProcess2.Sub(timeToProcess1))
		}
	}
	for {
		c := newLogstashClient(LogstashAddr)
		readStats(c)
	}
}

func (d *Daemon) getInlineLabelStr(id uint32) string {
	l, err := d.GetCachedLabelList(id)
	if err != nil {
		return ""
	}
	inlineLblSlice := []string{}
	for _, lbl := range l {
		inlineLblSlice = append(inlineLblSlice, lbl.String())
	}
	return strings.Join(inlineLblSlice, "\n")
}

func (d *Daemon) processStats(allPes map[uint16][]policymap.PolicyEntryDump) []LogstashStat {
	lss := []LogstashStat{}
	for k, v := range allPes {
		if len(v) == 0 {
			continue
		}
		for _, stat := range v {
			lss = append(lss, LogstashStat{
				FromID:  stat.ID,
				From:    d.getInlineLabelStr(stat.ID),
				ToID:    strconv.FormatUint(uint64(k), 10),
				Bytes:   stat.Bytes,
				Packets: stat.Packets,
				Action:  stat.PolicyEntry.String(),
			})
		}
	}
	return lss
}