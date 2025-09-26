//go:build !prod || full || playtech

package greatblue

import (
	"github.com/slotopol/server/game"
)

var Info = game.AlgInfo{
	Aliases: []game.GameAlias{
		{Prov: "Playtech", Name: "Great Blue", Date: game.Date(2013, 1, 1)}, // see: https://www.slotsmate.com/software/playtech/great-blue
		{Prov: "Playtech", Name: "Irish Luck", Date: game.Date(2009, 1, 1)}, // see: https://www.slotsmate.com/software/playtech/irish-luck
	},
	AlgDescr: game.AlgDescr{
		GT: game.GTslot,
		GP: game.GPlpay |
			game.GPlsel |
			game.GPretrig |
			game.GPfgmult |
			game.GPscat |
			game.GPwild |
			game.GPwmult,
		SX:  5,
		SY:  3,
		SN:  len(LinePay),
		LN:  len(BetLines),
		BN:  0,
		RTP: game.MakeRtpList(ReelsMap),
	},
}

func init() {
	Info.SetupFactory(func() game.Gamble { return NewGame() }, CalcStat)
	game.DataRouter["playtech/greatblue/reel"] = &ReelsMap
}
