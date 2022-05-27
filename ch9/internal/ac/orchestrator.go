package ac

import (
	"context"

	"github.com/stackus/errors"

	"eda-in-golang/ch9/internal/am"
	"eda-in-golang/ch9/internal/ddd"
)

type (
	Orchestrator[T any] interface {
		Start(ctx context.Context, id string, data T) error
		ReplyTopic() string
		HandleMessage(ctx context.Context, message am.IncomingReplyMessage) error
	}

	orchestrator[T any] struct {
		saga      Saga[T]
		repo      SagaRepository[T]
		publisher am.CommandPublisher
	}
)

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

	return o.processStep(ctx, result)
}

func (o orchestrator[T]) ReplyTopic() string {
	return o.saga.ReplyTopic()
}

func (o orchestrator[T]) HandleMessage(ctx context.Context, replyMsg am.IncomingReplyMessage) error {
	sagaID, sagaName := o.replySagaInfo(replyMsg)
	if sagaID == "" || sagaName == "" || sagaName != o.saga.Name() {
		// returning nil to drop bad replies
		return nil
	}

	sagaCtx, err := o.repo.Load(ctx, o.saga.Name(), sagaID)
	if err != nil {
		return err
	}

	result, err := o.handle(ctx, sagaCtx, replyMsg)
	if err != nil {
		return err
	}

	return o.processStep(ctx, result)
}

func (o orchestrator[T]) handle(ctx context.Context, sagaCtx *SagaContext[T], replyMsg am.ReplyMessage) (stepResult[T], error) {
	step := o.saga.Steps()[sagaCtx.Step]

	err := step.handle(ctx, sagaCtx, replyMsg)
	if err != nil {
		return stepResult[T]{}, err
	}

	var success bool
	if outcome, ok := replyMsg.Metadata().Get(am.ReplyOutcomeHdr).(string); !ok {
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

	steps := o.saga.Steps()
	stepCount := len(steps)

	for i := sagaCtx.Step + direction; i > -1 && i < stepCount; i += direction {
		if step := steps[i]; step != nil && step.isInvocable(sagaCtx.Compensating) {
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

func (o orchestrator[T]) processStep(ctx context.Context, result stepResult[T]) (err error) {
	if result.cmd != nil {
		err = o.publishCommand(ctx, result)
		if err != nil {
			return
		}
	}

	return o.repo.Save(ctx, o.saga.Name(), result.ctx)
}

func (o orchestrator[T]) publishCommand(ctx context.Context, result stepResult[T]) error {
	command := result.cmd

	command.Metadata().Set(am.CommandReplyChannelHdr, o.saga.ReplyTopic())
	command.Metadata().Set(SagaCommandIDHdr, result.ctx.ID)
	command.Metadata().Set(SagaCommandNameHdr, o.saga.Name())

	return o.publisher.Publish(ctx, command.Destination(), command)
}

func (o orchestrator[T]) replySagaInfo(reply ddd.Reply) (string, string) {
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
