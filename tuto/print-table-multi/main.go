package main

func main() {
	for l := 1; l <= 10; l++ {
		for c := 1; c <= 10; c++ {
			m := l * c
			switch {
			case m < 10:
				print("|  ")
			case m < 100:
				print("| ")
			default:
				print("|")
			}
			print(m)
		}
		print("|\n")
	}
}
