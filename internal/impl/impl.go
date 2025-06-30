package impl

import (
	"errors"
	"fmt"
)

type GqBuilderImpl struct {
	finalQuery string
	tmpQuery   [][]string
	err        error
}

func (gq *GqBuilderImpl) Select(tableName string, params ...string) *GqBuilderImpl {
	query := make([]string, 0, len(params)+1)
	query = append(query, "SELECT ")

	if len(params) != 0 {
		for _, v := range params {
			query = append(query, fmt.Sprintf("%s, ", v))
		}
	} else {
		query = append(query, "* ")
	}

	query = append(query, fmt.Sprintf("FROM %s ", tableName))

	gq.qAppend(query)
	return gq
}

func (gq *GqBuilderImpl) qAppend(queryPart []string) *GqBuilderImpl {
	if gq.tmpQuery == nil {
		gq.tmpQuery = make([][]string, 0)
	}

	gq.tmpQuery = append(gq.tmpQuery, queryPart)
	return gq
}

func (gq *GqBuilderImpl) Build() (string, error) {
	if gq.tmpQuery == nil {
		return "", errors.New("query is not init")
	}

	//TODO: err parsing

	var resultQuery string
	buildString(gq.tmpQuery, &resultQuery)

	return resultQuery, nil
}

func buildString(sl [][]string, value *string) {
	for _, s1 := range sl {
		for _, s2 := range s1 {
			*value += s2
		}
	}
}
