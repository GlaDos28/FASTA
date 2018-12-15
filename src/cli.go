package main

import (
    "./algo"
    "./structs"
    "./util"
    "fmt"
    "os"
    "runtime"
    "sort"
)

const GoRoutineNum = 128

func main() {
    if len(os.Args) <= 1 {
        fmt.Println("Usage: go run cli.go <converted DB clusters directory path>")
        return
    }

    sequence     := "MASSINGRKPSEIFKAQALLYKHIYAFIDSMSLKWAVEMNIPNIIQNHGKPISLSNLVSILQVPSSKIGNVRRLMRYLAHNGFFEIITKEEESYALTVASELLVRGSDLCLAPMVECVLDPTLSGSYHELKKWIYEEDLTLFGVTLGSGFWDFLDKNPEYNTSFNDAMASDSKLINLALRDCDFVFDGLESIVDVGGGTGTTAKIICETFPKLKCIVFDRPQVVENLSGSNNLTYVGGDMFTSIPNADAVLLKYILHNWTDKDCLRILKKCKEAVTNDGKRGKVTIIDMVIDKKKDENQVTQIKLLMDVNMACLNGKERNEEEWKKLFIEAGFQHYKISPLTGFLSLIEIYP"
    dbDirPath    := os.Args[1]
    weightMatrix := structs.Blosum62()
    //sequenceDb   := structs.DbBySequences([]string{
    //    "MASSINGRKPSEIFKAQALLYKHIYAFIDSMSLKWAVEMNIPNIIQNHGKPISLSNLVSILQVPSSKIGNVRRLMRYLAHNGFFEIITKEEESYALTVASELLVRGSDLCLAPMVECVLDPTLSGSYHELKKWIYEEDLTLFGVTLGSGFWDFLDKNPEYNTSFNDAMASDSKLINLALRDCDFVFDGLESIVDVGGGTGTTAKIICETFPKLKCIVFDRPQVVENLSGSNNLTYVGGDMFTSIPNADAVLLKYILHNWTDKDCLRILKKCKEAVTNDGKRGKVTIIDMVIDKKKDENQVTQIKLLMDVNMACLNGKERNEEEWKKLFIEAGFQHYKISPLTGFLSLIEIYP",
    //})
    sequenceDb := structs.FromClusters(dbDirPath)

    fmt.Println("DB sequences were successfully loaded")

    // Fill input bundle

    input := structs.InputBundle {
        TargetSequence: sequence,
        TargetSeqDots: structs.BuildSeqDotDataFor(sequence),
        WeightMat: weightMatrix,
        GapPenalty: -1,
        DiagFilterNum: 10,
        DotMatchCutOff: 11,
        CutOff: 26,
        StripExtraWidth: 0,
        BestMatchNum: 1,
    }

    // Call FASTA algorithm

    cReady  := make(chan *algo.FastaResult, GoRoutineNum)
    results := make([]*algo.FastaResult, 0, GoRoutineNum)
    chunk   := len(sequenceDb) / GoRoutineNum

    runtime.GOMAXPROCS(GoRoutineNum)

    t1 := util.CurTime()

    for i := 0; i < GoRoutineNum; i++ {
        go func(start int) {
            end := start + chunk

            if end > len(sequenceDb) {
                end = len(sequenceDb)
            }

            res := algo.FASTA(&input, sequenceDb[start : end])

            cReady <- &res
        }(i * chunk)
    }

    for i := 0; i < GoRoutineNum; i++ {
        res := <-cReady
        results = append(results, res)
    }

    t2 := util.CurTime()

    timeNano := t2 - t1

    // Print result

    fmt.Println("Input sequence:")
    fmt.Println(sequence)
    fmt.Println("Converted DB clusters directory path:")
    fmt.Println(dbDirPath)

    fmt.Println()
    fmt.Println("FASTA result:")

    bestResEntries := make([]algo.FastaResultEntry, input.BestMatchNum * GoRoutineNum)

    for i, result := range results {
        for j, entry := range *result {
            bestResEntries[i * input.BestMatchNum + j] = entry
        }
    }

    sort.Slice(bestResEntries, func (i, j int) bool {
        return bestResEntries[i].Score > bestResEntries[j].Score
    })

    for i := 0; i < input.BestMatchNum; i += 1 {
        fmt.Println()
        fmt.Println(bestResEntries[i].DbSequence)
        fmt.Printf("Score: %d\n", bestResEntries[i].Score)
    }

    fmt.Printf("\nTotal time: %.3f sec\n", float64(timeNano / 1000000) / 1000)
}
