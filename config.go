package quickfixdev

import (
	"fmt"
	"strings"
)

// Config is golang structured config for quickfix app.
// Usage: quickfix.ParseSettings(strings.NewReader(thisConfig.ToQuickFIXSetting()))
type Config struct {
	Default struct {
		IsClient bool
		Client struct {
			SocketConnectHost string // ex: 127.0.0.1
			SocketConnectPort string // ex: 5001
			HeartBtInt        int    // heartbeat interval in seconds, ex: 30
			LogonTimeout      int    // measure in seconds, FIX default: 10
			// measure in seconds, should be longer than LogonTimeout, FIX default: 30
			ReconnectInterval int
		}
		Server struct {
			SocketAcceptPort string // ex: 5001
			ResetOnLogon     bool   // if sequence should be reset when in logon
		}

		SenderCompID string // ex: TechX
		TargetCompID string // ex: HNX

		MongoStoreConnection string // ex: mongodb://name:pass@host:port/
		MongoStoreDatabase   string // ex: mongodb database name

		EnableLastMsgSeqNumProcessed bool
	}
	Sessions []Session
}

// Session _
type Session struct {
	BeginString string // version of FIX, ex: FIX.4.4
	HNXVersion  string // specific to HaNoiStockExchange
}

func boolToYN(b bool) string {
	if b {
		return "Y"
	}
	return "N"
}

// ToQuickFIXSetting returns the config in quickfix format
func (c Config) ToQuickFIXSetting() string {
	ret := []string{"[DEFAULT]"}
	if c.Default.IsClient {
		ret = append(ret,
			fmt.Sprintf("SocketConnectHost=%v", c.Default.Client.SocketConnectHost),
			fmt.Sprintf("SocketConnectPort=%v", c.Default.Client.SocketConnectPort),
			fmt.Sprintf("HeartBtInt=%v", c.Default.Client.HeartBtInt),
		)
		if c.Default.Client.ReconnectInterval > 0 {
			ret = append(ret, fmt.Sprintf("ReconnectInterval=%v", c.Default.Client.ReconnectInterval))
		}
		if c.Default.Client.LogonTimeout > 0 {
			ret = append(ret, fmt.Sprintf("LogonTimeout=%v", c.Default.Client.LogonTimeout))
		}
	} else {
		ret = append(ret,
			fmt.Sprintf("SocketAcceptPort=%v", c.Default.Server.SocketAcceptPort),
			fmt.Sprintf("ResetOnLogon=%v", boolToYN(c.Default.Server.ResetOnLogon)),
		)
	}

	ret = append(ret,
		fmt.Sprintf("SenderCompID=%v", c.Default.SenderCompID),
		fmt.Sprintf("TargetCompID=%v", c.Default.TargetCompID),
		fmt.Sprintf("MongoStoreConnection=%v", c.Default.MongoStoreConnection),
		fmt.Sprintf("MongoStoreDatabase=%v", c.Default.MongoStoreDatabase),
		fmt.Sprintf("EnableLastMsgSeqNumProcessed=%v", boolToYN(c.Default.EnableLastMsgSeqNumProcessed)),
	)
	for _, session := range c.Sessions {
		ret = append(ret,
			"\n[SESSION]",
			fmt.Sprintf("BeginString=%v", session.BeginString),
		)
		if session.HNXVersion != "" {
			ret = append(ret, fmt.Sprintf("HNXVersion=%v", session.HNXVersion))
		}
	}

	return strings.Join(ret, "\n") + "\n"
}

// NewMockServerConfig is an example of acceptor config in quickfix
func NewMockServerConfig() Config {
	c0 := Config{}
	c0.Default.IsClient = false
	c0.Default.Server.SocketAcceptPort = "5001"
	c0.Default.Server.ResetOnLogon = false
	c0.Default.SenderCompID = "HNX"
	c0.Default.TargetCompID = "TechX"
	c0.Default.MongoStoreConnection = "mongodb://127.0.0.1:27017/"
	c0.Default.MongoStoreDatabase = "fix_server"
	c0.Default.EnableLastMsgSeqNumProcessed = true
	c0.Sessions = []Session{{BeginString: "FIX.4.4"}}
	return c0
}

// NewMockClientConfig is an example of initiator config in quickfix
func NewMockClientConfig() Config {
	c0 := Config{}
	c0.Default.IsClient = true
	c0.Default.Client.SocketConnectHost = "127.0.0.1"
	c0.Default.Client.SocketConnectPort = "5001"
	c0.Default.Client.HeartBtInt = 30
	c0.Default.Client.ReconnectInterval = 30
	c0.Default.Client.LogonTimeout = 10
	c0.Default.SenderCompID = "TechX"
	c0.Default.TargetCompID = "HNX"
	c0.Default.MongoStoreConnection = "mongodb://127.0.0.1:27017/"
	c0.Default.MongoStoreDatabase = "fix_client"
	c0.Default.EnableLastMsgSeqNumProcessed = true
	c0.Sessions = []Session{
		{BeginString: "FIX.4.4"},
	}
	return c0
}
