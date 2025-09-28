//go:build !prod || full || agt

package icefruits

import (
	_ "embed"

	"github.com/slotopol/server/game"
)

//go:embed icefruits_data.yaml
var data []byte

var Info = game.AlgInfo{
	Aliases: []game.GameAlias{
		{Prov: "AGT", Name: "Ice Fruits"},
		{Prov: "AGT", Name: "Mega Shine"}, // see: https://agtsoftware.com/games/agt/megashine
	},
	AlgDescr: game.AlgDescr{
		GT: game.GTslot,
		GP: game.GPlpay |
			game.GPlsel |
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
	game.DataRouter["agt/icefruits/reel"] = &ReelsMap
	game.LoadMap = append(game.LoadMap, data)
}
