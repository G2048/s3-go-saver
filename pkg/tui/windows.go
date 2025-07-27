package tui

import "github.com/charmbracelet/lipgloss"

type WindowDrawing struct {
	LightColor string
	DarkColor  string
}

type Windows struct {
	*WindowDrawing
	ActiveTabStyle   lipgloss.Style
	InactiveTabStyle lipgloss.Style
	Style            lipgloss.Style
	DocStyle         lipgloss.Style
	FillerStyle      lipgloss.Style
}

func (s *Windows) ChangeBorderColor() {
	/* Usage:
	      mt.LightColor = mt.Black.Get()
	      mt.DarkColor = mt.Blue.Get()
	      mt.ChangeBorderColor()
	   OR:
	       mt.DarkColor = (*mt.Register)[mt.currentColor]
	       mt.ChangeBorderColor()
	*/
	s.ActiveTabStyle = s.WindowDrawing.ActiveTabStyle()
	s.InactiveTabStyle = s.WindowDrawing.InactiveTabStyle()
	s.Style = s.WindowDrawing.Style()
	s.DocStyle = s.WindowDrawing.DocStyle()
}

func NewWindows() *Windows {
	colors := NewColours()
	w := &WindowDrawing{
		LightColor: colors.Black.Get(),
		DarkColor:  colors.Violet.Get(),
	}
	return &Windows{
		WindowDrawing:    w,
		ActiveTabStyle:   w.ActiveTabStyle(),
		InactiveTabStyle: w.InactiveTabStyle(),
		Style:            w.Style(),
		DocStyle:         w.DocStyle(),
		FillerStyle:      w.FillerStyle(),
	}
}

func (wd *WindowDrawing) highlightColor() lipgloss.AdaptiveColor {
	return lipgloss.AdaptiveColor{Light: wd.LightColor, Dark: wd.DarkColor}
}
func (wd *WindowDrawing) FillerStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(wd.highlightColor())
}
func (wd *WindowDrawing) DocStyle() lipgloss.Style {
	return lipgloss.NewStyle().Padding(0, 1, 0, 2)
}

// Height
// Core Area
func (wd *WindowDrawing) Style() lipgloss.Style {
	return lipgloss.NewStyle().BorderForeground(wd.highlightColor()).
		Align(lipgloss.Center).         // allign Text in area
		AlignVertical(lipgloss.Center). // allign Text in area
		Border(lipgloss.RoundedBorder()).
		Padding(0, 2).
		UnsetBorderTop()
}

// Weight
// Tabs
func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}
func (wd *WindowDrawing) inactiveTabBorder() lipgloss.Border {
	return tabBorderWithBottom("┴", "─", "┴")
}
func (wd *WindowDrawing) InactiveTabStyle() lipgloss.Style {
	inactiveTabBorder := tabBorderWithBottom("┴", "─", "┴")

	return lipgloss.NewStyle().
		Border(inactiveTabBorder, true).
		BorderForeground(wd.highlightColor()).
		Padding(0, 1)
}
func (wd *WindowDrawing) ActiveTabStyle() lipgloss.Style {
	activeTabBorder := tabBorderWithBottom("┘", " ", "└")
	return wd.InactiveTabStyle().Border(activeTabBorder, true).Underline(true)
}
