package metrics

import (
	"fmt"
	"net"
	"strconv"

	statsd "github.com/etsy/statsd/examples/go"
)

//go:generate mockery --name StatsdClient --outpkg mocks --output ../lib/mocks/ --with-expecter --structname StatsdClient --filename statsd_monitor.go
type StatsdClient interface {
	Timing(string, int64)
	Increment(string)
}

func NewStatsdClient(statsdAddress string) *statsd.StatsdClient {
	statsdHost, statsdPortStr, _ := net.SplitHostPort(statsdAddress)
	statsdPort, _ := strconv.Atoi(statsdPortStr)
	return statsd.New(statsdHost, statsdPort)
}

type StatsdMonitor struct {
	client    StatsdClient
	prefix    string
	separator string
}

func NewStatsdMonitor(client StatsdClient, prefix string, separator string) *StatsdMonitor {
	return &StatsdMonitor{
		client:    client,
		prefix:    prefix,
		separator: separator,
	}
}

func (mm *StatsdMonitor) Duration(operation string, duration int) {
	statName := fmt.Sprintf("%s%s%s,operation=%s", mm.prefix, mm.separator, "duration", operation)
	mm.client.Timing(statName, int64(duration))
}

func (mm *StatsdMonitor) ResponseTime(requestMethod string, requestUrl string, responseTime int64) {
	statName := fmt.Sprintf("%s%s%s,%s",
		mm.prefix,
		mm.separator,
		"responseTime",
		Tags{requestMethod, requestUrl})
	mm.client.Timing(statName, responseTime)
}

func (mm *StatsdMonitor) ResponseStatus(requestMethod string, requestUrl string, responseCode int) {
	statName := fmt.Sprintf("%s%s%s,statusCode=%d,%s",
		mm.prefix,
		mm.separator,
		"responseStatusCode",
		responseCode,
		Tags{requestMethod, requestUrl})
	mm.client.Increment(statName)
}

func (mm *StatsdMonitor) ResponseTimeGRPC(FullMethod string, responseTime int64) {
	statName := fmt.Sprintf("%s%s%s,%s",
		mm.prefix,
		mm.separator,
		"responseTime",
		fmt.Sprintf("method=%s", FullMethod))
	mm.client.Timing(statName, responseTime)
}

func (mm *StatsdMonitor) ResponseStatusGRPC(fullMethod string, statusString string) {
	statName := fmt.Sprintf("%s%s%s,statusCode=%s,%s",
		mm.prefix,
		mm.separator,
		"responseStatusCode",
		statusString,
		fmt.Sprintf("method=%s", fullMethod))
	mm.client.Increment(statName)
}
