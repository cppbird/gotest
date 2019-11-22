package main

import (
	"fmt"

	"github.com/cppbird/cron/v3"
)

func main() {
	c := cron.New()
	fmt.Println(c.AddFunc("0 3 * * *", func() { fmt.Println("Every hour on the half hour") }))
}
