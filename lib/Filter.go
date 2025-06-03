package lib

import (
	"fmt"
	"strconv"
)

var newLine = "\n"
var tab = "\t"

type Filter struct {
	Text string
}

func (f *Filter) Show() {
	f.Text += "Show" + newLine
}

func (f *Filter) Hide() {
	f.Text += "Hide" + newLine
}

func (f *Filter) AddOption(text string) {
	f.Text += tab + text + newLine
}

func (f *Filter) SetFontSize(size string) {
	f.Text += tab + "SetFontSize " + size + newLine
}

func (f *Filter) Continue() {
	f.Text += tab + "Continue" + newLine
}

func (f *Filter) Class(class string) {
	f.Text += tab + "Class " + class + newLine
}

func (f *Filter) Rarity(comparison string, rarity string) {
	f.Text += tab + "Rarity " + comparison + " " + rarity + newLine
}

func (f *Filter) MinimapIcon(id string, color string, shape string) {
	f.Text += tab + "MinimapIcon " + id + " " + color + " " + shape + newLine
}

func (f *Filter) PlayEffect(color string) {
	f.Text += tab + "PlayEffect " + color + newLine
}

func (f *Filter) Sockets(count string) {
	f.Text += tab + "Sockets " + count + newLine
}

func (f *Filter) LinkedSockets(comparison string, count string) {
	f.Text += tab + "LinkedSockets " + comparison + " " + count + newLine
}

func (f *Filter) SetBorderColor(r255 string, g255 string, b255 string) {
	f.Text += tab + "SetBorderColor " + r255 + " " + g255 + " " + b255 + newLine
}

func (f *Filter) SetTextColor(r255 string, g255 string, b255 string) {
	f.Text += tab + "SetTextColor " + r255 + " " + g255 + " " + b255 + newLine
}

func (f *Filter) SetBackgroundColor(r255 string, g255 string, b255 string) {
	f.Text += tab + "SetBackgroundColor " + r255 + " " + g255 + " " + b255 + newLine
}

func (f *Filter) PlayAlertSound(id string, volume string) {
	f.Text += tab + "PlayAlertSound " + id + " " + volume + newLine
}

func (f *Filter) AreaLevel(comparison string, level string) {
	f.Text += tab + "AreaLevel " + comparison + " " + level + newLine
}

func (f *Filter) BaseType(text string) {
	f.Text += tab + "BaseType " + text + newLine
}

func (f *Filter) StackSize(comparison string, size string) {
	f.Text += tab + "StackSize " + comparison + " " + size + newLine
}

func (f *Filter) Quality(comparison string, size string) {
	f.Text += tab + "Quality " + comparison + " " + size + newLine
}

func (f *Filter) Height(comparison string, size string) {
	f.Text += tab + "Height " + comparison + " " + size + newLine
}

func (f *Filter) Width(comparison string, size string) {
	f.Text += tab + "Width " + comparison + " " + size + newLine
}

// Requires the sockets to be linked
func (f *Filter) SocketGroup(comparison string, group string) {
	f.Text += tab + "SocketGroup " + comparison + " " + group + newLine
}

func (f *Filter) addArmour(name string, belowAreaLevel int, socketgroup string) {
	f.Show()
	f.BaseType(fmt.Sprintf("\"%s\"", name))
	f.AreaLevel("<", strconv.Itoa(belowAreaLevel))

	if socketgroup != "" {
		f.SocketGroup("=", socketgroup)
	}
}

func (f *Filter) ShowArmourESGloves(socketgroup string) {
	f.addArmour("Chain Gloves", 19, socketgroup)
	f.addArmour("Ringmail Gloves", 32, socketgroup)
	f.addArmour("Mesh Gloves", 37, socketgroup)
	f.addArmour("Riveted Gloves", 43, socketgroup)
	f.addArmour("Zealot Gloves", 51, socketgroup)
	f.addArmour("Soldier Gloves", 57, socketgroup)
	f.addArmour("Legion Gloves", 66, socketgroup)
	f.addArmour("Crusader Gloves", 78, socketgroup)
	f.addArmour("Martyr Gloves", 84, socketgroup)
	f.BaseType("\"Paladin Gloves\"")

	// These are special
	f.Show()
	f.BaseType("\"Apothecary's Gloves\"")
}

func (f *Filter) ShowArmourESBoots(socketgroup string) {
	f.addArmour("Chain Boots", 13, socketgroup)
	f.addArmour("Ringmail Boots", 28, socketgroup)
	f.addArmour("Mesh Boots", 36, socketgroup)
	f.addArmour("Riveted Boots", 49, socketgroup)
	f.addArmour("Zealot Boots", 49, socketgroup)
	f.addArmour("Soldier Boots", 58, socketgroup)
	f.addArmour("Legion Boots", 64, socketgroup)
	f.addArmour("Crusader Boots", 78, socketgroup)
	f.addArmour("Martyr Boots", 84, socketgroup)
	f.BaseType("\"Paladin Boots\"")

	// These are special
	f.Show()
	f.BaseType("\"Two-Toned Boots\"")
}

