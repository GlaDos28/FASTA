package main

import (
    "./algo"
    "./structs"
    "fmt"
)

func main() {
    sequence     := "ATTG"
    weightMatrix := structs.Blosum62()

    var sequenceDb structs.SequenceDb

    // Fill input bundle

    input := structs.InputBundle {
        TargetSequence: sequence,
        TargetSeqDots: structs.BuildSeqDotDataFor(sequence),
        WeightMat: weightMatrix,
        GapPenalty: -1,
        DiagFilterNum: 10,
        CutOff: 26,
        StripExtraWidth: 4,
        BestMatchNum: 10,
    }

    // Call FASTA algorithm

    result := algo.FASTA(&input, sequenceDb)

    // Print result

    fmt.Println("Input sequence:")
    fmt.Println(sequence)

    for _, match := range result {
        fmt.Println()
        fmt.Println(match.DbSequence)
        fmt.Printf("Score: %d\n", match.Score)
    }
}
