package cmd

import (
	"github.com/charmbracelet/lipgloss"
)

var KeyForeground = lipgloss.AdaptiveColor{Light: "#8D8D8D", Dark: "#5F5F5F"}
var ValForeground1 = lipgloss.AdaptiveColor{Light: "#07C592", Dark: "#06D6A0"}
var ValForeground2 = lipgloss.AdaptiveColor{Light: "#07C592", Dark: "#50D8D7"}
var TitleForeground = lipgloss.Color("#00BCEB")

var ErrorStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#DF2935")).
	PaddingTop(1).
	PaddingBottom(1).
	PaddingLeft(2).
	PaddingRight(2)

var Container = lipgloss.NewStyle().Padding(1)

var TitleStyle = lipgloss.NewStyle().Foreground(TitleForeground)
var DescriptionStyle = lipgloss.NewStyle().Foreground(KeyForeground)

var KeyStyle = lipgloss.NewStyle().
	Foreground(KeyForeground).
	BorderStyle(lipgloss.NormalBorder()).
	BorderBottom(true).
	BorderBottomForeground(KeyForeground).Width(24)

var ValStyle1 = lipgloss.NewStyle().Foreground(ValForeground1).PaddingLeft(1).Bold(true)

var ValStyle2 = lipgloss.NewStyle().Foreground(ValForeground2).PaddingLeft(1).Bold(true)
