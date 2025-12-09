package strain

// Implement the "Keep" and "Discard" function in this file.
func Keep[T any](in []T, predicate func(T) bool) []T {
	// in the example they use comparable but here we use a predicate directly.
	out := make([]T, 0, len(in))
	for _, r := range in {
		if predicate(r) {
			out = append(out, r)
		}
	}
	return out
}

func Discard[T any](in []T, predicate func(T) bool) []T {
	out := make([]T, 0, len(in))
	for _, r := range in {
		if !predicate(r) {
			out = append(out, r)
		}
	}
	return out
}

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1
