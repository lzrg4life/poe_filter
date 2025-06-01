package filters

import (
	"os"
	"poefilter/lib"
)

func CreateDeadeye1() {
	var filter lib.Filter

	lib.AddTop(&filter)
	lib.AddLeveling(&filter)

	// Deadeye leveling stuff
	filter.Show()
	filter.Class("Bow Quiver")

	filter.Show()
	filter.LinkedSockets(">=", "3")
	filter.SocketGroup("", "RRG RGG")
	filter.PlayEffect("green")
	filter.PlayAlertSound("3", "150")

	lib.AddPhase1Vendoring(&filter)
	lib.AddHideStuff(&filter)

	filter.Show() // show everything else

	f, err := os.Create("C:\\Users\\wds85\\OneDrive\\Documents\\My Games\\Path of Exile\\deadeye_1.filter")
	lib.LogFatal(err)
	defer f.Close()
	f.WriteString(filter.Text)
}

func CreateDeadeye2() {
	var filter lib.Filter

	lib.AddTop(&filter)
	lib.AddLeveling(&filter)

	// Deadeye leveling stuff
	filter.Show()
	filter.Class("Bow Quiver")

	lib.AddPhase2Vendoring(&filter)
	lib.AddHideStuff(&filter)

	filter.Show() // show everything else

	f, err := os.Create("C:\\Users\\wds85\\OneDrive\\Documents\\My Games\\Path of Exile\\deadeye_1.filter")
	lib.LogFatal(err)
	defer f.Close()
	f.WriteString(filter.Text)
}
