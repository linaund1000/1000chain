package validators

//just for insert functions
import (
	"sort"
)
func insertFloat64(ss []float64, s float64) []float64 {
	i := sort.SearchFloat64s(ss, s)
	ss = append(ss, 0)
	copy(ss[i+1:], ss[i:])
	ss[i] = s
	return ss
}

func InsertSTR(ss []string, s string) []string {
	i := sort.SearchStrings(ss, s)
	ss = append(ss, "")
	copy(ss[i+1:], ss[i:])
	ss[i] = s
	return ss
}