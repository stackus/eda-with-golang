package sec

import (
	"context"

	"eda-in-golang/ch9/internal/am"
)

const (
	SagaCommandIDHdr   = am.CommandHdrPrefix + "SAGA_ID"
	SagaCommandNameHdr = am.CommandHdrPrefix + "SAGA_NAME"

	SagaReplyIDHdr   = am.ReplyHdrPrefix + "SAGA_ID"
	SagaReplyNameHdr = am.ReplyHdrPrefix + "SAGA_NAME"
)

type (
	SagaContext[T any] struct {
		ID           string
		Data         T
		Step         int
		Done         bool
		Compensating bool
	}

	SagaStore interface {
		Load(ctx context.Context, sagaName, sagaID string) (*SagaContext[[]byte], error)
		Save(ctx context.Context, sagaName string, sagaCtx *SagaContext[[]byte]) error
	}

	Saga[T any] interface {
		AddStep() SagaStep[T]
		Steps() []SagaStep[T]
		Name() string
		ReplyTopic() string
	}

	saga[T any] struct {
		name       string
		replyTopic string
		steps      []SagaStep[T]
	}
)

const (
	notCompensating = false
	isCompensating  = true
)

func NewSaga[T any](name, replyTopic string) Saga[T] {
	return &saga[T]{
		name:       name,
		replyTopic: replyTopic,
	}
}

func (s *saga[T]) AddStep() SagaStep[T] {
	step := &sagaStep[T]{
		actions: map[bool]StepActionFunc[T]{
			notCompensating: nil,
			isCompensating:  nil,
		},
		handlers: map[bool]map[string]StepReplyHandlerFunc[T]{
			notCompensating: {},
			isCompensating:  {},
		},
	}

	s.steps = append(s.steps, step)

	return step
}

func (s *saga[T]) Steps() []SagaStep[T] {
	return s.steps
}

func (s *saga[T]) Name() string {
	return s.name
}

func (s *saga[T]) ReplyTopic() string {
	return s.replyTopic
}

func (s *SagaContext[T]) advance(steps int) {
	var dir = 1
	if s.Compensating {
		dir = -1
	}

	s.Step += dir * steps
}

func (s *SagaContext[T]) complete() {
	s.Done = true
}

func (s *SagaContext[T]) compensate() {
	s.Compensating = true
}
