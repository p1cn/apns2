package apns2

import "github.com/ikenchina/apns2/payload"

// Your total notification payload cannot exceed 256 bytes for IOS7 and earlier, 4kb IOS8 and later.
const MAX_PAYLOAD_SIZE_BEFORE_IOS8_BYTES = 256
const MAX_PAYLOAD_SIZE_BYTES = 4096

type PushNotification struct {
	Notification
	maxPayloadSize int
}

// Constructor. Also initializes the pseudo-random identifier.
func NewPushNotification(isIOS7OrEarlier bool) (pn *PushNotification) {
	pn = new(PushNotification)
	pn.Payload = payload.NewPayload()
	pn.maxPayloadSize = MAX_PAYLOAD_SIZE_BYTES
	if isIOS7OrEarlier {
		pn.maxPayloadSize = MAX_PAYLOAD_SIZE_BEFORE_IOS8_BYTES
	}
	pn.Priority = PriorityHigh
	return
}

func (this *PushNotification) ExceededMaxPayload() (exceeded bool, extraLength int, err error) {
	exceeded = false
	payload, err := this.MarshalJSON()
	if err != nil {
		return
	}
	length := len(payload)
	if length > this.maxPayloadSize {
		exceeded = true
		extraLength = length - this.maxPayloadSize
	}
	return
}

func (this *PushNotification) GetPayload() *payload.Payload {
	return this.Payload.(*payload.Payload)
}
