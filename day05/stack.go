package day05

type stack []string

func (s stack) Push(v string) stack {
	return append(s, v)
}

func (s stack) InvertedAndWrongPush(v string) stack {
	return append([]string{v}, s...)
}

func (s stack) Pop() (stack, string) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) Peek() string {
	l := len(s)
	return s[l-1]
}
