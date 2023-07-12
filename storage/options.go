package storage

import commonv1 "github.com/nodeum-io/nodeum-proto/nodeum/common/v1"

type WriterOptions struct {
	// Exclusive will fail the writing if the destination already exists
	Exclusive bool
	// NodeInfo indicate metadata that also need to be applied.
	NodeInfo *commonv1.NodeInfo
}

type WriterOption func(o *WriterOptions)

func WithExclusive(v bool) WriterOption {
	return func(o *WriterOptions) {
		o.Exclusive = v
	}
}

func WithNodeInfo(v *commonv1.NodeInfo) WriterOption {
	return func(o *WriterOptions) {
		o.NodeInfo = v
	}
}
