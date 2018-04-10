package alg


var perms [][]rune

func heap(n int, s []rune) {
	if n == 1 {
		c := make([]rune, len(s))
		copy(c, s)
		perms = append(perms, c)
	} else {
		for i := 0; i < n-1; i++ {
			heap(n-1, s)
			if n%2 == 0 {
				s[i], s[n-1] = s[n-1], s[i]
			} else {
				s[0], s[n-1] = s[n-1], s[0]
			}
		}
		heap(n-1, s)
	}
}

//Atkin's sieve
func primes(amount int) []int {
	results := []int{2, 3, 5}
	sieve := make([]bool, amount)

	for n := range sieve {
		if n < 6 {
			continue
		}
		r := n % 60
		if r == 1 || r == 13 || r == 17 || r == 29 || r == 37 || r == 41 || r == 49 || r == 53 {
			for x := 0; 4*x*x <= n; x++ {
				for y := 0; y*y <= n; y++ {
					if opt := 4*x*x + (y * y); opt == n {
						sieve[n] = !sieve[n]
					}
				}
			}
		} else if r == 7 || r == 19 || r == 31 || r == 43 {
			for x := 0; 3*x*x <= n; x++ {
				for y := 0; y*y <= n; y++ {
					if opt := 3*x*x + (y * y); opt == n {
						sieve[n] = !sieve[n]
					}
				}
			}

		} else if r == 11 || r == 23 || r == 47 || r == 59 {
			//Added in an extra 2 factor to cancel out the y case so it does as many iterations as it needs to still.
			for x := 0; 3*x*x <= 2*n; x++ {
				for y := 0; y < x; y++ {
					if opt := 3*x*x - (y * y); opt == n {
						sieve[n] = !sieve[n]
					}
				}
			}
		}
	}

	for i, v := range sieve {
		if v {
			results = append(results, i)
			iSq := i * i
			for factor := 1; iSq*factor < len(sieve); factor++ {
				if factor == 5 || factor == 2 || factor == 3 {
					continue
				}
				sieve[iSq*factor] = false
			}
		}
	}

	return results
}

