package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	identRightEdge int = 5
	identHight     int = 4
)

type ModelTabs struct {
	*Windows
	list       list.Model
	Tabs       []string
	TabContent []string
	Keys       *ListKeyMap
	activeTab  int
	width      int
	height     int
}

func (m ModelTabs) Init() tea.Cmd {
	return nil
}
func (m *ModelTabs) nextTab() {
	m.activeTab = min(m.activeTab+1, len(m.Tabs)-1)
}

func (m *ModelTabs) prevTab() {
	m.activeTab = max(m.activeTab-1, 0)
}
func (m ModelTabs) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.Keys.Exit):
			return m, tea.Quit
		case key.Matches(msg, m.Keys.NextTab):
			m.nextTab()
		case key.Matches(msg, m.Keys.PrevTab):
			m.prevTab()
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	m.TabContent[m.activeTab] = fmt.Sprintf("Height: %d, Width: %d", m.height, m.width)

	return m, nil
}

func (m ModelTabs) View() string {
	doc := strings.Builder{}
	var renderedTabs []string
	var style lipgloss.Style

	for i, t := range m.Tabs {
		isFirst, isLast, isActive := i == 0, i == len(m.Tabs)-1, i == m.activeTab
		if isActive {
			style = m.ActiveTabStyle
		} else {
			style = m.InactiveTabStyle
		}
		border, _, _, _, _ := style.GetBorder()
		if isFirst && isActive {
			border.BottomLeft = "│"
		} else if isFirst && !isActive {
			border.BottomLeft = "├"
		} else if isLast && isActive {
			border.BottomRight = "└"
		} else if isLast && !isActive {
			border.BottomRight = "┴"
		}
		style = style.Border(border)
		renderedTabs = append(renderedTabs, style.Render(t))
	}
	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	fillerStringLen := m.width - lipgloss.Width(row) - identRightEdge
	if fillerStringLen > 0 {
		fillerString := strings.Repeat("─", fillerStringLen+1)
		fillerString += "┐"
		row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, m.FillerStyle.Render(fillerString))
	}

	// -5 от правого края к левому!!!!
	windowStyle := m.Style.Width(m.width - identRightEdge).
		Height(m.height - identHight)
	doc.WriteString(row + "\n")
	doc.WriteString(windowStyle.Render(m.TabContent[m.activeTab]))
	return m.DocStyle.Render(doc.String())
}

func NewModelTabs(tabs, tabsContent []string) *ModelTabs {
	var window = NewWindows()
	var listKeys = NewListKeyMap()
	return &ModelTabs{Windows: window, Keys: listKeys, Tabs: tabs, TabContent: tabsContent}
}
func TestModelTabs() *ModelTabs {
	tabs := []string{"Lip Gloss", "Blush", "Eye Shadow", "Mascara", "Foundation"}
	tabsContent := []string{"Lip Gloss Tab", "Blush Tab", "Eye Shadow Tab", "Mascara Tab", "Foundation Tab"}
	return NewModelTabs(tabs, tabsContent)
}
