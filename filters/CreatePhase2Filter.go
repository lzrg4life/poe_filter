package filters

import (
	"os"
	"poefilter/lib"
)

func CreatePhase2Filter() {
	var filter lib.Filter

	lib.AddTop(&filter)
	lib.AddLeveling(&filter)
	lib.AddPhase2Vendoring(&filter)
	lib.AddHideStuff(&filter)

	filter.Show() // show everything else

	f, err := os.Create("C:\\Users\\wds85\\OneDrive\\Documents\\My Games\\Path of Exile\\phase_2.filter")
	lib.LogFatal(err)
	defer f.Close()
	f.WriteString(filter.Text)
}
