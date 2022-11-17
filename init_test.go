package main

import (
	"flag"
	"strings"
	"testing"

	"git.restream.ru/itv-backend/reindexer/benchmarks/repo"
)

func TestMain(t *testing.T) {
	benches := flag.CommandLine.Lookup("test.bench").Value.String()
	repo.Start(strings.Index(benches, "Seed") >= 0, benches, itemsInDataSet)
}
