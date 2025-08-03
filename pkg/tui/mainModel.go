package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	identRightEdge int = 5
	identHight     int = 6
)

type Tab string
type Items []list.Item
type TabsItems map[Tab]Items

type Storage interface {
	GetTabsItems() TabsItems
	GetTabs() []Tab
	DownloadItems(tab Tab, item Item) Item
}

type ModelTabs struct {
	Storage
	windows   *Windows
	Keys      *ListKeyMap
	Tabs      []Tab
	TabsItems TabsItems
	list      list.Model
	help      help.Model
	activeTab int
	width     int
	height    int
}

func (m ModelTabs) Init() tea.Cmd {
	return nil
}
func (m *ModelTabs) initList() {
	tabKey := m.Tabs[0]          // [Tab1, Tab2]
	items := m.TabsItems[tabKey] // {Tab1: [Item1, Item2], Tab2: [Item3, Item4]}
	m.list = list.New(items, list.NewDefaultDelegate(), 0, 0)
	m.list.SetShowTitle(false)
}
func (m *ModelTabs) setListForCurrentTab() {
	tabKey := m.Tabs[m.activeTab] // [Tab1, Tab2]
	items := m.TabsItems[tabKey]  // {Tab1: [Item1, Item2], Tab2: [Item3, Item4]}
	m.list.SetItems(items)
	m.list.SetShowTitle(false)
}
func (m *ModelTabs) nextTab() {
	m.activeTab = min(m.activeTab+1, len(m.Tabs)-1)
	m.setListForCurrentTab()
}
func (m *ModelTabs) prevTab() {
	m.activeTab = max(m.activeTab-1, 0)
	m.setListForCurrentTab()
}

func (m *ModelTabs) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.Keys.Exit):
			return m, tea.Quit
		case key.Matches(msg, m.Keys.NextTab):
			m.nextTab()
		case key.Matches(msg, m.Keys.PrevTab):
			m.prevTab()
		case key.Matches(msg, m.Keys.Download):
			tab := m.Tabs[m.activeTab] // [Tab1, Tab2]
			// HACK: convert list.Item to tui.Item
			currentItem := m.list.SelectedItem().(Item)
			newItem := m.DownloadItems(tab, currentItem)
			m.list.VisibleItems()[m.list.Index()] = newItem
			m.list.SetItems(m.list.VisibleItems())
		}
	case tea.WindowSizeMsg:
		h, v := m.windows.DocStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h-identHight, msg.Height-v-identRightEdge)
		m.width = msg.Width
		m.height = msg.Height
	}
	// m.TabContent[m.activeTab] = fmt.Sprintf("Height: %d, Width: %d", m.height, m.width)

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m ModelTabs) View() string {
	doc := strings.Builder{}
	var renderedTabs []string
	var style lipgloss.Style
	help := m.help.ShortHelpView([]key.Binding{
		m.Keys.NextTab,
		m.Keys.PrevTab,
		m.Keys.Exit,
		m.Keys.HelpMenu,
	})

	for i, t := range m.Tabs {
		isFirst, isLast, isActive := i == 0, i == len(m.Tabs)-1, i == m.activeTab
		if isActive {
			style = m.windows.ActiveTabStyle
		} else {
			style = m.windows.InactiveTabStyle
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
		renderedTabs = append(renderedTabs, style.Render(string(t)))
	}
	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	fillerStringLen := m.width - lipgloss.Width(row) - identRightEdge
	if fillerStringLen > 0 {
		fillerString := strings.Repeat("─", fillerStringLen+1)
		fillerString += "┐"
		row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, m.windows.FillerStyle.Render(fillerString))
	}

	// -5 от правого края к левому!!!!
	windowStyle := m.windows.Style.Width(m.width - identRightEdge).
		Height(m.height - identHight)
	doc.WriteString(row + "\n")
	// doc.WriteString(windowStyle.Render(m.TabContent[m.activeTab]))
	doc.WriteString(windowStyle.Render(m.list.View()))
	doc.WriteString("\n" + help)
	return m.windows.DocStyle.Render(doc.String())
}

func NewModelTabs(s Storage) *ModelTabs {
	var window = NewWindows()
	var listKeys = NewListKeyMap()

	m := ModelTabs{
		Storage:   s,
		windows:   window,
		Keys:      listKeys,
		Tabs:      s.GetTabs(),
		TabsItems: s.GetTabsItems(),
		help:      help.New(),
	}
	m.initList()
	return &m
}

// func TestModelTabs() *ModelTabs {
// 	return NewModelTabs(testTabsItems{})
// }
