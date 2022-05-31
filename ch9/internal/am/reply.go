package am

// import (
// 	"strings"
//
// 	"eda-in-golang/internal/ddd"
// )
//
const (
	FailureReply = "am.Failure"
	SuccessReply = "am.Success"

	OutcomeSuccess = "SUCCESS"
	OutcomeFailure = "FAILURE"

	ReplyHdrPrefix  = "REPLY_"
	ReplyNameHdr    = ReplyHdrPrefix + "NAME"
	ReplyOutcomeHdr = ReplyHdrPrefix + "OUTCOME"
)

//
// type (
// 	Reply interface {
// 		ddd.Reply
// 		Destination() string
// 	}
//
// 	reply struct {
// 		ddd.Reply
// 		destination string
// 	}
// )
//
// func NewReply(name string, payload ddd.ReplyPayload, cmd ddd.Command, options ...ddd.ReplyOption) Reply {
// 	correlationHeaders := make(ddd.Metadata)
//
// 	for key, value := range cmd.Metadata() {
// 		if key == CommandNameHdr {
// 			continue
// 		}
//
// 		if strings.HasPrefix(key, CommandHdrPrefix) {
// 			hdr := ReplyHdrPrefix + key[len(CommandHdrPrefix):]
// 			correlationHeaders[hdr] = value
// 		}
// 	}
//
// 	destination := cmd.Metadata().Get(CommandReplyChannelHdr).(string)
//
// 	return reply{
// 		Reply:       ddd.NewReply(name, payload, append(options, correlationHeaders)...),
// 		destination: destination,
// 	}
// }
//
// func (r reply) Destination() string {
// 	return r.destination
// }
