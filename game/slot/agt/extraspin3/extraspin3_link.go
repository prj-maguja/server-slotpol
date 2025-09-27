//go:build !prod || full || agt

package extraspin3

import (
	"github.com/slotopol/server/game"
	"github.com/slotopol/server/game/slot/agt/extraspin"
)

var Info = game.AlgInfo{
	Aliases: []game.GameAlias{
		{Prov: "AGT", Name: "Extra Spin III"},
	},
	AlgDescr: game.AlgDescr{
		GT: game.GTslot,
		GP: game.GPlpay |
			game.GPlsel |
			game.GPretrig |
			game.GPscat |
			game.GPwild,
		SX: 5,
		SY: 3,
		SN: len(LinePay),
		LN: len(BetLines),
		BN: 0,
	},
	Update: func(ai *game.AlgInfo) { ai.RTP = game.MakeRtpList(extraspin.ReelsMap) },
}

func init() {
	Info.SetupFactory(func() game.Gamble { return NewGame() }, extraspin.CalcStat)
}
