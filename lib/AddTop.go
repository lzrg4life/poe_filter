package lib

func AddTop(f *Filter) {
	f.Show()
	f.SetFontSize("40")
	f.Continue()

	f.Show()
	f.Class("Quest")

	f.Show()
	f.Rarity("", "Unique")
	f.MinimapIcon("0", "Orange", "Star")
	f.PlayEffect("Orange")

	f.Show()
	f.Sockets("6")
	f.SetBorderColor("255", "255", "255")
	f.MinimapIcon("0", "Brown", "Diamond")
	f.PlayAlertSound("2", "150")

	f.Show()
	f.LinkedSockets(">=", "5")
	f.SetBorderColor("255", "255", "255")
	f.MinimapIcon("0", "Blue", "Diamond")
	f.PlayAlertSound("2", "150")

	f.Show()
	f.Class("Map")
	f.PlayEffect("Grey")
	f.MinimapIcon("0", "Grey", "Circle")

	f.Hide()
	f.AreaLevel(">", "68")
	f.BaseType("\"Scroll of Wisdom\" \"Portal Scroll\"")
	f.StackSize("<", "2")

	f.Show()
	f.BaseType("\"Scroll of Wisdom\" \"Portal Scroll\"")

	f.Show()
	f.BaseType("Chromatic Glassblower")
	f.SetBorderColor("0", "200", "0")
	f.SetTextColor("0", "200", "0")
	f.PlayEffect("Green")

	f.Show()
	f.BaseType("Instilling Chaos")
	f.SetBorderColor("255", "255", "0")
	f.SetTextColor("255", "255", "0")
	f.PlayEffect("Yellow")

	f.Show()
	f.Class("Currency")
	f.SetBorderColor("0", "200", "0")
	f.SetTextColor("0", "200", "0")

	f.Hide()
	f.Class("Divination")
	f.BaseType("\"The Carrion Crow\" \"Lantador's Lost Love\"")

	f.Show()
	f.Class("Divination")
	f.MinimapIcon("0", "Grey", "Square")

	f.Show()
	f.Class("Gem")
	f.BaseType("Vaal")

	f.Show()
	f.Class("Gem")
	f.Quality(">=", "10")

	f.Show()
	f.Class("Gem")
	f.AreaLevel("<", "12")

	f.Hide()
	f.Class("Gem")
	f.AreaLevel(">=", "12")
	f.Quality("<", "1")

	// Armour for crafting
	f.Show()
	f.AreaLevel("<", "73")
	f.BaseType("\"Eternal Burgonet\"")
	f.PlayEffect("White")

	f.Show()
	f.AreaLevel("<", "78")
	f.BaseType("\"Titan Plate\" \"Titan Gauntlets\" \"Titan Greaves\" \"General's Helmet\"")
	f.PlayEffect("White")

	f.Show()
	f.AreaLevel("<", "84")
	f.BaseType("\"Legion Plate\" \"Precursor Gauntlets\" \"Precursor Greaves\" \"Conqueror's Helmet\"")
	f.PlayEffect("White")

	f.Show()
	f.BaseType("\"Royal Plate\" \"Leviathan Gauntlets\" \"Leviathan Greaves\" \"Giantslayer Helmet\"")
	f.PlayEffect("White")

	// Evasion for crafting
	f.Show()
	f.AreaLevel("<", "73")
	f.BaseType("\"Lion Pelt\"")
	f.PlayEffect("White")

	f.Show()
	f.AreaLevel("<", "78")
	f.BaseType("\"Supreme Leather\" \"Slink Gloves\" \"Slink Boots\" \"Dire Pelt\"")
	f.PlayEffect("White")

	f.Show()
	f.AreaLevel("<", "84")
	f.BaseType("\"Astral Leather\" \"Harpyskin Gloves\" \"Harpyskin Boots\" \"Grizzly Pelt\"")
	f.PlayEffect("White")

	f.Show()
	f.BaseType("\"Syndicate's Garb\" \"Velour Gloves\" \"Velour Boots\" \"Stormrider Boots\" \"Majestic Pelt\"")
	f.PlayEffect("White")

	// Energy Shield for crafting
	f.Show()
	f.AreaLevel("<", "73")
	f.BaseType("\"Hubris Circlet\"")
	f.PlayEffect("White")

	f.Show()
	f.AreaLevel("<", "78")
	f.BaseType("\"Arcane Vestment\" \"Sorcerer Gloves\" \"Sorcerer Boots\" \"Moonlit Circlet\"")
	f.PlayEffect("White")

	f.Show()
	f.AreaLevel("<", "84")
	f.BaseType("\"Nightweave Robe\" \"Sage Gloves\" \"Sage Slippers\" \"Sunfire Circlet\"")
	f.PlayEffect("White")

	f.Show()
	f.BaseType("\"Twilight Regalia\" \"Warlock Gloves\" \"Fingerless Silk Gloves\" \"Dreamquest Slippers\" \"Warlock Boots\" \"Lich's Circlet\"")
	f.PlayEffect("White")

	// Armour/Evasion for crafting
	f.Show()
	f.AreaLevel("<", "73")
	f.BaseType("\"Nightmare Bascinet\"")
	f.PlayEffect("White")

	f.Show()
	f.AreaLevel("<", "78")
	f.BaseType("\"Full Wyvernscale\" \"Dragonscale Gauntlets\" \"Dragonscale Boots\" \"Knight Helm\"")
	f.PlayEffect("White")

	f.Show()
	f.AreaLevel("<", "84")
	f.BaseType("\"Marshall's Brigandine\" \"Chimerascale Gauntlets\" \"Chimerascale Boots\" \"Conquest Helmet\"")
	f.PlayEffect("White")

	f.Show()
	f.BaseType("\"Conquest Lamellar\" \"Wyvernscale Gauntlets\" \"Wyvernscale Boots\" \"Haunted Bascinet\"")
	f.PlayEffect("White")

	// Armour/Energy Shield for crafting
	f.Show()
	f.AreaLevel("<", "73")
	f.BaseType("\"Praetor Crown\"")
	f.PlayEffect("White")

	f.Show()
	f.AreaLevel("<", "78")
	f.BaseType("\"Grand Ringmail\" \"Crusader Gloves\" \"Crusader Boots\" \"Faithful Helmet\"")
	f.PlayEffect("White")

	f.Show()
	f.AreaLevel("<", "84")
	f.BaseType("\"Paladin's Hauberk\" \"Martyr Gloves\" \"Martyr Boots\" \"Paladin Crown\"")
	f.PlayEffect("White")

	f.Show()
	f.BaseType("\"Sacred Chainmail\" \"Paladin Gloves\" \"Apothecary's Gloves\" \"Paladin Boots\" \"Divine Crown\" \"Archdemon Crown\" \"Bone Helmet\"")
	f.PlayEffect("White")

	// Evasion/Energy Shield Shield for crafting
	f.Show()
	f.AreaLevel("<", "73")
	f.BaseType("\"Deicide Mask\"")
	f.PlayEffect("White")

	f.Show()
	f.AreaLevel("<", "78")
	f.BaseType("\"Sanguine Raiment\" \"Murder Mitts\" \"Murder Boots\" \"Jester Mask\"")
	f.PlayEffect("White")

	f.Show()
	f.AreaLevel("<", "84")
	f.BaseType("\"Torturer Garb\" \"Infiltrator Mitts\" \"Infiltrator Boots\" \"Ancient Mask\"")
	f.PlayEffect("White")

	f.Show()
	f.BaseType("\"Necrotic Armour\" \"Phantom Mitts\" \"Phantom Boots\" \"Fugitive Boots\" \"Torturer's Mask\" \"Blizzard Crown\"")
	f.PlayEffect("White")

	// defence for crafting
	f.Show()
	f.BaseType("\"Grasping Mail\" \"Sacrificial Garb\" Runic")
	f.PlayEffect("Orange")
	f.MinimapIcon("0", "Orange", "Star")
}
