package infrastructure

import "time"

func Map[TSource Entity, TTarget any](source []*TSource, convert func(*TSource) *TTarget) []*TTarget {
	target := make([]*TTarget, 0, len(source))
	for _, s := range source {
		target = append(target, convert(s))
	}

	return target
}

func TimeToIsoString(t time.Time) string {
	return t.Format(time.RFC3339)
}

func TimeFromIsoString(str string) (time.Time, error) {
	return time.Parse(time.RFC3339, str)
}
