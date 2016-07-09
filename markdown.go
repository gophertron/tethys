package main

import (
	"bytes"
	"fmt"
	"github.com/russross/blackfriday"
	"path"
	"strings"
)

func Markdown(data []byte) []byte {
	htmlFlags := 0
	renderer := &renderer{Html: blackfriday.HtmlRenderer(htmlFlags, "", "").(*blackfriday.Html)}
	extensions := 0
	return blackfriday.Markdown(data, renderer, extensions)
}

type renderer struct {
	*blackfriday.Html
}

func externalLink(content string, link string) string {
	return fmt.Sprintf("<a href=\"%s\" class=\"external\" target=\"_blank\">%s</a>", link, content)
}

func internalLink(content string, link string) string {
	return fmt.Sprintf("<a href=\"%s\" class=\"internal\">%s</a>", link, content)
}

func interLinkFromContent(content string) string {
	name := strings.ToLower(strings.Replace(content, " ", "-", -1))
	return fmt.Sprintf("<a href=\"/wiki/%s\" class=\"internal\">%s</a>", name, content)
}

func (r *renderer) Link(out *bytes.Buffer, link []byte, title []byte, content []byte) {

	strLink := string(link)
	strContent := string(content)

	switch {
	case strLink == "-":
		fmt.Fprintf(out, interLinkFromContent(strContent))
	case strings.HasPrefix(strLink, "http") || strings.HasPrefix(strLink, "https"):
		fmt.Fprintf(out, externalLink(strContent, strLink))
	case strings.HasPrefix(strLink, "/wiki/"):
		fmt.Fprintf(out, internalLink(strContent, strLink))
	default:
		fmt.Fprintf(out, internalLink(strContent, path.Join("/wiki", strLink)))
	}

}
