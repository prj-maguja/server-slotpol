//go:build !prod || full || ct

package halloweenhot

import (
	_ "embed"

	"github.com/slotopol/server/game"
)

//go:embed halloweenhot_data.yaml
var data []byte

var Info = game.AlgInfo{
	Aliases: []game.GameAlias{
		{Prov: "CT Interactive", Name: "Halloween Hot", Date: game.Date(2021, 10, 15)}, // see: https://www.slotsmate.com/software/ct-interactive/halloween-hot
	},
	AlgDescr: game.AlgDescr{
		GT: game.GTslot,
		GP: game.GPlpay |
			game.GPfill |
			game.GPfgno |
			game.GPscat |
			game.GPwild,
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
	game.DataRouter["ctinteractive/halloweenhot/reel"] = &ReelsMap
	game.LoadMap = append(game.LoadMap, data)
}
