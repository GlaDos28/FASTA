package structs

// Program input data. Principally stores data for FASTA algorithm.
type InputBundle struct {
    TargetSequence  string
    TargetSeqDots   *SeqDotData
    WeightMat       *WeightMatrix
    GapPenalty      int
    DiagFilterNum   int
    CutOff          int
    StripExtraWidth int
    BestMatchNum    int
}
