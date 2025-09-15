package protein

import (
	"errors"
)

var ErrStop = errors.New("stop codon encountered")
var ErrInvalidBase = errors.New("invalid base in codon")

func FromRNA(rna string) ([]string, error) {
	codons := []string{}
	for i := 0; i < len(rna); i += 3 {
		if i+2 >= len(rna) {
			return nil, ErrInvalidBase
		}
		codon := rna[i : i+3]
		if codon == "UAA" || codon == "UAG" || codon == "UGA" {
			break
		}
		protein, err := FromCodon(codon)
		if err != nil {
			return nil, err
		}
		codons = append(codons, protein)
	}
	return codons, nil
}

func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}
