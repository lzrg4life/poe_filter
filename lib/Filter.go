package lib

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

func (f *Filter) SocketGroup(comparison string, group string) {
	f.Text += tab + "SocketGroup " + comparison + " " + group + newLine
}
