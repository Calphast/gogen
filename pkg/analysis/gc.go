package analysis

import (
	"log"
	"os"
	"strings"

	"github.com/biogo/biogo/alphabet"
	"github.com/biogo/biogo/io/seqio"
	"github.com/biogo/biogo/io/seqio/fasta"
	"github.com/biogo/biogo/seq/linear"
)

func FATSACalculateGC(fatsaDataLocation string) float64 {
	file, err := os.ReadFile(fatsaDataLocation)
	if err != nil {
		log.Fatal(err)
	}

	var totalCountGC int
	var totalCount int

	data := strings.NewReader(string(file))

	template := linear.NewSeq("", nil, alphabet.DNAredundant)
	r := fasta.NewReader(data, template)

	sc := seqio.NewScanner(r)

	for sc.Next() {
		s := sc.Seq().(*linear.Seq)

		currentSeq := s.Seq.String()
		for i := 0; i < len(currentSeq); i++ {
			if currentSeq[i] == 'G' || currentSeq[i] == 'C' {
				totalCountGC++
			}
			totalCount++
		}
	}
	if err := sc.Error(); err != nil {
		log.Fatal(err)
	}

	var result float64 = float64(totalCountGC*100) / float64(totalCount)

	return result
}
