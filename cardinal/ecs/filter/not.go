package filter

import (
	"pkg.world.dev/world-engine/cardinal/ecs/component"
)

type not struct {
	filter LayoutFilter
}

func Not(filter LayoutFilter) LayoutFilter {
	return &not{filter: filter}
}

func (f *not) MatchesLayout(components []component.IComponentType) bool {
	return !f.filter.MatchesLayout(components)
}
