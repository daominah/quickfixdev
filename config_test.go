package quickfixdev

import (
	"testing"
)

func TestConfig(t *testing.T) {
	c0 := NewMockClientConfig()
	c0.Sessions[0].HNXVersion = "5.0"
	c0.Default.Client.LogonTimeout = 20
	if c0.ToQuickFIXSetting() != `[DEFAULT]
SocketConnectHost=127.0.0.1
SocketConnectPort=5001
HeartBtInt=30
ReconnectInterval=30
LogonTimeout=20
SenderCompID=TechX
TargetCompID=HNX
MongoStoreConnection=mongodb://127.0.0.1:27017/
MongoStoreDatabase=fix_client
EnableLastMsgSeqNumProcessed=Y

[SESSION]
BeginString=FIX.4.4
HNXVersion=5.0
` {
		t.Error(c0.ToQuickFIXSetting())
	}

	c1 := NewMockServerConfig()
	if c1.ToQuickFIXSetting() != `[DEFAULT]
SocketAcceptPort=5001
ResetOnLogon=N
SenderCompID=HNX
TargetCompID=TechX
MongoStoreConnection=mongodb://127.0.0.1:27017/
MongoStoreDatabase=fix_server
EnableLastMsgSeqNumProcessed=Y

[SESSION]
BeginString=FIX.4.4
` {
		t.Errorf(c1.ToQuickFIXSetting())
	}
}
