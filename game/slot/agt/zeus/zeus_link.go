//go:build !prod || full || agt

package zeus

import (
	_ "embed"

	"github.com/slotopol/server/game"
)

//go:embed zeus_data.yaml
var data []byte

var Info = game.AlgInfo{
	Aliases: []game.GameAlias{
		{Prov: "AGT", Name: "Zeus", Date: game.Year(2025)}, // see: https://agtsoftware.org/games/agt/zeus
	},
	AlgDescr: game.AlgDescr{
		GT: game.GTslot,
		GP: game.GPlpay |
			game.GPlsel |
			game.GPretrig |
			game.GPscat |
			game.GPwild,
		SX: 4,
		SY: 4,
		SN: len(LinePay),
		LN: len(BetLines),
		BN: 0,
	},
	Update: func(ai *game.AlgInfo) { ai.RTP = game.MakeRtpList(ReelsMap) },
}

func init() {
	Info.SetupFactory(func() game.Gamble { return NewGame() }, CalcStat)
	game.DataRouter["agt/zeus/reel"] = &ReelsMap
	game.LoadMap = append(game.LoadMap, data)
}
