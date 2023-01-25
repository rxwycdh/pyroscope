package stackbuilder

import (
	"context"
	"github.com/pyroscope-io/pyroscope/pkg/ingestion"
	"github.com/pyroscope-io/pyroscope/pkg/storage"
	"github.com/pyroscope-io/pyroscope/pkg/storage/tree"
)



type SamplesAppender interface {
	Append(stackID, value uint64)
}

type Label struct{ Key, Value string }
type Labels []Label

type WriteBatch interface {
	StackBuilder() tree.StackBuilder
	SamplesAppender(startTime, endTime int64, labels Labels) SamplesAppender
	Flush()
}

type WriteBatchFactory interface {
	NewWriteBatch(appName string) (WriteBatch, error)
}

type WriteBatchParser interface {
	ParseWithWriteBatch(c context.Context,
		wbf WriteBatchFactory,
		fallback storage.Putter,
		md ingestion.Metadata) error
}
