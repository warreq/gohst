package main

import (
	"time"
)

type invocation struct {
	Directory string
	Shell     string
	User      string
	Host      string
	Command   string
	Status    int8
	Tags      []string
	Timestamp time.Time
	HasStatus bool
	IsSynced  bool
}