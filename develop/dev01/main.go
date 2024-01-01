package main

import (
	"fmt"
	"os"

	ntp "github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Errorf("%w", err)
		os.Exit(-1)
	}

	fmt.Println(time)
}
