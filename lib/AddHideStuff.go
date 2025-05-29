package lib

func AddHideStuff(f *Filter) {
	f.Show()
	f.AreaLevel("", "1")

	f.Hide()
	f.BaseType("\"Life Flask\" \"Mana Flask\" \"Hybrid Flask\"")

	f.Hide()
	f.Class("Amulets Rings Belts")

	f.Hide()
	f.Class("Bows Staves \"Two Hand\" \"Body Armours\" Warstaves")

	f.Hide()
	f.Class("Gloves Boots Helmets")

	f.Hide()
	f.Class("Claws Dagger Wands \"One Hand\" Shields Sceptres")

	f.Hide()
	f.Class("Quivers")
}
