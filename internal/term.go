package internal

func ShowInTerminal(ch chan string) {
	for s := range ch {
		print(s)
	}
}
