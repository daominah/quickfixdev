package quickfixdev

import (
	"testing"
)

func TestConfig(t *testing.T) {
	c0 := NewMockClientConfig()

	if c0.ToQuickFIXSetting() != `[DEFAULT]
SocketConnectHost=127.0.0.1
SocketConnectPort=5001
HeartBtInt=30
ReconnectInterval=5
SenderCompID=TechX
TargetCompID=HNX
MongoStoreConnection=mongodb://127.0.0.1:27017/
MongoStoreDatabase=fix_client
EnableLastMsgSeqNumProcessed=Y

[SESSION]
BeginString=FIX.4.4` {
		t.Error(c0.ToQuickFIXSetting())
	}
}
