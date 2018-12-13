package algo

import . "../structs"

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

// Core algorithm for calculating best (by alignment score) sequence matches.
// Given input sequence input.TargetSequence (with DotData) and sequence database,
// the task is to find several best database sequences, i.e. whose alignment has the greatest score.
func FASTA(input *InputBundle, db SequenceDb) FastaResult {
    entryNum     := len(db)
    bestResult   := AlignResult{ Score: -1 }
    bestSequence := ""

    for i := 0; i < entryNum; i += 1 {
        alignResult := fastaEntry(db[i].Sequence, db[i].DotData, input)
        if alignResult.Score > bestResult.Score {
            bestResult = *alignResult
            bestSequence = db[i].Sequence
        }
    }

    result := make([]FastaResultEntry, 1) /* TODO: generalize for input.BestMatchNum entries */

    result[0] = FastaResultEntry {
        DbSequence: bestSequence,
        Score: bestResult.Score,
    }

    return result
}

// FASTA iteration with input sequence input.TargetSequence and DB sequence sDb.
// Can be parallelized: fastaEntry() works independently and can be executed in individual thread.
func fastaEntry(sDb string, sDbDots *SeqDotData, input *InputBundle) *AlignResult {
    seqPair := SeqPair {
        S1: input.TargetSequence,
        S2: sDb,
    }

    diagDotData  := FormDiagonalDotData(input.TargetSeqDots, sDbDots, len(seqPair.S1), len(seqPair.S2))
    diags        := diagDotData.SelectBestDiagonals(input.DiagFilterNum)
    segs         := TrimToBestSegments(diags, &seqPair, input.WeightMat)
    filteredSegs := FilterByCutOff(segs, input.CutOff)

    if len(filteredSegs) == 0 {
        return &AlignResult{ Score: 0 }
    }

    strip := GetStripOf(filteredSegs, input.StripExtraWidth, &seqPair)
    score := SmithWatermanStrip(&seqPair, strip, input.WeightMat, input.GapPenalty)

    return score
}
