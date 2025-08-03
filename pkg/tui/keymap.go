package tui

import "github.com/charmbracelet/bubbles/key"

type ListKeyMap struct {
	NextTab  key.Binding
	PrevTab  key.Binding
	HelpMenu key.Binding
	Exit     key.Binding
	Download key.Binding
}

func NewListKeyMap() *ListKeyMap {
	return &ListKeyMap{
		Exit: key.NewBinding(
			key.WithKeys("ctrl+c", "ctrl+d", "q", "esc"),
			key.WithHelp("q", "exit"),
		),
		NextTab: key.NewBinding(
			key.WithKeys("right", "l", "tab"),
			key.WithHelp("->/l/tab", "move to right tab"),
		),
		PrevTab: key.NewBinding(
			key.WithKeys("left", "h", "p", "shift+tab"),
			key.WithHelp("l", "move to left tab"),
		),
		HelpMenu: key.NewBinding(
			key.WithKeys("H"),
			key.WithHelp("H", "toggle help"),
		),
		Download: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "Download file"),
		),
	}
}
