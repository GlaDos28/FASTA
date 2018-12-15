package structs

import (
    "../util"
)

const MaxSequenceLength = 2 << 17

/* --- */

// Data from DotMatrix about matched diagonal dots of SeqPair.
// Structure implicitly relates to some SeqPair and their SeqDotData's.
// For every diagonal (determined by offset) number of matches is provided;
// StartOffset keeps offset of array, i.e. Data indices start from 0,
// but diagonal offsets starts from StartOffset.
type DiagonalDotData struct {
    Data        [MaxSequenceLength]uint
    StartOffset int
    length      int
}

/* --- */

// Calculates DiagonalDotData by given SeqDotData of each sequence S1 and S2.
// S1 length are used for determining array size and start diagonal offset.
// Result is written into given dddRef DiagonalDotData reference for memory issues.
func FormDiagonalDotData(dddRef *DiagonalDotData, s1Dots SeqDotData, s2 string, s1Len int) {

    // Initialize Data

    s2Len := len(s2)
    startOffset := -(s1Len - 1)

    dddRef.length = s1Len + s2Len - 1
    dddRef.StartOffset = startOffset

    for i := 0; i < dddRef.length; i += 1 {
        dddRef.Data[i] = 0
    }

    // Fill data

    for s2Ind := 0; s2Ind < s2Len - 1; s2Ind += 1 {
        key := util.CombineSymbolPair(s2[s2Ind], s2[s2Ind + 1])
        value := s1Dots[key]

        if value != nil {
            for _, s1Ind := range value {
                dddRef.Data[s2Ind - s1Ind - startOffset] += 1
            }
        }
    }
}

// Selects <amount> best (by dot match number) diagonals.
func (ddd *DiagonalDotData) SelectBestDiagonals(amount int) []Diagonal {

    // Store array of best values (and indices) with naive traverse and update.

    bestValues := make([]uint, amount + 1)
    bestValues[amount] = 1000000
    bestIndices := make([]Diagonal, amount)

    for i, j := 0, 0; i < ddd.length; i += 1 {
        j = 0
        for ; bestValues[j] < ddd.Data[i]; j += 1 {}
        j -= 1

        if j >= 0 {
            for k := 0; k < j; k += 1 {
                bestValues[k]  = bestValues[k + 1]
                bestIndices[k] = bestIndices[k + 1]
            }

            bestValues[j]  = ddd.Data[i]
            bestIndices[j] = Diagonal(i + ddd.StartOffset)
        }
    }

    // Return result

    return bestIndices
}
