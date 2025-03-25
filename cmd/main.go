package main

import (
	"fmt"
	"gogen/pkg/fatsa"
	fatsagc "gogen/pkg/fatsa/fatsa-gc"
	fatsakmer "gogen/pkg/fatsa/fatsa-kmer"
)

func main() {
	fatsaDataLocation := "tests/sequence.txt"

	fmt.Println("Parsing FASTA data...")
	fatsa.Parse(fatsaDataLocation)

	fmt.Printf("GC Content for %s: ", fatsaDataLocation)
	fmt.Println(fatsagc.CalculateGC(fatsaDataLocation))

	fmt.Printf("Kmer Count for %s: \n", fatsaDataLocation)
	kmerFrequencies := fatsakmer.KmerCount(fatsaDataLocation, 1)
	for kmer, count := range kmerFrequencies {
		fmt.Printf("%s: %d\n", kmer, count)
	}
}
