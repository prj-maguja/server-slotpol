//go:build !prod || full || agt

package doubleice

import (
	_ "embed"

	"github.com/slotopol/server/game"
)

//go:embed doubleice_data.yaml
var data []byte

var Info = game.AlgInfo{
	Aliases: []game.GameAlias{
		{Prov: "AGT", Name: "Double Ice"},
		{Prov: "AGT", Name: "Double Hot"}, // see: https://agtsoftware.com/games/agt/double
	},
	AlgDescr: game.AlgDescr{
		GT: game.GTslot,
		GP: game.GPfill |
			game.GPfgno,
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
	game.DataRouter["agt/doubleice/reel"] = &ReelsMap
	game.LoadMap = append(game.LoadMap, data)
}
