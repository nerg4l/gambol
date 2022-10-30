package card

import (
	"reflect"
	"testing"
)

func TestColor(t *testing.T) {
	tests := []struct {
		name  string
		suit  Suit
		color Color
		want  bool
	}{
		{
			name:  "spade black",
			suit:  SuitSpade,
			color: ColorBlack,
			want:  true,
		},
		{
			name:  "spade red",
			suit:  SuitSpade,
			color: ColorRed,
			want:  false,
		},
		{
			name:  "heart black",
			suit:  SuitHeart,
			color: ColorBlack,
			want:  false,
		},
		{
			name:  "heart red",
			suit:  SuitHeart,
			color: ColorRed,
			want:  true,
		},
		{
			name:  "club black",
			suit:  SuitClub,
			color: ColorBlack,
			want:  true,
		},
		{
			name:  "club red",
			suit:  SuitClub,
			color: ColorRed,
			want:  false,
		},
		{
			name:  "diamond black",
			suit:  SuitDiamond,
			color: ColorBlack,
			want:  false,
		},
		{
			name:  "diamond red",
			suit:  SuitDiamond,
			color: ColorRed,
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want != (Color(tt.suit)&tt.color > 0) {
				t.Errorf("Suit %b & Color %b, want %v", tt.suit, tt.color, tt.want)
			}
		})
	}
}

