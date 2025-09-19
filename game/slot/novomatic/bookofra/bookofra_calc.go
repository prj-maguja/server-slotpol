package bookofra

import (
	"context"
	"fmt"
	"io"

	"github.com/slotopol/server/game/slot"
)

func CalcStatBon(ctx context.Context, es slot.Sym) float64 {
	var reels = ReelsBon
	var g = NewGame()
	g.Sel = 1
	g.FSR = 10 // set free spins mode
	g.ES = es
	var s slot.Stat

	var calc = func(w io.Writer) float64 {
		var cost, _ = g.Cost()
		var lrtp, srtp = s.SymRTP(cost)
		var rtpsym = lrtp + srtp
		var q, sq = s.FSQ()
		var rtp = sq * rtpsym
		fmt.Fprintf(w, "symbols: %.5g(lined) + %.5g(scatter) = %.6f%%\n", lrtp, srtp, rtpsym)
		fmt.Fprintf(w, "free spins %d, q = %.5g, sq = 1/(1-q) = %.6f\n", s.FreeCountU(), q, sq)
		fmt.Fprintf(w, "free games frequency: 1/%.5g\n", s.FGF())
		fmt.Fprintf(w, "RTP[%d] = sq*rtp(sym) = %.5g*%.5g = %.6f%%\n", es, sq, rtpsym, rtp)
		return rtp
	}

	return slot.ScanReels5x(ctx, &s, g, reels, calc)
}

func CalcStatReg(ctx context.Context, mrtp float64) float64 {
	fmt.Printf("*bonus reels calculations*\n")
	var rtpe [9]float64
	var es slot.Sym
	for es = 1; es < book; es++ {
		fmt.Printf("*calculations for expanding symbol [%d]*\n", es)
		rtpe[es-1] = CalcStatBon(ctx, es)
		if ctx.Err() != nil {
			return 0
		}
	}
	var rtpfs float64
	for _, rtp := range rtpe {
		rtpfs += rtp
	}
	rtpfs /= 9
	fmt.Printf("RTPfs = %.6f%%\n", rtpfs)

	fmt.Printf("*regular reels calculations*\n")
	var reels, _ = slot.FindClosest(ReelsMap, mrtp)
	var g = NewGame()
	g.Sel = 1
	var s slot.Stat

	var calc = func(w io.Writer) float64 {
		var cost, _ = g.Cost()
		var lrtp, srtp = s.SymRTP(cost)
		var rtpsym = lrtp + srtp
		var q, sq = s.FSQ()
		var rtp = rtpsym + q*rtpfs
		fmt.Fprintf(w, "symbols: %.5g(lined) + %.5g(scatter) = %.6f%%\n", lrtp, srtp, rtpsym)
		fmt.Fprintf(w, "free spins %d, q = %.5g, sq = 1/(1-q) = %.6f\n", s.FreeCountU(), q, sq)
		fmt.Fprintf(w, "free games frequency: 1/%.5g\n", s.FGF())
		fmt.Fprintf(w, "RTP = %.5g(sym) + %.5g*%.5g(fg) = %.6f%%\n", rtpsym, q, rtpfs, rtp)
		return rtp
	}

	return slot.ScanReels5x(ctx, &s, g, reels, calc)
}
