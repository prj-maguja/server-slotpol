package valkyrie

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/slotopol/server/game/slot"
)

func CalcStatBon(ctx context.Context, mrtp float64) float64 {
	var reels, _ = ReelsMap.FindClosest(mrtp)
	var g = NewGame()
	g.Sel = 1
	g.FSR = 15 // set free spins mode
	var s slot.Stat

	var calc = func(w io.Writer) float64 {
		var cost, _ = g.Cost()
		var lrtp, srtp = s.SymRTP(cost)
		if srtp > 0 {
			panic("scatters are presents on bonus games")
		}
		var rtp = lrtp + srtp
		fmt.Fprintf(w, "RTP = %.6f%%\n", rtp)
		return rtp
	}

	func() {
		var ctx2, cancel2 = context.WithCancel(ctx)
		defer cancel2()
		slot.BruteForce5x3Big(ctx2, &s, g, reels.Reel(1), ReelsBon, reels.Reel(5))
	}()
	return calc(os.Stdout)
}

func CalcStatReg(ctx context.Context, mrtp float64) float64 {
	fmt.Printf("*bonus reels calculations*\n")
	var rtpfs = CalcStatBon(ctx, mrtp)
	if ctx.Err() != nil {
		return 0
	}
	fmt.Printf("*regular reels calculations*\n")
	var reels, _ = ReelsMap.FindClosest(mrtp)
	var g = NewGame()
	g.Sel = 1
	var s slot.Stat

	var calc = func(w io.Writer) float64 {
		var cost, _ = g.Cost()
		var lrtp, srtp = s.SymRTP(cost)
		var rtpsym = lrtp + srtp
		var q, _ = s.FSQ()
		var rtp = rtpsym + q*rtpfs
		fmt.Fprintf(w, "symbols: %.5g(lined) + %.5g(scatter) = %.6f%%\n", lrtp, srtp, rtpsym)
		fmt.Fprintf(w, "free spins %d, q = %.6f\n", s.FreeCountU(), q)
		fmt.Fprintf(w, "free games frequency: 1/%.5g\n", s.FGF())
		fmt.Fprintf(w, "RTP = %.5g(sym) + %.5g*%.5g(fg) = %.6f%%\n", rtpsym, q, rtpfs, rtp)
		return rtp
	}

	return slot.ScanReels5x(ctx, &s, g, reels, calc)
}
