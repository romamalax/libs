package convert

func BytesToInt(in []byte) int {
	out := 0
	t := byte(0)
	for i := 0; i < len(in); i++ {
		t = in[i]
		if t < '0' || t > '9' {
			return 0
		}
		out = out*10 + int(t-'0')
	}
	return out
}
