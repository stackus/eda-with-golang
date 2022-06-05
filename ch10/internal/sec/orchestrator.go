package sec

import (
	"context"

	"github.com/stackus/errors"

	"eda-in-golang/internal/am"
	"eda-in-golang/internal/ddd"
)

type (
	Orchestrator[T any] interface {
		Start(ctx context.Context, id string, data T) error
		ReplyTopic() string
		HandleReply(ctx context.Context, reply ddd.Reply) error
	}

	orchestrator[T any] struct {
		saga      Saga[T]
		repo      SagaRepository[T]
		publisher am.CommandPublisher
	}
)

var _ Orchestrator[any] = (*orchestrator[any])(nil)

func NewOrchestrator[T any](saga Saga[T], repo SagaRepository[T], publisher am.CommandPublisher) Orchestrator[T] {
	return orchestrator[T]{
		saga:      saga,
		repo:      repo,
		publisher: publisher,
	}
}

func (o orchestrator[T]) Start(ctx context.Context, id string, data T) error {
	sagaCtx := &SagaContext[T]{
		ID:   id,
		Data: data,
		Step: -1,
	}

	err := o.repo.Save(ctx, o.saga.Name(), sagaCtx)
	if err != nil {
		return err
	}

	result := o.execute(ctx, sagaCtx)
	if result.err != nil {
		return err
	}

	return o.processResult(ctx, result)
}

func (o orchestrator[T]) ReplyTopic() string {
	return o.saga.ReplyTopic()
}

func (o orchestrator[T]) HandleReply(ctx context.Context, reply ddd.Reply) error {
	sagaID, sagaName := o.getSagaInfoFromReply(reply)
	if sagaID == "" || sagaName == "" || sagaName != o.saga.Name() {
		// returning nil to drop bad replies
		return nil
	}

	sagaCtx, err := o.repo.Load(ctx, o.saga.Name(), sagaID)
	if err != nil {
		return err
	}

	result, err := o.handle(ctx, sagaCtx, reply)
	if err != nil {
		return err
	}

	return o.processResult(ctx, result)
}

func (o orchestrator[T]) handle(ctx context.Context, sagaCtx *SagaContext[T], reply ddd.Reply) (stepResult[T], error) {
	step := o.saga.getSteps()[sagaCtx.Step]

	err := step.handle(ctx, sagaCtx, reply)
	if err != nil {
		return stepResult[T]{}, err
	}

	var success bool
	if outcome, ok := reply.Metadata().Get(am.ReplyOutcomeHdr).(string); !ok {
		success = false
	} else {
		success = outcome == am.OutcomeSuccess
	}

	switch {
	case success:
		return o.execute(ctx, sagaCtx), nil
	case sagaCtx.Compensating:
		return stepResult[T]{}, errors.ErrInternal.Msg("received failed reply but already compensating")
	default:
		sagaCtx.compensate()
		return o.execute(ctx, sagaCtx), nil
	}
}

func (o orchestrator[T]) execute(ctx context.Context, sagaCtx *SagaContext[T]) stepResult[T] {
	var delta = 1
	var direction = 1
	var step SagaStep[T]

	if sagaCtx.Compensating {
		direction = -1
	}

	steps := o.saga.getSteps()
	stepCount := len(steps)

	for i := sagaCtx.Step + direction; i > -1 && i < stepCount; i += direction {
		if step = steps[i]; step != nil && step.isInvocable(sagaCtx.Compensating) {
			break
		}
		delta += 1
	}

	if step == nil {
		sagaCtx.complete()
		return stepResult[T]{ctx: sagaCtx}
	}

	sagaCtx.advance(delta)

	return step.execute(ctx, sagaCtx)
}

func (o orchestrator[T]) processResult(ctx context.Context, result stepResult[T]) (err error) {
	if result.cmd != nil {
		err = o.publishCommand(ctx, result)
		if err != nil {
			return
		}
	}

	return o.repo.Save(ctx, o.saga.Name(), result.ctx)
}

func (o orchestrator[T]) publishCommand(ctx context.Context, result stepResult[T]) error {
	cmd := result.cmd

	cmd.Metadata().Set(am.CommandReplyChannelHdr, o.saga.ReplyTopic())
	cmd.Metadata().Set(SagaCommandIDHdr, result.ctx.ID)
	cmd.Metadata().Set(SagaCommandNameHdr, o.saga.Name())

	return o.publisher.Publish(ctx, cmd.Destination(), cmd)
}

func (o orchestrator[T]) getSagaInfoFromReply(reply ddd.Reply) (string, string) {
	var ok bool
	var sagaID, sagaName string

	if sagaID, ok = reply.Metadata().Get(SagaReplyIDHdr).(string); !ok {
		return "", ""
	}

	if sagaName, ok = reply.Metadata().Get(SagaReplyNameHdr).(string); !ok {
		return "", ""
	}

	return sagaID, sagaName
}
