package groovypowers

import (
	"context"
	"fmt"
	"io"

	"github.com/slotopol/server/game/slot"
)

func CalcStatBon(ctx context.Context, mrtp float64) float64 {
	var reels, _ = ReelsMap.FindClosest(mrtp)
	var g = NewGame()
	g.Sel = 1
	g.BM = true // set bonus mode
	var s slot.Stat

	var calc = func(w io.Writer) float64 {
		var cost, _ = g.Cost()
		var lrtp, srtp = s.SymRTP(cost)
		var rtpsym = lrtp + srtp
		fmt.Fprintf(w, "RTP = %.5g(lined) + %.5g(scatter) = %.6f%%\n", lrtp, srtp, rtpsym)
		return rtpsym
	}

	return slot.ScanReels5x(ctx, &s, g, reels, calc)
}

func CalcStatReg(ctx context.Context, mrtp float64) float64 {
	fmt.Printf("*bonus mode calculations*\n")
	var rtpbm = CalcStatBon(ctx, mrtp)
	if ctx.Err() != nil {
		return 0
	}
	fmt.Printf("*regular mode calculations*\n")
	var reels, _ = ReelsMap.FindClosest(mrtp)
	var g = NewGame()
	g.Sel = 1
	var s slot.Stat

	var calc = func(w io.Writer) float64 {
		var cost, _ = g.Cost()
		var lrtp, srtp = s.SymRTP(cost)
		var rtpsym = lrtp + srtp
		var rtp = rtpsym*(1-Pbm) + rtpbm*Pbm
		fmt.Fprintf(w, "symbols: %.5g(lined) + %.5g(scatter) = %.6f%%\n", lrtp, srtp, rtpsym)
		fmt.Fprintf(w, "RTP = %.5g(reg)*%.5g + %.5g(bm)*%.5g = %.6f%%\n", rtpsym, 1-Pbm, rtpbm, Pbm, rtp)
		return rtp
	}

	return slot.ScanReels5x(ctx, &s, g, reels, calc)
}
