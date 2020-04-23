package popcount

// pc[i] は i のポピュレーションカウントです
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount は x のポピュレーションカウント（1 が設定されているビット数）を返します
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCount23 は x のポピュレーションカウント（1 が設定されているビット数）を返します（練習問題2.3バージョン）
func PopCount23(x uint64) int {
	var b byte
	for i := 0; i < 8; i++ {
		b += pc[byte(x>>(i*8))]
	}
	return int(b)
}

// PopCount24 は x のポピュレーションカウント（1 が設定されているビット数）を返します（練習問題2.4バージョン）
func PopCount24(x uint64) int {
	var cnt uint64
	for i := 0; i < 64; i++ {
		cnt += x & 1
		x = x >> 1
	}
	return int(cnt)
}

// PopCount25 は x のポピュレーションカウント（1 が設定されているビット数）を返します（練習問題2.5バージョン）
func PopCount25(x uint64) int {
	cnt := 0
	for x != 0 {
		x = x & (x - 1)
		cnt++
	}
	return cnt
}
