package strand

func ToRNA(dna string) string {
	var rna string
	for _, c := range dna {
		switch c {
		case 'G':
			rna += "C"
		case 'C':
			rna += "G"
		case 'T':
			rna += "A"
		case 'A':
			rna += "U"
		}
	}
	return rna
}
