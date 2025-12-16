package defaultmap

type (
	DefaultMap[K comparable, V any] struct {
		data map[K]V
		defaultFunc DefaultFunc[V]
	}

	DefaultFunc[V any] func() V
)

func ForKey[K comparable, V any](_ K, defaultFunc DefaultFunc[V]) DefaultMap[K, V] {
	return DefaultMap[K, V]{
		map[K]V{},
		defaultFunc,
	}
}

func (dm DefaultMap[K, V]) Len() int {
	return len(dm.data)
}

func (dm DefaultMap[K, V]) Set(key K, value V) (V, bool) {
	old, ok := dm.data[key]
	dm.data[key] = value
	return old, ok
}

func (dm DefaultMap[K, V]) Get(key K) V {
	value, ok := dm.data[key]
	if ok {
		return value
	}
	value = dm.defaultFunc()
	dm.data[key] = value
	return value
}

func (dm DefaultMap[K, V]) ForEach(forEach func(K, V)) {
	for key, value := range dm.data {
		forEach(key, value)
	}
}

func (dm DefaultMap[K, V]) Entries() map[K]V {
	return dm.data
}