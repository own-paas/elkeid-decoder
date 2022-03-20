package parse

func Decoding(source []byte) []string {
	var (
		ch  byte
		sid int
		eid int
	)

	result := []string{}
	size := len(source)
	sid = 0
	eid = 0

	for i, v := range source {
		eid = i
		ch = v
		eid++
		if ch == 30 {
			result = append(result, string(source[sid:eid-1]))
			sid = eid
		}

		if eid == size && eid > sid {
			result = append(result, string(source[sid:eid-1]))
		}
	}

	return result
}
