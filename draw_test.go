package sterm

import "testing"

var tbl = [][]string{
	{"name", "ver", "cpu", "ram", "up", "ping", "ip"},
	{"fedora", "36", "82%", "2300/15884 mb", "1d 22h 37m", "10 ms", "157.116.227.217"},
	{"arch", ">3", "1%", "2300/15884 mb", "1h 20m", "1345 ms", "211.216.51.124"},
	{"debian", "11", "43%", "158/15884 mb", "12h 1m", "128 ms", "157.116.227.217"},
	{"slackware", "15", "0%", "4583/15884 mb", "1d 1h 1m", "29 ms", "146.132.226.133"},
}

func BenchmarkDrawTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = DrawTable(tbl, RoundedBorders)
	}
}