var (
	faceUpDeck = Deck{
		Card{suit: SuitSpade, rank: 0, face: FaceUp},
		Card{suit: SuitSpade, rank: 1, face: FaceUp},
		Card{suit: SuitSpade, rank: 2, face: FaceUp},
		Card{suit: SuitSpade, rank: 3, face: FaceUp},
		Card{suit: SuitSpade, rank: 4, face: FaceUp},
		Card{suit: SuitSpade, rank: 5, face: FaceUp},
		Card{suit: SuitSpade, rank: 6, face: FaceUp},
		Card{suit: SuitSpade, rank: 7, face: FaceUp},
		Card{suit: SuitSpade, rank: 8, face: FaceUp},
		Card{suit: SuitSpade, rank: 9, face: FaceUp},
		Card{suit: SuitSpade, rank: 10, face: FaceUp},
		Card{suit: SuitSpade, rank: 11, face: FaceUp},
		Card{suit: SuitSpade, rank: 12, face: FaceUp},
		Card{suit: SuitHeart, rank: 0, face: FaceUp},
		Card{suit: SuitHeart, rank: 1, face: FaceUp},
		Card{suit: SuitHeart, rank: 2, face: FaceUp},
		Card{suit: SuitHeart, rank: 3, face: FaceUp},
		Card{suit: SuitHeart, rank: 4, face: FaceUp},
		Card{suit: SuitHeart, rank: 5, face: FaceUp},
		Card{suit: SuitHeart, rank: 6, face: FaceUp},
		Card{suit: SuitHeart, rank: 7, face: FaceUp},
		Card{suit: SuitHeart, rank: 8, face: FaceUp},
		Card{suit: SuitHeart, rank: 9, face: FaceUp},
		Card{suit: SuitHeart, rank: 10, face: FaceUp},
		Card{suit: SuitHeart, rank: 11, face: FaceUp},
		Card{suit: SuitHeart, rank: 12, face: FaceUp},
		Card{suit: SuitClub, rank: 0, face: FaceUp},
		Card{suit: SuitClub, rank: 1, face: FaceUp},
		Card{suit: SuitClub, rank: 2, face: FaceUp},
		Card{suit: SuitClub, rank: 3, face: FaceUp},
		Card{suit: SuitClub, rank: 4, face: FaceUp},
		Card{suit: SuitClub, rank: 5, face: FaceUp},
		Card{suit: SuitClub, rank: 6, face: FaceUp},
		Card{suit: SuitClub, rank: 7, face: FaceUp},
		Card{suit: SuitClub, rank: 8, face: FaceUp},
		Card{suit: SuitClub, rank: 9, face: FaceUp},
		Card{suit: SuitClub, rank: 10, face: FaceUp},
		Card{suit: SuitClub, rank: 11, face: FaceUp},
		Card{suit: SuitClub, rank: 12, face: FaceUp},
		Card{suit: SuitDiamond, rank: 0, face: FaceUp},
		Card{suit: SuitDiamond, rank: 1, face: FaceUp},
		Card{suit: SuitDiamond, rank: 2, face: FaceUp},
		Card{suit: SuitDiamond, rank: 3, face: FaceUp},
		Card{suit: SuitDiamond, rank: 4, face: FaceUp},
		Card{suit: SuitDiamond, rank: 5, face: FaceUp},
		Card{suit: SuitDiamond, rank: 6, face: FaceUp},
		Card{suit: SuitDiamond, rank: 7, face: FaceUp},
		Card{suit: SuitDiamond, rank: 8, face: FaceUp},
		Card{suit: SuitDiamond, rank: 9, face: FaceUp},
		Card{suit: SuitDiamond, rank: 10, face: FaceUp},
		Card{suit: SuitDiamond, rank: 11, face: FaceUp},
		Card{suit: SuitDiamond, rank: 12, face: FaceUp},
	}
	faceDownDeck = Deck{
		Card{suit: SuitSpade, rank: 0, face: FaceDown},
		Card{suit: SuitSpade, rank: 1, face: FaceDown},
		Card{suit: SuitSpade, rank: 2, face: FaceDown},
		Card{suit: SuitSpade, rank: 3, face: FaceDown},
		Card{suit: SuitSpade, rank: 4, face: FaceDown},
		Card{suit: SuitSpade, rank: 5, face: FaceDown},
		Card{suit: SuitSpade, rank: 6, face: FaceDown},
		Card{suit: SuitSpade, rank: 7, face: FaceDown},
		Card{suit: SuitSpade, rank: 8, face: FaceDown},
		Card{suit: SuitSpade, rank: 9, face: FaceDown},
		Card{suit: SuitSpade, rank: 10, face: FaceDown},
		Card{suit: SuitSpade, rank: 11, face: FaceDown},
		Card{suit: SuitSpade, rank: 12, face: FaceDown},
		Card{suit: SuitHeart, rank: 0, face: FaceDown},
		Card{suit: SuitHeart, rank: 1, face: FaceDown},
		Card{suit: SuitHeart, rank: 2, face: FaceDown},
		Card{suit: SuitHeart, rank: 3, face: FaceDown},
		Card{suit: SuitHeart, rank: 4, face: FaceDown},
		Card{suit: SuitHeart, rank: 5, face: FaceDown},
		Card{suit: SuitHeart, rank: 6, face: FaceDown},
		Card{suit: SuitHeart, rank: 7, face: FaceDown},
		Card{suit: SuitHeart, rank: 8, face: FaceDown},
		Card{suit: SuitHeart, rank: 9, face: FaceDown},
		Card{suit: SuitHeart, rank: 10, face: FaceDown},
		Card{suit: SuitHeart, rank: 11, face: FaceDown},
		Card{suit: SuitHeart, rank: 12, face: FaceDown},
		Card{suit: SuitClub, rank: 0, face: FaceDown},
		Card{suit: SuitClub, rank: 1, face: FaceDown},
		Card{suit: SuitClub, rank: 2, face: FaceDown},
		Card{suit: SuitClub, rank: 3, face: FaceDown},
		Card{suit: SuitClub, rank: 4, face: FaceDown},
		Card{suit: SuitClub, rank: 5, face: FaceDown},
		Card{suit: SuitClub, rank: 6, face: FaceDown},
		Card{suit: SuitClub, rank: 7, face: FaceDown},
		Card{suit: SuitClub, rank: 8, face: FaceDown},
		Card{suit: SuitClub, rank: 9, face: FaceDown},
		Card{suit: SuitClub, rank: 10, face: FaceDown},
		Card{suit: SuitClub, rank: 11, face: FaceDown},
		Card{suit: SuitClub, rank: 12, face: FaceDown},
		Card{suit: SuitDiamond, rank: 0, face: FaceDown},
		Card{suit: SuitDiamond, rank: 1, face: FaceDown},
		Card{suit: SuitDiamond, rank: 2, face: FaceDown},
		Card{suit: SuitDiamond, rank: 3, face: FaceDown},
		Card{suit: SuitDiamond, rank: 4, face: FaceDown},
		Card{suit: SuitDiamond, rank: 5, face: FaceDown},
		Card{suit: SuitDiamond, rank: 6, face: FaceDown},
		Card{suit: SuitDiamond, rank: 7, face: FaceDown},
		Card{suit: SuitDiamond, rank: 8, face: FaceDown},
		Card{suit: SuitDiamond, rank: 9, face: FaceDown},
		Card{suit: SuitDiamond, rank: 10, face: FaceDown},
		Card{suit: SuitDiamond, rank: 11, face: FaceDown},
		Card{suit: SuitDiamond, rank: 12, face: FaceDown},
	}
)

func TestNewDeck(t *testing.T) {
	type args struct {
		f Face
	}
	tests := []struct {
		name string
		args args
		want Deck
	}{
		{
			name: "face up",
			args: args{f: FaceUp},
			want: faceUpDeck,
		},
		{
			name: "face down",
			args: args{f: FaceDown},
			want: faceDownDeck,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDeck(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDeck() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
