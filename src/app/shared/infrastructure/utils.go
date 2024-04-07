package infrastructure

func Map[TSource Entity, TTarget any](source []*TSource, convert func(*TSource) *TTarget) []*TTarget {
	target := make([]*TTarget, 0, len(source))
	for _, s := range source {
		target = append(target, convert(s))
	}

	return target
}
