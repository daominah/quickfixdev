package quickfixdev

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/daominah/gomicrokit/log"
)

// FIXFields map tagNumber to fieldName,
// ex: 34: MsgSeqNum, 35: MsgType, ..
var FIXFields = make(map[string]string, 8192)

// FIXFieldsFull map tagNumber to detail of the tag,
// ex: 34: MsgSeqNum, an integer message sequence number, added in FIX.2.7
var FIXFieldsFull = make(map[int]FIXField, 8192)

func init() {
	var array []FIXField
	err := json.Unmarshal([]byte(jsonDataFIXFields), &array)
	if err != nil {
		log.Fatalf("error when read FIXFields: %v", err)
	}
	for _, v := range array {
		FIXFields[fmt.Sprintf("%v", v.Tag)] = v.FieldName
		FIXFieldsFull[v.Tag] = v
	}
	//log.Debugf("nFIXFields: %v", len(FIXFields))

}

// HumanString convert a FIX message to a human readable string
func HumanString(fixMsg string) string {
	separator := ""
	tagValues := strings.SplitAfter(fixMsg, separator)
	humanTags := make([]string, 0)
	for _, tagValue := range tagValues {
		if tagValue == "" {
			continue
		}
		// eqIdx is index of "=" symbol in the tagValue
		eqIdx := strings.Index(tagValue, "=")
		if eqIdx == -1 {
			continue
		}
		var humanTag string
		if tagValue[:eqIdx] == "35" { // MsgType
			if eqIdx+1 > len(tagValue)-1 {
				continue
			}
			humanTag = fmt.Sprintf("%v(%v)=%v(%v)",
				tagValue[:eqIdx], FIXFields[tagValue[:eqIdx]],
				tagValue[eqIdx+1:len(tagValue)-1],
				FIXMsgTypes[tagValue[eqIdx+1:len(tagValue)-1]])
		} else {
			humanTag = fmt.Sprintf("%v(%v)%v",
				tagValue[:eqIdx], FIXFields[tagValue[:eqIdx]], tagValue[eqIdx:])
		}
		humanTags = append(humanTags, humanTag)
	}
	return strings.Join(humanTags, "")
}
