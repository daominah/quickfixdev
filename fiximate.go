package quickfixdev

import (
	"strconv"
	"strings"

	"github.com/daominah/gomicrokit/textproc"
	"golang.org/x/net/html"
)

// FIXField is detail of a FIX protocol tag
type FIXField struct {
	Tag           int
	FieldName     string
	XMLName       string
	DataType      string
	UnionDataType string
	Description   string
	Added         string
	Deprecated    string
	RelateTag     string
}

// ParseFIXFieldsHtml _
func ParseFIXFieldsHtml(htmlStr string) ([]FIXField, error) {
	root, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		return nil, err
	}
	rows, err := textproc.HtmlXpath(root, `//tr[@xmlns]`)
	if err != nil {
		return nil, err
	}
	ret := make([]FIXField, 0)
	for _, row := range rows {
		cells, _ := textproc.HtmlXpath(row, `//td`)
		if len(cells) != 9 {
			continue
		}
		tag, _ := strconv.ParseInt(textproc.HtmlGetText(cells[0]), 10, 64)
		field := FIXField{
			Tag:           int(tag),
			FieldName:     textproc.HtmlGetText(cells[1]),
			XMLName:       textproc.HtmlGetText(cells[2]),
			DataType:      textproc.HtmlGetText(cells[3]),
			UnionDataType: textproc.HtmlGetText(cells[4]),
			Description:   textproc.HtmlGetText(cells[5]),
			Added:         textproc.HtmlGetText(cells[6]),
			Deprecated:    textproc.HtmlGetText(cells[7]),
			RelateTag:     textproc.HtmlGetText(cells[8]),
		}
		ret = append(ret, field)
	}
	return ret, nil
}

type msgType struct {
	MsgTypeSymbol string
	Name          string
	Description   string
}

// parseMsgTypesHtml _.
// :arg htmlStr: from https://fiximate.fixtrading.org/en/FIX.5.0SP2_EP254/tag35.htmls
func parseMsgTypesHtml(htmlStr string) ([]msgType, error) {
	root, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		return nil, err
	}
	rows, err := textproc.HtmlXpath(root, `//td[@class="FldDesEven"]//tr`)
	if err != nil {
		return nil, err
	}
	ret := make([]msgType, 0)
	for _, row := range rows {
		cells, _ := textproc.HtmlXpath(row, `//td`)
		if len(cells) != 6 {
			continue
		}
		name := textproc.HtmlGetText(cells[5])
		name = strings.Trim(name, "[]")
		field := msgType{
			MsgTypeSymbol: textproc.HtmlGetText(cells[0]),
			Name:          name,
			Description:   textproc.HtmlGetText(cells[2]),
		}
		ret = append(ret, field)
	}
	return ret, nil
}
