package main

func main() {
	find([]string{"a"}, "a")
}
func find(list []string, item string) string {
	for _, i := range list {
		if i == item {
			return i
		}
	}

	return ""
}
