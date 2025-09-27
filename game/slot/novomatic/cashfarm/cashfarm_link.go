//go:build !prod || full || novomatic

package cashfarm

import (
	_ "embed"

	"github.com/slotopol/server/game"
)

//go:embed cashfarm_reel.yaml
var reels []byte

var Info = game.AlgInfo{
	Aliases: []game.GameAlias{
		{Prov: "Novomatic", Name: "Cash Farm", Date: game.Date(2013, 1, 21)}, // see: https://casino.ru/cash-farm-novomatic/
	},
	AlgDescr: game.AlgDescr{
		GT: game.GTslot,
		GP: game.GPlpay |
			game.GPcasc |
			game.GPretrig |
			game.GPscat |
			game.GPwild,
		SX: 5,
		SY: 3,
		SN: len(LinePay),
		LN: len(BetLines),
		BN: 1,
	},
	Update: func(ai *game.AlgInfo) { ai.RTP = game.MakeRtpList(ReelsMap) },
}

func init() {
	Info.SetupFactory(func() game.Gamble { return NewGame() }, CalcStat)
	game.DataRouter["novomatic/cashfarm/reel"] = &ReelsMap
	game.LoadMap = append(game.LoadMap, reels)
}
