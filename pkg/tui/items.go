package tui

import (
	"github.com/charmbracelet/bubbles/list"
)

type Item struct {
	Top, Desc string
}

func (i Item) Title() string       { return i.Top }
func (i Item) Description() string { return i.Desc }
func (i Item) FilterValue() string { return i.Top }

func NewItems() []list.Item {
	return []list.Item{
		Item{Top: "Raspberry Pi’s", Desc: "I have ’em all over my house"},
		Item{Top: "Nutella", Desc: "It's good on toast"},
		Item{Top: "Bitter melon", Desc: "It cools you down"},
		Item{Top: "Nice socks", Desc: "And by that I mean socks without holes"},
		Item{Top: "Eight hours of sleep", Desc: "I had this once"},
		Item{Top: "Cats", Desc: "Usually"},
		Item{Top: "Plantasia, the album", Desc: "My plants love it too"},
		Item{Top: "Pour over coffee", Desc: "It takes forever to make though"},
		Item{Top: "VR", Desc: "Virtual reality...what is there to say?"},
		Item{Top: "Noguchi Lamps", Desc: "Such pleasing organic forms"},
		Item{Top: "Linux", Desc: "Pretty much the best OS"},
		Item{Top: "Business school", Desc: "Just kidding"},
		Item{Top: "Pottery", Desc: "Wet clay is a great feeling"},
		Item{Top: "Shampoo", Desc: "Nothing like clean hair"},
		Item{Top: "Table tennis", Desc: "It’s surprisingly exhausting"},
		Item{Top: "Milk crates", Desc: "Great for packing in your extra stuff"},
		Item{Top: "Afternoon tea", Desc: "Especially the tea sandwich part"},
		Item{Top: "Stickers", Desc: "The thicker the vinyl the better"},
		Item{Top: "20° Weather", Desc: "Celsius, not Fahrenheit"},
		Item{Top: "Warm light", Desc: "Like around 2700 Kelvin"},
		Item{Top: "The vernal equinox", Desc: "The autumnal equinox is pretty good too"},
		Item{Top: "Gaffer’s tape", Desc: "Basically sticky fabric"},
		Item{Top: "Terrycloth", Desc: "In other words, towel fabric"},
	}
}
