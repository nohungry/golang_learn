package main

func main() {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	for key, value := range m {
		println(key, value)
	}
}
