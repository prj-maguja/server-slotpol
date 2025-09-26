//go:build !prod || full || ct

package groovyautomat

import (
	"github.com/slotopol/server/game"
)

var Info = game.AlgInfo{
	Aliases: []game.GameAlias{
		{Prov: "CT Interactive", Name: "Groovy Automat", Date: game.Date(2020, 11, 26)}, // see: https://www.slotsmate.com/software/ct-interactive/groovy-automat
		{Prov: "CT Interactive", Name: "Golden Amulet", Date: game.Date(2020, 11, 26)},  // see: https://www.slotsmate.com/software/ct-interactive/golden-amulet
	},
	AlgDescr: game.AlgDescr{
		GT: game.GTslot,
		GP: game.GPfgno |
			game.GPscat,
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
	game.LoadMap["ctinteractive/groovyautomat/reel"] = &ReelsMap
}
