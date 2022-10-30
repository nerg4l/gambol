package klondike

import (
	"errors"
	"github.com/nerg4l/gambol/card"
	"math/rand"
	"time"
)

type Game struct {
	stock      card.Deck
	waste      card.Deck
	foundation [4]card.Deck
	tableau    [7]card.Deck
}

func NewGame() Game {
	d := card.NewDeck(card.FaceDown)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
	g := Game{
		waste:      make(card.Deck, 0, 24),
		foundation: [4]card.Deck{},
		tableau:    [7]card.Deck{},
	}
	for i := range g.tableau {
		dd := append(card.Deck{}, d[:i+1]...)
		dd[len(dd)-1] = dd[len(dd)-1].Flip()
		g.tableau[i] = dd
		d = d[i+1:]
	}
	g.stock = d[1:]
	g.waste = card.Deck{d[0].Flip()}
	return g
}

const (
	DeckStock DeckID = iota
	DeckWaste
	DeckFoundation1
	DeckFoundation2
	DeckFoundation3
	DeckFoundation4
	DeckTableau1
	DeckTableau2
	DeckTableau3
	DeckTableau4
	DeckTableau5
	DeckTableau6
	DeckTableau7

	DeckS  = DeckStock       // DeckS is an alias for DeckStock
	DeckW  = DeckWaste       // DeckW is an alias for DeckWaste
	DeckF1 = DeckFoundation1 // DeckF1 is an alias for DeckFoundation1
	DeckF2 = DeckFoundation2 // DeckF2 is an alias for DeckFoundation2
	DeckF3 = DeckFoundation3 // DeckF3 is an alias for DeckFoundation3
	DeckF4 = DeckFoundation4 // DeckF4 is an alias for DeckFoundation4
	DeckT1 = DeckTableau1    // DeckT1 is an alias for DeckTableau1
	DeckT2 = DeckTableau2    // DeckT2 is an alias for DeckTableau2
	DeckT3 = DeckTableau3    // DeckT3 is an alias for DeckTableau3
	DeckT4 = DeckTableau4    // DeckT4 is an alias for DeckTableau4
	DeckT5 = DeckTableau5    // DeckT5 is an alias for DeckTableau5
	DeckT6 = DeckTableau6    // DeckT6 is an alias for DeckTableau6
	DeckT7 = DeckTableau7    // DeckT7 is an alias for DeckTableau7
)

type DeckID int8

var indexesFoundation = map[DeckID]int8{
	DeckF1: 0,
	DeckF2: 1,
	DeckF3: 2,
	DeckF4: 3,
}

func IsFoundation(d DeckID) bool {
	_, ok := indexesFoundation[d]
	return ok
}

var indexesTableau = map[DeckID]int8{
	DeckT1: 0,
	DeckT2: 1,
	DeckT3: 2,
	DeckT4: 3,
	DeckT5: 4,
	DeckT6: 5,
	DeckT7: 6,
}

func IsTableau(d DeckID) bool {
	_, ok := indexesTableau[d]
	return ok
}

func (g Game) Deck(dt DeckID) card.Deck {
	switch {
	case dt == DeckS:
		return g.stock
	case dt == DeckW:
		return g.waste
	case IsFoundation(dt):
		return g.foundation[indexesFoundation[dt]]
	case IsTableau(dt):
		return g.tableau[indexesTableau[dt]]
	}
	return nil
}

var ErrInvalidMove = errors.New("invalid move")

func (g Game) Move(from, to DeckID, c *card.Card) (Game, error) {
	switch {
	case from == DeckS:
		if c != nil {
			return g, ErrInvalidMove
		}
		return moveStock(g, to)
	case from == DeckW:
		if c != nil {
			return g, ErrInvalidMove
		}
		return moveWaste(g, to)
	case IsFoundation(from):
		if c != nil {
			return g, ErrInvalidMove
		}
		return moveFoundation(g, from, to)
	case IsTableau(from):
		if c == nil { // card is required for deck split
			return g, ErrInvalidMove
		}
		return moveTableau(g, from, to, *c)
	}
	// make sure a card is always available in the waste when possible
	if len(g.waste) == 0 && len(g.stock) != 0 {
		l := len(g.stock)
		g.stock, g.waste = g.stock[:l-1], append(g.waste, g.stock[l-1].Flip())
	}
	return g, nil
}