func (f *Filter) ShowArmourESHelmets(socketgroup string) {
	f.addArmour("Rusted Coif", 12, socketgroup)
	f.addArmour("Soldier Helmet", 22, socketgroup)
	f.addArmour("Great Helmet", 31, socketgroup)
	f.addArmour("Crusader Helmet", 37, socketgroup)
	f.addArmour("Aventail Helmet", 44, socketgroup)
	f.addArmour("Zealot Helmet", 53, socketgroup)
	f.addArmour("Great Crown", 58, socketgroup)
	f.addArmour("Magistrate Crown", 63, socketgroup)
	f.addArmour("Prophet Crown", 68, socketgroup)
	f.addArmour("Praetor Crown", 73, socketgroup)
	f.addArmour("Faithful Helmet", 78, socketgroup)
	f.addArmour("Paladin Crown", 84, socketgroup)
	f.BaseType("\"Divine Crown\"")

	// These are special
	f.Show()
	f.BaseType("\"Bone Helmet\"")
	f.BaseType("\"Archdemon Crown\"")
}

func (f *Filter) ShowArmourChests(socketgroup string) {
	f.addArmour("Plate Vest", 6, socketgroup)
	f.addArmour("Chestplate", 17, socketgroup)
	f.addArmour("Copper Plate", 21, socketgroup)
	f.addArmour("War Plate", 28, socketgroup)
	f.addArmour("Full Plate", 32, socketgroup)
	f.addArmour("Arena Plate", 35, socketgroup)
	f.addArmour("Lordly Plate", 37, socketgroup)
	f.addArmour("Bronze Plate", 41, socketgroup)
	f.addArmour("Battle Plate", 45, socketgroup)
	f.addArmour("Sun Plate", 49, socketgroup)
	f.addArmour("Colosseum Plate", 53, socketgroup)
	f.addArmour("Majestic Plate", 56, socketgroup)
	f.addArmour("Golden Plate", 59, socketgroup)
	f.addArmour("Crusader Plate", 62, socketgroup)
	f.addArmour("Astral Plate", 65, socketgroup)
	f.addArmour("Gladiator Plate", 68, socketgroup)
	f.addArmour("Titan Plate", 78, socketgroup)
	f.addArmour("Legion Plate", 84, socketgroup)
	f.BaseType("\"Royal Plate\"")
}

func (f *Filter) ShowArmourShields(socketgroup string) {
	f.addArmour("Splintered Tower Shield", 5, socketgroup)
	f.addArmour("Corroded Tower Shield", 11, socketgroup)
	f.addArmour("Rawhide Tower Shield", 17, socketgroup)
	f.addArmour("Cedar Tower Shield", 24, socketgroup)
	f.addArmour("Copper Tower Shield", 30, socketgroup)
	f.addArmour("Reinforced Tower Shield", 35, socketgroup)
	f.addArmour("Painted Tower Shield", 39, socketgroup)
	f.addArmour("Buckskin Tower Shield", 43, socketgroup)
	f.addArmour("Mahogany Tower Shield", 47, socketgroup)
	f.addArmour("Bronze Tower Shield", 51, socketgroup)
	f.addArmour("Girded Tower Shield", 55, socketgroup)
	f.addArmour("Crested Tower Shield", 58, socketgroup)
	f.addArmour("Shagreen Tower Shield", 61, socketgroup)
	f.addArmour("Ebony Tower Shield", 64, socketgroup)

	// All three of these might be viable
	f.Show()
	f.BaseType("\"Ezomyte Tower Shield\"")  // more life
	f.BaseType("\"Colossal Tower Shield\"") // more armour
	f.BaseType("\"Pinnacle Tower Shield\"") // midway life/armour

	// Special
	f.Show()
	f.BaseType("\"Magmatic Tower Shield\"")
	f.BaseType("\"Heat-attuned Tower Shield\"")
}

func (f *Filter) ShowAmythestRings() {
	f.Show()
	f.Class("Amethyst Ring")
}

func (f *Filter) ShowRubyRings() {
	f.Show()
	f.Class("Ruby Ring")
}

func (f *Filter) addLifeFlask(name string, belowAreaLevel int) {
	f.Show()
	f.BaseType(fmt.Sprintf("\"%s\"", name))
	f.AreaLevel("<", strconv.Itoa(belowAreaLevel))
}

func (f *Filter) ShowLifeFlasks() {
	f.addLifeFlask("Small Life Flask", 4)
	f.addLifeFlask("Medium Life Flask", 8)
	f.addLifeFlask("Large Life Flask", 13)
	f.addLifeFlask("Greater Life Flask", 18)
	f.addLifeFlask("Grand Life Flask", 24)
	f.addLifeFlask("Giant Life Flask", 30)
	f.addLifeFlask("Colossal Life Flask", 36)
	f.addLifeFlask("Sacred Life Flask", 42)
	f.addLifeFlask("Hallowed Life Flask", 50)
	f.addLifeFlask("Sanctified Life Flask", 60)

	f.Show()
	f.BaseType("\"Divine Life Flask\"")

	// We don't talk about "Enternal Life Flask"
}
