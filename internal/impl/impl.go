package impl

import (
	"errors"
	"strings"
)

type QueryPart struct {
	clause string
	parts  []string
	args   []any
}

type GqBuilderImpl struct {
	parts []*QueryPart
	err   error
}

func (gq *GqBuilderImpl) Select(columns ...string) *GqBuilderImpl {
	part := &QueryPart{
		clause: "SELECT",
		parts:  []string{"SELECT "},
	}

	if len(columns) > 0 {
		part.parts = append(part.parts, strings.Join(columns, ", "))
	} else {
		part.parts = append(part.parts, "* ")
	}

	gq.parts = append(gq.parts, part)
	return gq
}

func (gq *GqBuilderImpl) From(tableName string) *GqBuilderImpl {
	part := &QueryPart{
		clause: "FROM",
		parts:  []string{"FROM "},
	}

	if len(tableName) > 0 {
		part.parts = append(part.parts, tableName)
	} else {
		if gq.err == nil {
			gq.err = errors.New("table name is not provide")
			return gq
		}
	}

	gq.parts = append(gq.parts, part)
	return gq
}

func (gq *GqBuilderImpl) Where(conditions string, args ...any) *GqBuilderImpl {

	part := &QueryPart{
		clause: "WHERE",
		parts:  []string{"WHERE "},
		args:   args,
	}

	if len(conditions) == 0 {
		if gq.err == nil {
			gq.err = errors.New("conditions in WHERE statement is not provide")
			return gq
		}
	}

	part.parts = append(part.parts, conditions)
	gq.parts = append(gq.parts, part)
	return gq
}

func (gq *GqBuilderImpl) Build() (string, any, error) {
	if len(gq.parts) == 0 {
		return "", nil, errors.New("query is not init")
	}

	if gq.err != nil {
		return "", nil, gq.err
	}

	var (
		sqlQuery strings.Builder
		args     []any
	)

	for _, part := range gq.parts {
		sqlQuery.WriteString(strings.Join(part.parts, ""))
		sqlQuery.WriteString(" ")
		args = append(args, part.args...)
	}

	return strings.TrimSpace(sqlQuery.String()), args, nil
}
