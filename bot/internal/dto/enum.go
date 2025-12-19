package dto

type EventType int32

const (
	EVENT_TYPE_UNSPECIFIED EventType = 0
	MESSAGE_DELETE         EventType = 1
	MESSAGE_EDIT           EventType = 2
	USER_JOIN              EventType = 3
	USER_LEAVE             EventType = 4
	CREATE_INVITE          EventType = 5
	JOIN_CHANNEL           EventType = 6
	LEAVE_CHANNEL          EventType = 7
	MOVE_CHANNEL           EventType = 8
)

func (e EventType) String() string {
	switch e {
	case MESSAGE_DELETE:
		return "MESSAGE_DELETE"
	case MESSAGE_EDIT:
		return "MESSAGE_EDIT"
	case USER_JOIN:
		return "USER_JOIN"
	case USER_LEAVE:
		return "USER_LEAVE"
	case CREATE_INVITE:
		return "CREATE_INVITE"
	case JOIN_CHANNEL:
		return "JOIN_CHANNEL"
	case LEAVE_CHANNEL:
		return "LEAVE_CHANNEL"
	case MOVE_CHANNEL:
		return "MOVE_CHANNEL"
	default:
		return "UNKNOWN"
	}
}