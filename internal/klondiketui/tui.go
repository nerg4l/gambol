package klondiketui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nerg4l/gambol/card"
	"github.com/nerg4l/gambol/klondike"
)

const deckNone klondike.DeckID = -1

const (
	stateDeckSelection state = iota
	stateCardSelection
)

type state int8

type TUI struct {
	game  klondike.Game
	state state

	focusedRow    int8
	focusedColumn int8
	focusedCard   int8

	selectedRow    int8
	selectedColumn int8

	rows [][]klondike.DeckID
}

func New() TUI {
	t := TUI{
		game: klondike.NewGame(),

		rows: [][]klondike.DeckID{make([]klondike.DeckID, 7), make([]klondike.DeckID, 7)},
	}

	t.rows[0][0] = klondike.DeckS
	t.rows[0][1] = klondike.DeckW
	t.rows[0][2] = deckNone
	t.rows[0][3] = klondike.DeckF1
	t.rows[0][4] = klondike.DeckF2
	t.rows[0][5] = klondike.DeckF3
	t.rows[0][6] = klondike.DeckF4

	for i := int8(0); i < 7; i++ {
		t.rows[1][i] = klondike.DeckT1 + klondike.DeckID(i)
	}

	return t.reset()
}

func (t TUI) reset() TUI {
	t.focusedCard = 0
	t.selectedRow = -1
	t.selectedColumn = -1
	t.state = stateDeckSelection
	return t
}

func (t TUI) Init() tea.Cmd {
	return nil
}

func (t TUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return t, tea.Quit
		case tea.KeyEsc:
			return t.reset(), nil
		}
	}

	if t.state == stateDeckSelection {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
			case tea.KeyRight:
				row := t.rows[t.focusedRow]
				if t.focusedColumn < (int8(len(row)) - 1) {
					t.focusedColumn++
				} else {
					t.focusedColumn = 0
				}
			case tea.KeyLeft:
				row := t.rows[t.focusedRow]
				if t.focusedColumn > 0 {
					t.focusedColumn--
				} else if len(row) > 0 {
					t.focusedColumn = int8(len(row)) - 1
				}
			case tea.KeyDown:
				if t.focusedRow < (int8(len(t.rows)) - 1) {
					t.focusedRow++
				}
			case tea.KeyUp:
				if t.focusedRow > 0 {
					t.focusedRow--
				}
			case tea.KeySpace, tea.KeyEnter:
				d := t.focusedDeckID()
				if d == deckNone {
					break
				}
				if t.selectedDeckID() == deckNone {
					if klondike.IsTableau(d) {
						t.state = stateCardSelection
						t.focusedCard = int8(len(t.game.Deck(d))) - 1
					}
					t.selectedRow = t.focusedRow
					t.selectedColumn = t.focusedColumn
				} else {
					g, err := t.game.Move(t.selectedDeckID(), t.focusedDeckID(), t.selectedCard())
					if err != nil {
						break
					}
					t.game = g
					t = t.reset()
				}
			}
		}
		return t, nil
	}

	if t.state == stateCardSelection {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			d := t.game.Deck(t.selectedDeckID())
			switch msg.Type {
			case tea.KeyDown:
				if t.focusedCard < (int8(len(d)) - 1) {
					t.focusedCard++
				} else {
					t.focusedCard = 0
				}
			case tea.KeyUp:
				if t.focusedCard > 0 {
					t.focusedCard--
				} else if len(d) > 0 {
					t.focusedCard = int8(len(d)) - 1
				}
			case tea.KeySpace, tea.KeyEnter:
				t.state = stateDeckSelection
			}
		}
	}

	return t, nil
}

var (
	colorFocused  = lipgloss.AdaptiveColor{Light: "3", Dark: "11"}
	colorSelected = lipgloss.AdaptiveColor{Light: "2", Dark: "10"}

	styleCellNormal   = lipgloss.NewStyle().Border(lipgloss.NormalBorder())
	styleCellFocused  = styleCellNormal.Copy().BorderForeground(colorFocused)
	styleCellSelected = styleCellNormal.Copy().BorderForeground(colorSelected)

	styleCardNormal  = lipgloss.NewStyle().Padding(0, 1)
	styleCardFocused = lipgloss.NewStyle().Border(func() lipgloss.Border {
		border := lipgloss.NormalBorder()
		border.Left = ">"
		border.Right = "<"
		return border
	}(), false, true).BorderForeground(colorFocused)
	styleCardSelected = styleCardFocused.Copy().BorderForeground(colorSelected)
)

func (t TUI) View() string {
	rows := make([]string, len(t.rows))
	for i := range t.rows {
		row := make([]string, len(t.rows[i]))
		for j, c := range t.rows[i] {
			style := styleCellNormal
			if deckNone != c && t.selectedDeckID() == c {
				style = styleCellSelected
			} else if t.focusedDeckID() == c {
				style = styleCellFocused
			}
			row[j] = style.Render(t.renderDeck(c))
		}
		rows[i] = lipgloss.JoinHorizontal(lipgloss.Top, row...)
	}
	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}

func (t TUI) selectedCard() *card.Card {
	dt := t.selectedDeckID()
	if klondike.IsTableau(dt) {
		return &t.game.Deck(dt)[t.focusedCard]
	}
	return nil
}

func (t TUI) selectedDeckID() klondike.DeckID {
	if t.selectedRow == -1 && t.selectedColumn == -1 {
		return deckNone
	}
	return t.rows[t.selectedRow][t.selectedColumn]
}

func (t TUI) focusedDeckID() klondike.DeckID {
	return t.rows[t.focusedRow][t.focusedColumn]
}

var (
	rankSymbols = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	suitSymbols = map[card.Suit]rune{
		card.SuitSpade:   '♠',
		card.SuitHeart:   '♥',
		card.SuitClub:    '♣',
		card.SuitDiamond: '♦',
	}
	colorStyles = map[card.Color]lipgloss.Style{
		card.ColorRed:   lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "1", Dark: "9"}),
		card.ColorBlack: lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "8", Dark: "15"}),
	}
)

func renderCard(c *card.Card) string {
	if c == nil {
		return "░░░"
	}
	if c.Face() == card.FaceDown {
		return "▓▓▓"
	}
	r := string(suitSymbols[c.Suit()])
	return colorStyles[c.Suit().Color()].Render(fmt.Sprintf("%2s%s", rankSymbols[c.Rank()], r))
}

func (t TUI) renderDeck(dt klondike.DeckID) string {
	selected := t.selectedDeckID() == dt
	focused := t.focusedDeckID() == dt
	switch {
	case dt == klondike.DeckS, dt == klondike.DeckW, klondike.IsFoundation(dt):
		d := t.game.Deck(dt)
		r := renderCard(nil)
		if len(d) > 0 {
			r = renderCard(&d[len(d)-1])
		}
		return styleCardNormal.Render(r)
	case klondike.IsTableau(dt):
		d := t.game.Deck(dt)
		cards := make([]string, 0, len(d))
		for i, ca := range d {
			s := styleCardNormal
			if t.state == stateCardSelection && focused && t.focusedCard == int8(i) {
				s = styleCardFocused
			} else if selected && t.focusedCard == int8(i) {
				s = styleCardSelected
			}
			cards = append(cards, s.Render(renderCard(&ca)))
		}
		if len(cards) == 0 {
			cards = append(cards, styleCardNormal.Render(renderCard(nil)))
		}
		return lipgloss.JoinVertical(lipgloss.Left, cards...)
	}
	return "     "
}
