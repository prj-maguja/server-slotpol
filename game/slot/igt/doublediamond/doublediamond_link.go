//go:build !prod || full || igt

package doublediamond

import (
	_ "embed"

	"github.com/slotopol/server/game"
)

//go:embed doublediamond_data.yaml
var data []byte

var Info = game.AlgInfo{
	Aliases: []game.GameAlias{
		{Prov: "IGT", Name: "Double Diamond", Date: game.Year(2005)}, // see: https://www.slotsmate.com/software/igt/double-diamond
	},
	AlgDescr: game.AlgDescr{
		GT: game.GTslot,
		GP: game.GPfgno |
			game.GPwild |
			game.GPwmult,
		SX: 3,
		SY: 3,
		SN: len(LinePay),
		LN: len(BetLines),
		BN: 0,
	},
	Update: func(ai *game.AlgInfo) { ai.RTP = game.MakeRtpList(ReelsMap) },
}

func init() {
	Info.SetupFactory(func() game.Gamble { return NewGame() }, CalcStat)
	game.DataRouter["igt/doublediamond/reel"] = &ReelsMap
	game.LoadMap = append(game.LoadMap, data)
}
