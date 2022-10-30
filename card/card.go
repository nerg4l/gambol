package card

const (
	SuitSpade Suit = 1 << iota
	SuitHeart
	SuitClub
	SuitDiamond
)

type Suit uint8

const (
	ColorRed   = Color(SuitHeart | SuitDiamond)
	ColorBlack = Color(SuitSpade | SuitClub)
)

var colors = map[Suit]Color{
	SuitSpade:   ColorBlack,
	SuitHeart:   ColorRed,
	SuitClub:    ColorBlack,
	SuitDiamond: ColorRed,
}

type Color uint8

func (s Suit) Color() Color {
	return colors[s]
}

const (
	IndexAce Rank = iota
	Index2
	Index3
	Index4
	Index5
	Index6
	Index7
	Index8
	Index9
	Index10
	IndexJack
	IndexQueen
	IndexKing

	IndexA = IndexAce
	IndexJ = IndexJack
	IndexQ = IndexQueen
	IndexK = IndexKing
)

type Rank uint8

const (
	FaceDown Face = iota
	FaceUp
)

type Face uint8

type Card struct {
	suit Suit
	rank Rank
	face Face
}

func (c Card) Suit() Suit {
	return c.suit
}

func (c Card) Rank() Rank {
	return c.rank
}

func (c Card) Face() Face {
	return c.face
}

func (c Card) Flip() Card {
	if c.face == FaceDown {
		c.face = FaceUp
	} else {
		c.face = FaceDown
	}
	return c
}

type Deck []Card

func NewDeck(f Face) Deck {
	d := make(Deck, 4*13)
	for i := uint8(0); i < 4; i++ {
		for j := uint8(0); j < 13; j++ {
			d[i*13+j] = Card{suit: Suit(1 << i), rank: Rank(j), face: f}
		}
	}
	return d
}
