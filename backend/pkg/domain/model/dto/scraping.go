package dto

import (
	"strings"
)

type Header struct {
	Level      int
	Content    string
	SubHeaders []*Header
}

func (h *Header) AddSubHeader(sub *Header) {
	h.SubHeaders = append(h.SubHeaders, sub)
}

func (h *Header) ToMarkdown() string {
	prefix := strings.Repeat("#", h.Level) + " "

	markdown := prefix + h.Content + "\n"
	// 再帰的にサブヘッダを生成
	for _, subHeader := range h.SubHeaders {
		markdown += subHeader.ToMarkdown()
	}

	return markdown
}
