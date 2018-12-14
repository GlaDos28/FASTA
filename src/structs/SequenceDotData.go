package structs

import "../util"

/* --- */

// Sequence dot Data stores for every combination of two symbols
// array of indices where this combination acts as substring.
// Structure implicitly relates to some sequence.
// Pair of symbols encoded in byte=uint8, so their production is uint16.
type SeqDotData [][]int

/* --- */

// Builds sequence dot data for given sequence.
func BuildSeqDotDataFor(seq string) *SeqDotData {
    n := len(seq) - 1
    data := make(SeqDotData, 256 * 256)

    for i := 0; i < n; i += 1 {
        key   := util.CombineSymbolPair(seq[i], seq[i + 1])
        value := data[key]

        if value != nil {
            data[key] = append(value, i)
        } else {
            data[key] = []int{ i }
        }
    }

    return &data
}
