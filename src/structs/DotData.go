package structs

/* --- */

// Sequence dot Data stores for every combination of two symbols
// array of indices where this combination acts as substring.
// Structure implicitly relates to some sequence.
// Pair of symbols encoded in byte=uint8, so their production is uint16.
type SeqDotData map[uint16][]int

/* --- */

// Builds sequence dot data for given sequence.
func BuildSeqDotDataFor(seq string) *SeqDotData {
    return nil /* TODO */
}
