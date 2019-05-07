package tmpl

import (
	"book/conf"
	"book/util"
	"bufio"
	"io"
	"os"
	"strings"
)

type html struct {
	html_head string
	html_body string
	html_foot string
}

type htmlAction interface {
	head(st string) string
	foot(st string) string
	body(st string) string
	Show(t string) string
}

var (
	h_header string
	h_footer string
)

func init() {
	filePath := conf.ROOT + "view/header.html"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 066)
	util.ShowError(err)

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		con, err := reader.ReadString('\n')
		h_header += con
		if err == io.EOF {
			break
		}
	}
	h_footer = "</div></body></html>"
}

func NewHtml() html {
	var h html;
	h.html_head = h_header;
	h.html_foot = h_footer;
	return h;
}

func (h *html) head(st string) string {
	if len(st) > 0 {
		h.html_head = st;
	}
	return h.html_head
}

func (h *html) foot(st string) string {
	if len(st) > 0 {
		h.html_foot = st;
	}
	return h.html_foot
}

func (h *html) body(st string) string {
	if len(st) > 0 {
		h.html_body = st;
	}
	return h.html_body
}

func (h *html) Show(t string) string {
	return strings.Replace(h.html_head,"{{title}}",t,-1) + "\n" + h.html_body + "\n" + h.html_foot
}
