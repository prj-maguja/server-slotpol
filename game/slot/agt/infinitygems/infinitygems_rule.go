package infinitygems

// See: https://agtsoftware.com/games/agt/infinitygems

import (
	"github.com/slotopol/server/game/slot"
)

var ReelsBon *slot.Reels5x

var ReelsMap slot.ReelsMap[*slot.Reels5x]

// Lined payment.
var LinePay = [13][5]float64{
	{0, 10, 250, 2500, 9000}, //  1 wild
	{0, 2, 30, 120, 800},     //  2 emerald
	{0, 2, 30, 120, 800},     //  3 heliodor
	{0, 0, 20, 100, 400},     //  4 ruby
	{0, 0, 20, 70, 250},      //  5 amethyst
	{0, 0, 20, 70, 250},      //  6 sapphire
	{0, 0, 10, 50, 120},      //  7 ace
	{0, 0, 10, 50, 120},      //  8 king
	{0, 0, 4, 30, 100},       //  9 queen
	{0, 0, 4, 30, 100},       // 10 jack
	{0, 0, 4, 30, 100},       // 11 ten
	{0, 2, 4, 30, 100},       // 12 nine
	{},                       // 13 scatter
}

// Scatters payment.
var ScatPay = [5]float64{0, 2, 4, 20, 500} // 13 scatter

// Scatter freespins table
var ScatFreespin = [5]int{0, 0, 50, 50, 50} // 13 scatter

// Bet lines
var BetLines = slot.BetLinesAgt5x3[:20]

type Game struct {
	slot.Screen5x3 `yaml:",inline"`
	slot.Slotx     `yaml:",inline"`
}

// Declare conformity with SlotGame interface.
var _ slot.SlotGame = (*Game)(nil)

func NewGame() *Game {
	return &Game{
		Slotx: slot.Slotx{
			Sel: len(BetLines),
			Bet: 1,
		},
	}
}

func (g *Game) Clone() slot.SlotGame {
	var clone = *g
	return &clone
}

const wild, scat = 1, 13

func (g *Game) Scanner(wins *slot.Wins) error {
	g.ScanLined(wins)
	g.ScanScatters(wins)
	return nil
}

// Lined symbols calculation.
func (g *Game) ScanLined(wins *slot.Wins) {
	for li, line := range BetLines[:g.Sel] {
		var mw float64 = 1 // mult wild
		var numw, numl slot.Pos = 0, 5
		var syml slot.Sym
		var x slot.Pos
		for x = 1; x <= 5; x++ {
			var sx = g.LY(x, line)
			if sx == wild {
				if syml == 0 {
					numw = x
				}
				mw = 2
			} else if syml == 0 {
				syml = sx
			} else if sx != syml {
				numl = x - 1
				break
			}
		}

		var payw, payl float64
		if numw >= 2 {
			payw = LinePay[wild-1][numw-1]
		}
		if numl >= 2 && syml > 0 {
			payl = LinePay[syml-1][numl-1]
		}
		if payl*mw > payw {
			*wins = append(*wins, slot.WinItem{
				Pay: g.Bet * payl,
				MP:  mw,
				Sym: syml,
				Num: numl,
				LI:  li + 1,
				XY:  line.CopyL(numl),
			})
		} else if payw > 0 {
			*wins = append(*wins, slot.WinItem{
				Pay: g.Bet * payw,
				MP:  1,
				Sym: wild,
				Num: numw,
				LI:  li + 1,
				XY:  line.CopyL(numw),
			})
		}
	}
}

// Scatters calculation.
func (g *Game) ScanScatters(wins *slot.Wins) {
	if count := g.ScatNum(scat); count >= 2 {
		var pay, fs = ScatPay[count-1], ScatFreespin[count-1]
		*wins = append(*wins, slot.WinItem{
			Pay: g.Bet * float64(g.Sel) * pay,
			MP:  1,
			Sym: scat,
			Num: count,
			XY:  g.ScatPos(scat),
			FS:  fs,
		})
	}
}

func (g *Game) Spin(mrtp float64) {
	if g.FSR == 0 {
		var reels, _ = ReelsMap.FindClosest(mrtp)
		g.ReelSpin(reels)
	} else {
		g.ReelSpin(ReelsBon)
	}
}

func (g *Game) SetSel(sel int) error {
	return g.SetSelNum(sel, len(BetLines))
}
