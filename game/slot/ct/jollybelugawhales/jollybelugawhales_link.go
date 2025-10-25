//go:build !prod || full || ct

package jollybelugawhales

import (
	_ "embed"

	"github.com/slotopol/server/game"
)

//go:embed jollybelugawhales_data.yaml
var data []byte

var Info = game.AlgInfo{
	Aliases: []game.GameAlias{
		{Prov: "CT Interactive", Name: "Jolly Beluga Whales", Date: game.Date(2020, 11, 26)}, // see: https://www.slotsmate.com/software/ct-interactive/jolly-beluga-whales
	},
	AlgDescr: game.AlgDescr{
		GT: game.GTslot,
		GP: game.GPlpay |
			game.GPretrig |
			game.GPscat |
			game.GPrwild,
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
	game.DataRouter["ctinteractive/jollybelugawhales/reel"] = &ReelsMap
	game.LoadMap = append(game.LoadMap, data)
}
