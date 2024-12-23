package serializer

import "github.com/samber/lo"

func List[FROM any, TO any](items []FROM, serializer func(FROM) TO) []TO {
	return lo.Map(items, func(item FROM, index int) TO {
		return serializer(item)
	})
}
