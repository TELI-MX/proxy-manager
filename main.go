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
		fmt.Println(s)
		fmt.Println(s.Container)
		fmt.Println(s.Memory)
		fmt.Println(s.CPU)
	}
}
