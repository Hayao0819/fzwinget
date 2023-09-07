package ui

import (
	fzf "github.com/ktr0731/go-fuzzyfinder"
)

type FinderItem interface {
	Label() string
	Detail() string
}

type FinderList []FinderItem

func FindPackage(list FinderList) ([]int, error) {
	return fzf.FindMulti(
		list,
		func(i int) string {
			return list[i].Label()
		},

		fzf.WithMode(fzf.ModeCaseInsensitive),
		fzf.WithPreviewWindow(func(i, width, height int) string {
			return list[i].Detail()
		}),
	)
}
