package iteration

func Repeat(character string, repeatCount int) string {
	var repeated string // Is declared as ""
	for i := 0; i < repeatCount; i++ {
		repeated = repeated + character
	}
	return repeated
}
