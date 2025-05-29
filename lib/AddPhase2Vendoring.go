package lib

func AddPhase2Vendoring(f *Filter) {

	// For cold resist or fire resist recipe
	f.Show()
	f.AreaLevel("<", "14")
	f.BaseType("\"Iron Ring\"")
	f.Rarity("", "Normal")

	// Vendoring for Chromatic Orbs
	f.Show()
	f.Height("<=", "2")
	f.Width("<=", "2")
	f.SocketGroup(">=", "RGB")
	f.SetBorderColor("0", "0", "0")
	f.SetTextColor("0", "0", "0")
	f.SetBackgroundColor("0", "0", "0")

	// Vendoring for transmutes/alterations

	f.Show()
	f.Class("Ring Amulet Belt")
	f.Rarity("", "Magic Rare")

	f.Show()
	f.Class("Flask")
	f.Rarity("", "Magic")
}
