//go:build !prod || full || novomatic

package helena

import (
	_ "embed"

	"github.com/slotopol/server/game"
)

//go:embed helena_bon.yaml
var rbon []byte

//go:embed helena_reel.yaml
var reels []byte

var Info = game.AlgInfo{
	Aliases: []game.GameAlias{
		{Prov: "Novomatic", Name: "Helena", Date: game.Date(2016, 8, 15)}, // see: https://www.slotsmate.com/software/novomatic/helena
	},
	AlgDescr: game.AlgDescr{
		GT: game.GTslot,
		GP: game.GPlpay |
			game.GPlsel |
			game.GPretrig |
			game.GPfgreel |
			game.GPfgmult |
			game.GPscat |
			game.GPwild |
			game.GPwmult,
		SX: 5,
		SY: 3,
		SN: len(LinePay),
		LN: len(BetLines),
		BN: 0,
	},
	Update: func(ai *game.AlgInfo) { ai.RTP = game.MakeRtpList(ReelsMap) },
}

func init() {
	Info.SetupFactory(func() game.Gamble { return NewGame() }, CalcStatReg)
	game.DataRouter["novomatic/helena/bon"] = &ReelsBon
	game.DataRouter["novomatic/helena/reel"] = &ReelsMap
	game.LoadMap = append(game.LoadMap, rbon, reels)
}
