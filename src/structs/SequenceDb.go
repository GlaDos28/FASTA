package structs

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

const ClusterFileExtension = ".cl"
const SequenceInitCap      = 1000000

/* --- */

// Adapter for database of sequences and their dots data.
type SequenceDb []SequenceEntry

// Sequence database entry.
// Stores sequence in string form and data about precalculated dots.
type SequenceEntry struct {
    Sequence string
}

/* --- */

func FromClusters(clustersDirPath string) *SequenceDb {
    files, err := ioutil.ReadDir(clustersDirPath)

    if err != nil {
        panic(err)
    }

    sequences := make(SequenceDb, 0, SequenceInitCap)

    for _, f := range files {
        if strings.HasSuffix(f.Name(), ClusterFileExtension) {
            fmt.Printf("Processing cluster %v ...\n", f.Name())
            file, err := os.Open(clustersDirPath + "/" + f.Name())

            if err != nil {
                panic(err)
            }

            scanner := bufio.NewScanner(file)
            buf := make([]byte, 0, 1024 * 1024)
            scanner.Buffer(buf, 10 * 1024 * 1024)

            for scanner.Scan() {
                text  := scanner.Text()
                lines := strings.Split(text, "\n")

                for _, sequence := range lines {
                    sequences = append(sequences, SequenceEntry{ Sequence: string(sequence) })
                }
            }

            file.Close()
            //break /* TODO: remove */
        }
    }

    return &sequences
}

func SequencesToEntries(seqs []string) []SequenceEntry {
    entries := make([]SequenceEntry, len(seqs))

    for i, seq := range seqs {
        entries[i] = SequenceEntry{ Sequence: seq }
    }

    return entries
}