func moveStock(g Game, to DeckID) (Game, error) {
	switch {
	case to == DeckW:
		l := len(g.stock)
		if l == 0 {
			if len(g.waste) == 0 {
				return g, ErrInvalidMove
			}
			l = len(g.waste)
			g.stock = make(card.Deck, l)
			for i := 0; i < l; i++ {
				g.stock[l-1-i] = g.waste[i].Flip()
			}
			g.waste = nil
		} else {
			g.stock, g.waste = g.stock[:l-1], append(g.waste, g.stock[l-1].Flip())
		}
	default:
		return g, ErrInvalidMove
	}
	return g, nil
}

func moveWaste(g Game, to DeckID) (Game, error) {
	l := len(g.waste)
	if l == 0 {
		return g, ErrInvalidMove
	}
	c := g.waste[l-1]
	switch {
	case IsFoundation(to):
		ti := indexesFoundation[to]
		if isValidMoveFoundation(g.foundation[ti], c) {
			return g, ErrInvalidMove
		}
		g.waste, g.foundation[ti] = g.waste[:l-1], append(g.foundation[ti], c)
	case IsTableau(to):
		ti := indexesTableau[to]
		if isValidMoveTableau(g.tableau[ti], c) {
			return g, ErrInvalidMove
		}
		g.waste, g.tableau[ti] = g.waste[:l-1], append(g.tableau[ti], c)
	default:
		return g, ErrInvalidMove
	}
	return g, nil
}

func moveFoundation(g Game, from DeckID, to DeckID) (Game, error) {
	fi := indexesFoundation[from]
	l := len(g.foundation[fi])
	if l == 0 {
		return g, ErrInvalidMove
	}
	switch {
	case IsFoundation(to):
		ti := indexesFoundation[to]
		if len(g.foundation[fi]) != 1 && len(g.foundation[ti]) != 0 {
			return g, ErrInvalidMove
		}
		g.foundation[ti], g.foundation[fi] = g.foundation[fi], g.foundation[ti]
	case IsTableau(to):
		ti := indexesTableau[to]
		c := g.foundation[fi][l-1]
		if isValidMoveTableau(g.tableau[ti], c) {
			return g, ErrInvalidMove
		}
		g.foundation[fi], g.tableau[ti] = g.foundation[fi][:l-1], append(g.tableau[ti], c)
	default:
		return g, ErrInvalidMove
	}
	return g, nil
}

func moveTableau(g Game, from DeckID, to DeckID, c card.Card) (Game, error) {
	fi := indexesTableau[from]
	l := len(g.tableau[fi])
	if l == 0 {
		return g, ErrInvalidMove
	}
	ci := -1
	for i, cc := range g.tableau[fi] {
		if cc.Face() == card.FaceUp && c == cc {
			ci = i
			break
		}
	}
	if ci == -1 {
		return g, ErrInvalidMove
	}
	cc := g.tableau[fi][ci]
	switch {
	case IsFoundation(to):
		ti := indexesFoundation[to]
		if ci != l-1 {
			return g, ErrInvalidMove
		}
		if isValidMoveFoundation(g.foundation[ti], cc) {
			return g, ErrInvalidMove
		}
		g.tableau[fi], g.foundation[ti] = g.tableau[fi][:ci], append(g.foundation[ti], g.tableau[fi][ci:]...)
	case IsTableau(to):
		ti := indexesTableau[to]
		if isValidMoveTableau(g.tableau[ti], cc) {
			return g, ErrInvalidMove
		}
		g.tableau[fi], g.tableau[ti] = g.tableau[fi][:ci], append(g.tableau[ti], g.tableau[fi][ci:]...)
	default:
		return g, ErrInvalidMove
	}
	l = len(g.tableau[fi])
	if l != 0 {
		c := g.tableau[fi][l-1]
		if c.Face() == card.FaceDown {
			g.tableau[fi][l-1] = c.Flip()
		}
	}
	return g, nil
}

func isValidMoveFoundation(d card.Deck, c card.Card) bool {
	l := len(d)
	if l == 0 {
		if c.Rank() != card.IndexA {
			return true
		}
	} else {
		hi, lo := c, d[l-1]
		if hi.Rank()-1 != lo.Rank() || hi.Suit() != lo.Suit() {
			return true
		}
	}
	return false
}

func isValidMoveTableau(d card.Deck, c card.Card) bool {
	l := len(d)
	if l == 0 {
		if c.Rank() != card.IndexK {
			return true
		}
	} else {
		lo, hi := c, d[l-1]
		if hi.Rank()-1 != lo.Rank() || hi.Suit().Color() == lo.Suit().Color() {
			return true
		}
	}
	return false
}
