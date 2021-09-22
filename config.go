package main

import "time"

type flags struct {
	Kubeconfig          string
	ResyncPeriodS       string
	ResyncPeriod        time.Duration
	StatusAddr          string
	AllowAll            bool
	StripOwnerReference bool
	LogLevel            string
	LogFormat           string
}
