package main

func asyncProcessor(ch chan string) {
	for m := range ch {
		print(m)
	}
}
