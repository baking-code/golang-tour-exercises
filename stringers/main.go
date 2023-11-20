package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	ipSlice := make([]interface{}, len(ip))
	for ind, val := range ip {
		ipSlice[ind] = val
	}
	//	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
	// spread the slice to avoid needing to specifiy each value!
	return fmt.Sprintf("%v.%v.%v.%v", ipSlice...)
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
