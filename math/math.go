package math

// Computes GCD using Euclidean algorithm
func GCD(a, b int) int {
    if (b == 0) {
        return a;
    }
    return GCD(b, a % b);
}

func LCM(a, b int) int {
	return (a * b) / GCD(a, b)
}
