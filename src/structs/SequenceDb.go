package structs

// Adapter for database of sequences and their dots data.
type SequenceDb []SequenceEntry

// Sequence database entry.
// Stores sequence in string form and data about precalculated dots.
type SequenceEntry struct {
    Sequence string
    DotData  *SeqDotData
}
