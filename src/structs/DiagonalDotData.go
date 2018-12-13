package structs

import "sort"
import "../util"

/* --- */

// Data from DotMatrix about matched diagonal dots of SeqPair.
// Structure implicitly relates to some SeqPair and their SeqDotData's.
// For every diagonal (determined by offset) number of matches is provided;
// StartOffset keeps offset of array, i.e. Data indices start from 0,
// but diagonal offsets starts from StartOffset.
type DiagonalDotData struct {
    Data        []uint
    StartOffset int
}

/* --- */

// Calculates DiagonalDotData by given SeqDotData of each sequence S1 and S2.
// Lengths are used for determining array size and start diagonal offset.
func FormDiagonalDotData(s1Dots, s2Dots *SeqDotData, s1Len, s2Len int) *DiagonalDotData {

    // Initialize Data

    ddd := DiagonalDotData {
        make([]uint, s1Len + s2Len - 1),
        -(s1Len - 1),
    }

    // Fill Data

    for k, v1 := range *s1Dots {
        v2 := (*s2Dots)[k]

        if v2 != nil {
            for i := range v1 {
                for j := range v2 {
                    ddd.Data[i - j] += 1
                }
            }
        }
    }

    // Return result

    return &ddd
}

// Selects <amount> best (by dot match number) diagonals.
func (ddd *DiagonalDotData) SelectBestDiagonals(amount int) []Diagonal {

    // Initialize index array for getting best indices (i.e. diagonals)

    indices := make([]Diagonal, len(ddd.Data))

    for i := 0; i < len(indices); i += 1 {
        indices[i] = Diagonal(i + ddd.StartOffset)
    }

    // Sort according to ddd.Data descending order
    sort.Slice(indices, func (i, j int) bool {
        return ddd.Data[i] > ddd.Data[j]
    })

    // Prepare and return result

    result := make([]Diagonal, util.MinInt(amount, len(ddd.Data)))
    copy(indices, result)

    return result
}
