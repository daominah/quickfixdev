package quickfixdev

import (
	"encoding/json"
	"testing"
)

func TestParseFIXFieldsHtml(t *testing.T) {
	mapTagName, err := ParseFIXFieldsHtml(html0)
	if err != nil {
		t.Error(err)
	}
	if len(mapTagName) != 94 {
		t.Error(len(mapTagName))
	}
}

func TestParseMsgTypesHtml(t *testing.T) {
	msgTypes, err := ParseMsgTypesHtml(html1)
	if err != nil {
		t.Error(err)
	}
	if len(msgTypes) != 164 {
		t.Error(len(msgTypes))
	}
	if msgTypes[0].MsgTypeSymbol != "0" || msgTypes[0].Name != "Heartbeat" {
		t.Errorf("msgTypes[0]: %#v", msgTypes[0])
	}
}

// Print the map FIXMsgTypes
func _TestGenMapMsgTypes(t *testing.T) {
	msgTypes, err := ParseMsgTypesHtml(html1)
	if err != nil {
		t.Error(err)
	}
	ret := make(map[string]string)
	for _, v := range msgTypes {
		ret[v.MsgTypeSymbol] = v.Name
	}
	beauty, err := json.MarshalIndent(ret, "", "    ")
	t.Error(string(beauty))
}
