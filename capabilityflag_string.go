// Code generated by "stringer -type=CapabilityFlag"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CLIENT_LONG_PASSWORD-1]
	_ = x[CLIENT_FOUND_ROWS-2]
	_ = x[CLIENT_LONG_FLAG-4]
	_ = x[CLIENT_CONNECT_WITH_DB-8]
	_ = x[CLIENT_NO_SCHEMA-16]
	_ = x[CLIENT_COMPRESS-32]
	_ = x[CLIENT_ODBC-64]
	_ = x[CLIENT_LOCAL_FILES-128]
	_ = x[CLIENT_IGNORE_SPACE-256]
	_ = x[CLIENT_PROTOCOL_41-512]
	_ = x[CLIENT_INTERACTIVE-1024]
	_ = x[CLIENT_SSL-2048]
	_ = x[CLIENT_IGNORE_SIGPIPE-4096]
	_ = x[CLIENT_TRANSACTIONS-8192]
	_ = x[CLIENT_RESERVED-16384]
	_ = x[CLIENT_SECURE_CONNECTION-32768]
	_ = x[CLIENT_MULTI_STATEMENTS-65536]
	_ = x[CLIENT_MULTI_RESULTS-131072]
	_ = x[CLIENT_PS_MULTI_RESULTS-262144]
	_ = x[CLIENT_PLUGIN_AUTH-524288]
	_ = x[CLIENT_CONNECT_ATTRS-1048576]
	_ = x[CLIENT_PLUGIN_AUTH_LENENC_CLIENT_DATA-2097152]
	_ = x[CLIENT_CAN_HANDLE_EXPIRED_PASSWORDS-4194304]
	_ = x[CLIENT_SESSION_TRACK-8388608]
	_ = x[CLIENT_DEPRECATE_EOF-16777216]
}

const _CapabilityFlag_name = "CLIENT_LONG_PASSWORDCLIENT_FOUND_ROWSCLIENT_LONG_FLAGCLIENT_CONNECT_WITH_DBCLIENT_NO_SCHEMACLIENT_COMPRESSCLIENT_ODBCCLIENT_LOCAL_FILESCLIENT_IGNORE_SPACECLIENT_PROTOCOL_41CLIENT_INTERACTIVECLIENT_SSLCLIENT_IGNORE_SIGPIPECLIENT_TRANSACTIONSCLIENT_RESERVEDCLIENT_SECURE_CONNECTIONCLIENT_MULTI_STATEMENTSCLIENT_MULTI_RESULTSCLIENT_PS_MULTI_RESULTSCLIENT_PLUGIN_AUTHCLIENT_CONNECT_ATTRSCLIENT_PLUGIN_AUTH_LENENC_CLIENT_DATACLIENT_CAN_HANDLE_EXPIRED_PASSWORDSCLIENT_SESSION_TRACKCLIENT_DEPRECATE_EOF"

var _CapabilityFlag_map = map[CapabilityFlag]string{
	1:        _CapabilityFlag_name[0:20],
	2:        _CapabilityFlag_name[20:37],
	4:        _CapabilityFlag_name[37:53],
	8:        _CapabilityFlag_name[53:75],
	16:       _CapabilityFlag_name[75:91],
	32:       _CapabilityFlag_name[91:106],
	64:       _CapabilityFlag_name[106:117],
	128:      _CapabilityFlag_name[117:135],
	256:      _CapabilityFlag_name[135:154],
	512:      _CapabilityFlag_name[154:172],
	1024:     _CapabilityFlag_name[172:190],
	2048:     _CapabilityFlag_name[190:200],
	4096:     _CapabilityFlag_name[200:221],
	8192:     _CapabilityFlag_name[221:240],
	16384:    _CapabilityFlag_name[240:255],
	32768:    _CapabilityFlag_name[255:279],
	65536:    _CapabilityFlag_name[279:302],
	131072:   _CapabilityFlag_name[302:322],
	262144:   _CapabilityFlag_name[322:345],
	524288:   _CapabilityFlag_name[345:363],
	1048576:  _CapabilityFlag_name[363:383],
	2097152:  _CapabilityFlag_name[383:420],
	4194304:  _CapabilityFlag_name[420:455],
	8388608:  _CapabilityFlag_name[455:475],
	16777216: _CapabilityFlag_name[475:495],
}

func (i CapabilityFlag) String() string {
	if str, ok := _CapabilityFlag_map[i]; ok {
		return str
	}
	return "CapabilityFlag(" + strconv.FormatInt(int64(i), 10) + ")"
}
