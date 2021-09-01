package helper

import (
	"context"
	"strings"
	"github.com/99designs/gqlgen/graphql"
)

func GoStrctStrFormatter(inpt string) string {
	words := strings.Split(inpt, ".")
  for index, word := range words {
    rslt := strings.Title(word)
		if strings.Contains(rslt, "Id") {
			rslt = strings.Replace(rslt, "Id", "ID", -1)
		}
		words[index] = rslt
	}
	return strings.Join(words, ".")
}

func GetPreloads(ctx context.Context) []string {
	return GetNestedPreloads(
		graphql.GetOperationContext(ctx),
		graphql.CollectFieldsCtx(ctx, nil),
		"",
	)
}

func GetNestedPreloads(ctx *graphql.OperationContext, fields []graphql.CollectedField, prefix string) (preloads []string) {
	for _, column := range fields {
		prefixColumn := GetPreloadString(prefix, column.Name)
		// column.Selections is subfield
		if len(column.Selections)==0 {
			continue
		}
		preloads = delWord(preloads, findObj(prefixColumn))
		preloads = append(preloads, prefixColumn)
		preloads = append(preloads, GetNestedPreloads(ctx, graphql.CollectFields(ctx, column.Selections, nil), prefixColumn)...)
	}
	return
}

func GetPreloadString(prefix, name string) string {
	if len(prefix) > 0 {
		return GoStrctStrFormatter(prefix + "." + name)
	}
	return GoStrctStrFormatter(name)
}
