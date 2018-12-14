package algo

import (
    . "../structs"
    "../util"
    "fmt"
)

/* --- */

// Data returned by FASTA algorithm
type FastaResult []FastaResultEntry

// Data entry of FastaResult structure.
// Signifies one of the best sequence matches.
type FastaResultEntry struct {
    DbSequence string
    Score      int
}

/* --- */

var C1 int64 = 0
var C2 int64 = 0
var C3 int64 = 0
var C4 int64 = 0
var C5 int64 = 0
var C6 int64 = 0

var PassedCutOffNum int64 = 0

// Core algorithm for calculating best (by alignment score) sequence matches.
// Given input sequence input.TargetSequence (with DotData) and sequence database,
// the task is to find several best database sequences, i.e. whose alignment has the greatest score.
func FASTA(input *InputBundle, db *SequenceDb) FastaResult {
    entryNum     := len(*db)
    bestResult   := AlignResult{ Score: -1 }
    bestSequence := ""

    for i := 0; i < entryNum; i += 1 {
        alignResult := fastaEntry((*db)[i].Sequence, input)
        if alignResult.Score > bestResult.Score {
            bestResult = *alignResult
            bestSequence = (*db)[i].Sequence
        }
    }

    fmt.Printf("Debug:\n\t%v (%v)\n\t%v (%v)\n\t%v (%v)\n\t%v (%v)\n\t%v (%v)\n\t%v (%v)\n%v passed to SW\n", C1 / int64(entryNum), C1, C2 / int64(entryNum), C2, C3 / int64(entryNum), C3, C4 / int64(entryNum), C4, C5 / PassedCutOffNum, C5, C6 / PassedCutOffNum, C6, PassedCutOffNum)

    result := make([]FastaResultEntry, 1) /* TODO: generalize for input.BestMatchNum entries */

    result[0] = FastaResultEntry {
        DbSequence: bestSequence,
        Score: bestResult.Score,
    }

    return result
}

// FASTA iteration with input sequence input.TargetSequence and DB sequence sDb.
// Can be parallelized: fastaEntry() works independently and can be executed in individual thread.
func fastaEntry(sDb string, input *InputBundle) *AlignResult {
    seqPair := SeqPair {
        S1: input.TargetSequence,
        S2: sDb,
    }

    t1 := util.CurTime()
    diagDotData  := FormDiagonalDotData(input.TargetSeqDots, seqPair.S2, len(seqPair.S1))
    t2 := util.CurTime()
    C1 += t2 - t1
    diags        := diagDotData.SelectBestDiagonals(input.DiagFilterNum)
    t3 := util.CurTime()
    C2 += t3 - t2
    segs         := TrimToBestSegments(diags, &seqPair, input.WeightMat)
    t4 := util.CurTime()
    C3 += t4 - t3
    filteredSegs := FilterByCutOff(segs, input.CutOff)
    t5 := util.CurTime()
    C4 += t5 - t4

    if len(filteredSegs) == 0 {
        return &AlignResult{ Score: 0 }
    }

    PassedCutOffNum += 1
    strip := GetStripOf(filteredSegs, input.StripExtraWidth, &seqPair)
    t6 := util.CurTime()
    C5 += t6 - t5
    score := SmithWatermanStrip(&seqPair, strip, input.WeightMat, input.GapPenalty)
    t7 := util.CurTime()
    C6 += t7 - t6

    return score
}
