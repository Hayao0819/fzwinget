package ui

import (
	fzf "github.com/ktr0731/go-fuzzyfinder"
	"github.com/sheepla/fzwinget/api"
)

type finderItem interface {
	Label() string
	Detail() string
}

func FindPackage(result api.SearchResult) ([]int, error) {
	var items []finderItem
	for _, p := range result {
		items = append(items, p)
	}

	return fzf.FindMulti(
		items,
		func(i int) string {
			return items[i].Label()
		},

		fzf.WithMode(fzf.ModeCaseInsensitive),
		fzf.WithPreviewWindow(func(i, width, height int) string {
			return items[i].Detail()
		}),
	)
}
