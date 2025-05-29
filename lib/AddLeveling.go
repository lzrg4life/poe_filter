package lib

func AddLeveling(f *Filter) {

	f.Show()
	f.AreaLevel("<", "14")
	f.BaseType("\"Life Flask\" \"Mana Flask\"")

	f.Show()
	f.AreaLevel("<", "14")
	f.Class("Ring Amulet Belt")

	f.Show()
	f.AreaLevel("<", "68")
	f.Rarity("", "Rare")
	f.Height("<", "4")

}
