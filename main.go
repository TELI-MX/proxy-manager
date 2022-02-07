package main

import (
	"fmt"

	"github.com/KyleBanks/dockerstats"
)

func main() {
	stats, err := dockerstats.Current()
	if err != nil {
		panic(err)
	}

	for _, s := range stats {
		fmt.Println("Container", s.Container, "CPU Usage:", s.CPU, "Memory Usage:", s.Memory.Percent, "Network Usage:", s.IO.Network)
	}
}
