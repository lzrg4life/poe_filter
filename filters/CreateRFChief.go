package filters

import (
	"os"
	"poefilter/lib"
)

func CreateRFChief() {
	var filter lib.Filter

	lib.AddTop(&filter)
	lib.AddLeveling(&filter)
	lib.AddPhase1Vendoring(&filter)

	// Gem: (BBB, BBR) RF, Efficacy, Elemental Damage, Burning Damage, Area of Effect
	// Gear: Armour/ES gloves/helmet/boots, pure armour chest
	// Jewelry: Fire rings and Chaos rings
	// Weapon: Sceptres
	// Shield: armour shield
	sockets := "BBB BBR RRG RGG GBR"
	filter.ShowArmourESBoots(sockets)
	filter.ShowArmourESGloves(sockets)
	filter.ShowArmourESHelmets(sockets)
	filter.ShowArmourChests("")
	filter.ShowArmourShields("")
	filter.ShowAmythestRings()
	filter.ShowRubyRings()
	filter.ShowLifeFlasks()

	filter.Show()
	filter.Class("Sceptre")

	lib.AddHideStuff(&filter)

	filter.Show() // show everything else

	f, err := os.Create("C:\\Users\\wds85\\OneDrive\\Documents\\My Games\\Path of Exile\\rf_chief.filter")
	lib.LogFatal(err)
	defer f.Close()
	f.WriteString(filter.Text)
}
