package iteration

const count = 5

func Repeat(ch string) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += ch
	}
	return repeated
}
