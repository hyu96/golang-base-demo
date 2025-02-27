/*
 * Copyright (c) 2021 yedf. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package dtmimp

const (
	// ResultFailure for result of a trans/trans branch
	// Same as HTTP status 409 and GRPC code 10
	ResultFailure = "FAILURE"

	ResultIgnoreRollback = "IGNORE_ROLLBACK"

	// ResultSuccess for result of a trans/trans branch
	// Same as HTTP status 200 and GRPC code 0
	ResultSuccess = "SUCCESS"
	// ResultOngoing for result of a trans/trans branch
	// Same as HTTP status 425 and GRPC code 9
	ResultOngoing = "ONGOING"

	// OpTry branch type for TCC
	OpTry = "try"
	// OpConfirm branch type for TCC
	OpConfirm = "confirm"
	// OpCancel branch type for TCC
	OpCancel = "cancel"
	// OpAction branch type for message, SAGA, XA
	OpAction = "action"
	// OpCompensate branch type for SAGA
	OpCompensate = "compensate"
	// OpCommit branch type for XA
	OpCommit = "commit"
	// OpRollback branch type for XA
	OpRollback = "rollback"

	// DBTypeMysql util for driver mysql
	DBTypeMysql = "mysql"
	// DBTypePostgres util for driver postgres
	DBTypePostgres = "postgres"
	// DBTypeRedis util for driver redis
	DBTypeRedis = "redis"
	// Jrpc util for json-rpc
	Jrpc = "json-rpc"
	// JrpcCodeFailure util for json-rpc failure
	JrpcCodeFailure = -32901

	// JrpcCodeOngoing util for json-rpc ongoing
	JrpcCodeOngoing = -32902

	// MsgDoBranch0 util for DoAndSubmit barrier branch
	MsgDoBranch0 = "00"
	// MsgDoBarrier1 util for DoAndSubmit barrier barrierID
	MsgDoBarrier1 = "01"
	// MsgDoOp util for DoAndSubmit barrier op
	MsgDoOp = "msg"
	//MsgTopicPrefix util for Add topic msg
	MsgTopicPrefix = "topic://"

	// XaBarrier1 util for xa barrier id
	XaBarrier1 = "01"

	// ProtocolGRPC util for protocol grpc
	ProtocolGRPC = "grpc"
	// ProtocolHTTP util for protocol grpc
	ProtocolHTTP = "grpc"
)
