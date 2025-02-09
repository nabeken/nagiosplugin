package main

import (
	"flag"

	"github.com/nabeken/nagiosplugin/v2"
)

func main() {
	check := nagiosplugin.NewCheck("TEST CHECK")

	var (
		critical = flag.Bool("critical", false, "trigger CRITICAL")
		warning  = flag.Bool("warning", false, "trigger WARNING")
		ok       = flag.Bool("ok", false, "trigger OK")
	)

	flag.Parse()

	switch {
	case *critical:
		check.AddResult(nagiosplugin.CRITICAL, "CRITICAL")
		check.AddResult(nagiosplugin.WARNING, "WARNING")
	case *warning:
		check.AddResult(nagiosplugin.WARNING, "WARNING")
		check.AddResult(nagiosplugin.OK, "OK")
	case *ok:
		check.AddResult(nagiosplugin.OK, "OK")
		check.AddResult(nagiosplugin.UNKNOWN, "UNKNOWN")
	}

	defer check.Finish()
}
