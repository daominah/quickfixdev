package quickfixdev

import (
	"fmt"
	"testing"
)

func TestHumanString(t *testing.T) {
	fixMsg := fmt.Sprintf("%v%v%v",
		"8=FIX.4.49=12535=D34=249=TW52=20200109-09:21:35.81656=ISLD",
		"11=521=138=11.0040=244=111.0054=155=VNM59=060=20200109-09:21:21.530",
		"10=225")
	humanStr := HumanString(fixMsg)
	if humanStr != fmt.Sprintf("%v%v%v%v%v",
		"8(BeginString)=FIX.4.49(BodyLength)=12535(MsgType)=D(NewOrderSingle)",
		"34(MsgSeqNum)=249(SenderCompID)=TW52(SendingTime)=20200109-09:21:35.816",
		"56(TargetCompID)=ISLD11(ClOrdID)=521(HandlInst)=138(OrderQty)=11.00",
		"40(OrdType)=244(Price)=111.0054(Side)=155(Symbol)=VNM59(TimeInForce)=0",
		"60(TransactTime)=20200109-09:21:21.53010(CheckSum)=225") {
		t.Error(humanStr)
	}
}
