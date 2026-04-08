package main

import (
	"github.com/krish0723/ait/internal/doctor"
	"github.com/krish0723/ait/internal/rules"
)

func init() {
	doctor.SetBuiltinRules(rules.All())
}
