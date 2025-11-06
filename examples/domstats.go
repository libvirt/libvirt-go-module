package main

import (
	"fmt"
	libvirt "libvirt.org/go/libvirt"
	"os"
	"encoding/json"
)

func main() {
	var uri string

	if len(os.Args) >= 2 {
		uri = os.Args[1]
	}

	conn, err := libvirt.NewConnect(uri)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	var domains []*libvirt.Domain

	if len(os.Args) >= 3 {
		dom, err := conn.LookupDomainByName(os.Args[2])
		if (err != nil) {
			panic(err)
		}

		domains = make([]*libvirt.Domain, 1)
		domains[0] = dom
	}

	stats, err := conn.GetAllDomainStats(domains, 0, 0)
	if err != nil {
		panic(err)
	}

	for _, st := range stats {
		js, err := json.MarshalIndent(st, "", "  ")
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s\n", js)
	}

}
