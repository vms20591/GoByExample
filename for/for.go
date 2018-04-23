package main

func main() {
	i := 0
	for i < 10 {
		i++
	}

	for j := 0; j < 10; j++ {

	}

	k := 0
	for {
		if k >= 10 {
			break
		}

		k++
	}

	for l := 0; l < 10; l++ {
		if l%2 == 0 {
			continue
		}
	}
}
