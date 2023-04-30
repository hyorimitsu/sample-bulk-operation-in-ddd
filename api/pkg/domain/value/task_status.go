package value

import (
	"fmt"
)

type TaskStatus int

const (
	TaskStatusUndefined  TaskStatus = -1
	TaskStatusNone       TaskStatus = 0
	TaskStatusInProgress TaskStatus = 10
	TaskStatusDone       TaskStatus = 20
	TaskStatusWaiting    TaskStatus = 30
	TaskStatusPending    TaskStatus = 99

	TaskStatusUndefinedString  string = "undefined"
	TaskStatusNoneString       string = "none"
	TaskStatusInProgressString string = "in_progress"
	TaskStatusDoneString       string = "done"
	TaskStatusWaitingString    string = "waiting"
	TaskStatusPendingString    string = "pending"
)

func (v TaskStatus) Value() int {
	return int(v)
}

func (v TaskStatus) String() string {
	switch v {
	case TaskStatusNone:
		return TaskStatusNoneString
	case TaskStatusInProgress:
		return TaskStatusInProgressString
	case TaskStatusDone:
		return TaskStatusDoneString
	case TaskStatusWaiting:
		return TaskStatusWaitingString
	case TaskStatusPending:
		return TaskStatusPendingString
	default:
		return TaskStatusUndefinedString
	}
}

func TaskStatusFromString(value string) (TaskStatus, error) {
	switch value {
	case TaskStatusNoneString:
		return TaskStatusNone, nil
	case TaskStatusInProgressString:
		return TaskStatusInProgress, nil
	case TaskStatusDoneString:
		return TaskStatusDone, nil
	case TaskStatusWaitingString:
		return TaskStatusWaiting, nil
	case TaskStatusPendingString:
		return TaskStatusPending, nil
	default:
		return TaskStatusUndefined, fmt.Errorf("invalid task status [value: %s]", value)
	}
}

func TaskStatusFromInt(value int) (TaskStatus, error) {
	switch value {
	case TaskStatusNone.Value():
		return TaskStatusNone, nil
	case TaskStatusInProgress.Value():
		return TaskStatusInProgress, nil
	case TaskStatusDone.Value():
		return TaskStatusDone, nil
	case TaskStatusWaiting.Value():
		return TaskStatusWaiting, nil
	case TaskStatusPending.Value():
		return TaskStatusPending, nil
	default:
		return TaskStatusUndefined, fmt.Errorf("invalid task status [value: %d]", value)
	}
}
