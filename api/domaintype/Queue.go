/**
 *  This file was auto generated, please do not edit it directly.
**/

package domaintype

type Queue struct {
}

func NewQueue() *Queue {
	return &Queue{}
}

type QueueCommandMessageLevelType struct {
	Slug        string
	Name        string
	Description string
	IsError     bool
	IsWarning   bool
	IsInfo      bool
}

type QueueTaskMessageLevelType struct {
	Slug        string
	Name        string
	Description string
	IsError     bool
	IsWarning   bool
	IsInfo      bool
}

func (f *Queue) ListQueueCommandMessageLevelTypes() []QueueCommandMessageLevelType {
	return []QueueCommandMessageLevelType{
		{
			Slug:        `error`,
			Name:        `Error`,
			Description: ``,
			IsError:     true,
			IsWarning:   false,
			IsInfo:      false,
		},
		{
			Slug:        `warning`,
			Name:        `Warning`,
			Description: ``,
			IsError:     false,
			IsWarning:   true,
			IsInfo:      false,
		},
		{
			Slug:        `info`,
			Name:        `Info`,
			Description: ``,
			IsError:     false,
			IsWarning:   false,
			IsInfo:      true,
		},
	}
}

func (f *Queue) ListQueueTaskMessageLevelTypes() []QueueTaskMessageLevelType {
	return []QueueTaskMessageLevelType{
		{
			Slug:        `error`,
			Name:        `Error`,
			Description: ``,
			IsError:     true,
			IsWarning:   false,
			IsInfo:      false,
		},
		{
			Slug:        `warning`,
			Name:        `Warning`,
			Description: ``,
			IsError:     false,
			IsWarning:   true,
			IsInfo:      false,
		},
		{
			Slug:        `info`,
			Name:        `Info`,
			Description: ``,
			IsError:     false,
			IsWarning:   false,
			IsInfo:      true,
		},
	}
}
