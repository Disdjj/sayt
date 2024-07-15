package internal

func ShowInTerminal(ch chan string) {
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				return
			}
			// debug
			print(msg)
		}
	}
}
