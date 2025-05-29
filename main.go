package main

import (
	"poefilter/filters"
)

func main() {
	filters.CreatePhase1Filter()
	filters.CreatePhase2Filter()
	filters.CreateDeadeye1()
	filters.CreateDeadeye2()
}
