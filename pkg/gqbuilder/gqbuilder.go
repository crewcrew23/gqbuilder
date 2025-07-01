package gqbuilder

import (
	"github.com/crewcrew23/gqbuilder/internal/impl"
	"github.com/crewcrew23/gqbuilder/pkg/interfaces"
)

func Builder() interfaces.GqBuilder {
	return &impl.GqBuilderImpl{}
}
