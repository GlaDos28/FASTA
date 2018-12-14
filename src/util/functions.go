package util

import "time"

func MinInt(a, b int) int {
    if a < b { return a } else { return b }
}

func MaxInt(a, b int) int {
    if a > b { return a } else { return b }
}

func Max4(a, b, c, d int) int {
    return MaxInt(MaxInt(MaxInt(a, b), c), d)
}

func CombineSymbolPair(s1, s2 byte) uint16 {
    return (uint16(s1) << 8) | uint16(s2)
}

func CurTime() int64 {
    return time.Now().UnixNano()
}
