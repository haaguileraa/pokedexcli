package main

import "time"

func main() {
	timeout := 5 * time.Second
	interval := 7 * time.Second
	config := NewConfig(timeout, interval)
	listen(config)
}

