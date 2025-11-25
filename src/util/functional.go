package util


type (
	ForEachFunction[T any] func(T)
 	ForEachErrorFunction[T any] func(T) error
 	ForEachStoppingFunction[T any] func(T) bool
)

func PermutationDo[T any](k int, values []T, doer func(permutation []T)) {
	c := []int{}
	for i := 0; i < k; i++ {
		c = append(c, 0)
	}

	array := values

	doer(array)

	i := 1
	for i < k {
		if c[i] >= i {
			c[i] = 0
			i += 1
			continue
		}

		if i%2 == 0 {
			array[0], array[i] = array[i], array[0]
		} else {
			array[c[i]], array[i] = array[i], array[c[i]]
		}
		doer(append([]T(nil), array...))

		c[i] += 1
		i = 1
	}
}

func CombinationDo[T any](values []T, k int, doer func([]T)) {
	if k == 0 {
		doer([]T{})
		return;
	}
	
	for _, v := range values {
		CombinationDo(values, k-1, func(km1 []T) {
			doer(append([]T{v}, km1...))
		})
	}
}

func CombinationNoRepeatDo[T any](values []T, k int, doer func([]T)) {
	if k == 0 {
		doer([]T{})
		return;
	}

	if len(values) < k {
		return
	}

	for idx, v := range values {
		vals := append([]T{}, values[idx+1:]...)
		CombinationNoRepeatDo(vals, k-1, func(km1 []T) {
			doer(append([]T{v}, km1...))
		})
	}
}

func PairWiseDo[T any](values []T, doer func(a, b T)) {
	if len(values) < 2 {
		return
	}

	prev := values[0]
	for _, val := range values[1:] {
		doer(prev, val)
		prev = val
	}
}

func ForEach[T any](values []T, forEach ForEachFunction[T]) {
	for _, value := range values {
		forEach(value)
	}
}

func ForEachError[T any](values []T, forEach ForEachErrorFunction[T]) error {
	for _, value := range values {
		err := forEach(value)
		if err != nil {
			return err
		}
	}
	return nil
}

func ForEachStopping[T any](values []T, forEach ForEachStoppingFunction[T]) bool {
	for _, value := range values {
		if !forEach(value) {
			return true
		}
	}
	return false
}