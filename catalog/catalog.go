package catalog

import (
	"slices"
)

type SortFunc[T any] func(T, T) int
type CreateFunc[T any] func([]string) T
type KeyFunc[K comparable, T any] func(*T) K

type Catalog[K comparable, T any] struct {
	actual_catalog []T
	Hash           map[K]*T
	View           []*T
}

func New[K comparable, T any](data [][]string, fcf CreateFunc[T], fkf KeyFunc[K, T]) Catalog[K, T] {
	cur_catalog := Catalog[K, T]{
		actual_catalog: []T{},
		Hash:           make(map[K]*T),
		View:           []*T{},
	}

	for i, d := range data[1:] {
		content := fcf(d)
		key := fkf(&content)
		cur_catalog.actual_catalog = append(cur_catalog.actual_catalog, content)
		cur_catalog.Hash[key] = &cur_catalog.actual_catalog[i]
		cur_catalog.View = append(cur_catalog.View, &cur_catalog.actual_catalog[i])
	}

	return cur_catalog
}

func (c Catalog[K, T]) Sort(f SortFunc[*T]) {
	slices.SortStableFunc(c.View, f)
}
