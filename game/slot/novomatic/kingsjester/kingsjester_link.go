//go:build !prod || full || novomatic

package kingsjester

import (
	_ "embed"

	"github.com/slotopol/server/game"
)

//go:embed kingsjester_data.yaml
var data []byte

var Info = game.AlgInfo{
	Aliases: []game.GameAlias{
		{Prov: "Novomatic", Name: "King's Jester", Date: game.Date(2015, 9, 1)}, // see: https://www.slotsmate.com/software/novomatic/kings-jester
	},
	AlgDescr: game.AlgDescr{
		GT: game.GTslot,
		GP: game.GPlpay |
			game.GPlsel |
			game.GPretrig |
			game.GPjack |
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
	Info.SetupFactory(func() game.Gamble { return NewGame() }, CalcStat)
	game.DataRouter["novomatic/kingsjester/reel"] = &ReelsMap
	game.DataRouter["novomatic/kingsjester/jack"] = &JackMap
	game.LoadMap = append(game.LoadMap, data)
}
