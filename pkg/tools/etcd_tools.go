package tools

import (
	"etcd"
)

const (
	EtcdErrorCodeNotFound      = 100
	EtcdErrorCodeTestFailed    = 101
	EtcdErrorCodeNodeExists    = 105
	EtcdErrorCodeValueRequired = 200
)

var (
	EtcdErrorNotFound      = &etcd.EtcdError{ErrorCode: EtcdErrorCodeNotFound}
	EtcdErrorTestFailed    = &etcd.EtcdError{ErrorCode: EtcdErrorCodeTestFailed}
	EtcdErrorNodeExists    = &etcd.EtcdError{ErrorCode: EtcdErrorCodeNodeExists}
	EtcdErrorValueRequired = &etcd.EtcdError{ErrorCode: EtcdErrorCodeValueRequired}
)
