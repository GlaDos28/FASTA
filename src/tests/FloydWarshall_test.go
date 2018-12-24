package tests

import "testing"

import (
   "../structs"
   "../algo"
)


func TestFloydWarshell1(t *testing.T){
  gap := -1
  e := 4
  testScore1 := 3
  testScore2 := 1

  seqs := make([]structs.Segment,4)

  seqs[0].Score = testScore2
  seqs[0].P1 = 0
  seqs[0].P2 = 1
  seqs[0].Diag = structs.Diagonal(2)

  seqs[1].Score = testScore1
  seqs[1].P1 = 3
  seqs[1].P2 = 6
  seqs[1].Diag = (structs.Diagonal)(-2)

  seqs[2].Score = testScore1
  seqs[2].P1 = 2
  seqs[2].P2 = 5
  seqs[2].Diag = (structs.Diagonal)(5)

  seqs[3].Score = testScore2
  seqs[3].P1 = 8
  seqs[3].P2 = 9
  seqs[3].Diag = (structs.Diagonal)(2)

  res := algo.FloydWarshall(seqs, gap, e)

  if len(res) != 1 {
		t.Error("Expected ", 1, ", got ", len(res))
	}

}

func TestFloydWarshell2(t *testing.T){
  gap := -1
  e := 2
  testScore := 5

  seqs := make([]structs.Segment,3)

  seqs[0].Score = testScore
  seqs[0].P1 = 3
  seqs[0].P2 = 5
  seqs[0].Diag = (structs.Diagonal)(-1)

  seqs[1].Score = testScore
  seqs[1].P1 = 2
  seqs[1].P2 = 4
  seqs[1].Diag = structs.Diagonal(5)

  seqs[2].Score = testScore
  seqs[2].P1 = 0
  seqs[2].P2 = 2
  seqs[2].Diag = structs.Diagonal(12)

  res := algo.FloydWarshall(seqs, gap, e)

  if len(res) != 1 {
		t.Error("Expected ", 1, ", got ", len(res))
	}
}
