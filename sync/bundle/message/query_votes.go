package message

import (
	"fmt"

	"github.com/zarbchain/zarb-go/util/errors"
)

type QueryVotesMessage struct {
	Height uint32 `cbor:"1,keyasint"`
	Round  int16  `cbor:"2,keyasint"`
}

func NewQueryVotesMessage(h uint32, r int16) *QueryVotesMessage {
	return &QueryVotesMessage{
		Height: h,
		Round:  r,
	}
}

func (m *QueryVotesMessage) SanityCheck() error {
	if m.Round < 0 {
		return errors.Error(errors.ErrInvalidRound)
	}
	return nil
}

func (m *QueryVotesMessage) Type() Type {
	return MessageTypeQueryVotes
}

func (m *QueryVotesMessage) Fingerprint() string {
	return fmt.Sprintf("{%d/%d}", m.Height, m.Round)
}
