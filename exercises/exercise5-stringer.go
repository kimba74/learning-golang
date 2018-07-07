package exercises

import (
	"fmt"
	"learning-golang/utils"
)

type ipaddr [4]byte

func (ip ipaddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

// Exercise5 runs the corresponding exercise
func Exercise5() {
	utils.ExerciseHeader(5)
	hosts := map[string]ipaddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v (%v)\n", name, ip)
	}
}
