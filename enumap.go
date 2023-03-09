package enumstr

import (
	"github.com/semichkin-gopkg/conf"
	"github.com/semichkin-gopkg/mapping"
)

type e interface{ ~int32 }
type s interface{ ~string }

type Mapping[E e, S s] struct {
	m *mapping.Mapping[E, S]
}

func New[E e, S s](
	initialMap map[int32]string,
	configurators ...conf.Updater[Config[E, S]],
) *Mapping[E, S] {
	config := conf.NewBuilder[Config[E, S]]().
		Append(WithFormatter[E, S](FormatterLower)).
		Append(configurators...).
		Build()

	if config.Formatter == nil {
		config.Formatter = func(input string) string { return input }
	}

	preparedMap := map[E]S{}
	for k, v := range initialMap {
		preparedMap[E(k)] = S(config.Formatter(v))
	}

	defaultStr, _ := preparedMap[config.DefaultEnum]

	return &Mapping[E, S]{
		m: mapping.New(
			preparedMap,
			mapping.WithDefaultLeft[E, S](config.DefaultEnum),
			mapping.WithDefaultRight[E, S](defaultStr),
		),
	}
}

func (m *Mapping[E, S]) ToStr(val E) S {
	return m.m.ToRight(val)
}

func (m *Mapping[E, S]) ToEnum(val S) E {
	return m.m.ToLeft(val)
}
