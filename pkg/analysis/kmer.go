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

func FATSAKmerCount(fatsaDataLocation string, k int) map[string]int {
	file, err := os.ReadFile(fatsaDataLocation)
	if err != nil {
		log.Fatal(err)
	}

	kmerCounts := make(map[string]int)

	data := strings.NewReader(string(file))

	template := linear.NewSeq("", nil, alphabet.DNAredundant)
	r := fasta.NewReader(data, template)

	sc := seqio.NewScanner(r)

	for sc.Next() {
		// Get the current sequence and type assert to *linear.Seq.
		// While this is unnecessary here, it can be useful to have
		// the concrete type.
		s := sc.Seq().(*linear.Seq)

		currentSeq := s.Seq.String()
		for i := 0; i < len(currentSeq)-k; i++ {
			kmer := currentSeq[i : i+k]
			kmerCounts[kmer]++
		}
	}
	if err := sc.Error(); err != nil {
		log.Fatal(err)
	}

	return kmerCounts
}
