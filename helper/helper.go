package helper

import 	"strconv"

func ConvertAtoI(a string) (i int) {
	convert, _ := strconv.Atoi(a)
	return convert
}
