package tui

import (
	"errors"
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var Register *[]string = &[]string{}
var ErrStop = errors.New("Stop Iteration")

type Iterator struct {
	series []string
}

func (iter *Iterator) Next() (string, error) {
	if len(iter.series) == 0 {
		return "", ErrStop
	}
	currentValue := iter.series[0]
	iter.series = iter.series[1:]
	return currentValue, nil
}

func (iter *Iterator) Iteration() string {
	color, err := iter.Next()
	if err != nil {
		return ""
	}
	return color
}
func (iter *Iterator) New(iterable []string) {
	iter.series = iterable
}

type baseColor struct {
	Iterator
	color string
}

func (c *baseColor) Registration(hexColor string) *baseColor {
	c.color = hexColor
	*Register = append(*Register, hexColor)
	return c
}
func (c *baseColor) Get() string {
	return c.color
}
func (c *baseColor) GetColor() *string {
	return &c.color
}

func (c *baseColor) style() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(lipgloss.Color(c.color))
}

func (c *baseColor) Print(str string) {
	fmt.Printf("%s\n", c.style().Render(str))
}
func (c *baseColor) Iter() string {
	iter := c.CreateIterator()
	return iter.Iteration()
}

func (c *baseColor) CreateIterator() *Iterator {
	iter := &Iterator{
		series: *Register,
	}
	return iter
}

type Purple struct {
	baseColor
}
type Pink struct {
	baseColor
}
type Violet struct {
	baseColor
}
type White struct {
	baseColor
}
type Green struct {
	baseColor
}
type Red struct {
	baseColor
}
type Black struct {
	baseColor
}
type Blue struct {
	baseColor
}
type Cyan struct {
	baseColor
}

type Colours struct {
	Register *[]string
	baseColor
	Purple
	Pink
	Violet
	White
	Green
	Red
	Black
	Blue
	Cyan
}

func NewColours() *Colours {
	c := &Colours{}
	c.Register = Register

	c.Purple.Registration("#874BFD")
	c.Violet.Registration("#7d56f4")
	c.Pink.Registration("#A020F0")
	c.White.Registration("#FFFFFF")
	c.Green.Registration("#16D90F")
	c.Red.Registration("#E31415")
	c.Black.Registration("#000000")
	c.Blue.Registration("#2720f0")
	c.Cyan.Registration("#00FFFF")
	return c
}
