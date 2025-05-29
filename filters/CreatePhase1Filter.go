package filters

import (
	"os"
	"poefilter/lib"
)

func CreatePhase1Filter() {
	var filter lib.Filter

	lib.AddTop(&filter)
	lib.AddLeveling(&filter)
	lib.AddPhase1Vendoring(&filter)
	lib.AddHideStuff(&filter)

	filter.Show() // show everything else

	f, err := os.Create("C:\\Users\\wds85\\OneDrive\\Documents\\My Games\\Path of Exile\\phase_1.filter")
	lib.LogFatal(err)
	defer f.Close()
	f.WriteString(filter.Text)
}
