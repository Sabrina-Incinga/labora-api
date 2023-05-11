package rotateChain

func RotateChainLeft(chain string) string {
	var result string
	n := len(chain)

	for i := 0; i < n; i++ {
		result += string(chain[(n+i-1)%n])
	}

	return result
}

func RotateChainRight(chain string) string {
	var result string
	n := len(chain)

	for i := 0; i < n; i++ {
		result += string(chain[(n+i+1)%n])
	}

	return result
}

