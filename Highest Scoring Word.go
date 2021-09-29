package main

func High(s string) string {
	m, pre, at := 0, -1, 0
	mx := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if m < at {
				m = at
				mx = s[pre+1 : i]
			}
			pre = i
			at = 0
		} else {
			at += int(s[i] - 'a' + 1)
		}
	}
	if m < at {
		mx = s[pre+1:]
	}
	return mx
}

func main() {
	print(High("hyxxmrlf giagpesorb xqija rhjg pnousyz xbmpebxbpm nih hovxhl clet cszlw ojnodfcngi rtepyjramb zkovuv xeyf zfgzajfrw ukiboi dmuuszhea udtufqre jgaz hsfa"))
}
