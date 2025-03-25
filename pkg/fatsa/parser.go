package fatsa

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/biogo/biogo/alphabet"
	"github.com/biogo/biogo/io/seqio"
	"github.com/biogo/biogo/io/seqio/fasta"
	"github.com/biogo/biogo/seq/linear"
)

func Parse(fatsaDataLocation string) {
	file, err := os.ReadFile(fatsaDataLocation)
	if err != nil {
		log.Fatal(err)
	}

	data := strings.NewReader(string(file))

	template := linear.NewSeq("", nil, alphabet.DNAredundant)
	r := fasta.NewReader(data, template)

	sc := seqio.NewScanner(r)

	for sc.Next() {
		// Get the current sequence and type assert to *linear.Seq.
		// While this is unnecessary here, it can be useful to have
		// the concrete type.
		s := sc.Seq().(*linear.Seq)

		// Print the sequence ID, description and sequence data.
		fmt.Printf("%q %q %s\n", s.ID, s.Desc, s.Seq)
	}
	if err := sc.Error(); err != nil {
		log.Fatal(err)
	}
}
