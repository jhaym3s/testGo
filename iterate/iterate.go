package iterate

func Iterate(char string,times int) string {
var repeated string
	for i := 0; i < times; i++ {
		repeated = repeated + char
	}
	return repeated
}