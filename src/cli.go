package main

import (
    "./algo"
    "./structs"
    "./util"
    "fmt"
    "os"
)

func main() {
    if len(os.Args) <= 1 {
        fmt.Println("Usage: go run cli.go <converted DB clusters directory path>")
        return
    }

    sequence     := "MASSINGRKPSEIFKAQALLYKHIYAFIDSMSLKWAVEMNIPNIIQNHGKPISLSNLVSILQVPSSKIGNVRRLMRYLAHNGFFEIITKEEESYALTVASELLVRGSDLCLAPMVECVLDPTLSGSYHELKKWIYEEDLTLFGVTLGSGFWDFLDKNPEYNTSFNDAMASDSKLINLALRDCDFVFDGLESIVDVGGGTGTTAKIICETFPKLKCIVFDRPQVVENLSGSNNLTYVGGDMFTSIPNADAVLLKYILHNWTDKDCLRILKKCKEAVTNDGKRGKVTIIDMVIDKKKDENQVTQIKLLMDVNMACLNGKERNEEEWKKLFIEAGFQHYKISPLTGFLSLIEIYP"
    dbDirPath    := os.Args[1]
    weightMatrix := structs.Blosum62()

    sequenceDb := structs.FromClusters(dbDirPath)
    fmt.Println("Clusters were successfully loaded")

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

    t1 := util.CurTime()
    result := algo.FASTA(&input, sequenceDb)
    t2 := util.CurTime()

    timeNano := t2 - t1

    // Print result

    fmt.Println("Input sequence:")
    fmt.Println(sequence)
    fmt.Println("Converted DB clusters directory path:")
    fmt.Println(dbDirPath)

    fmt.Println()
    fmt.Println("FASTA result:")

    for _, match := range result {
        fmt.Println()
        fmt.Println(match.DbSequence)
        fmt.Printf("Score: %d\n", match.Score)
    }

    fmt.Printf("\nTotal time: %.3f sec\n", float64(timeNano / 1000000) / 1000)
}
