package usecase

import (
	"gosdk"
	"net/http"
	"net/url"
	"strconv"
)

// Disable unused import error
var _ = strconv.IntSize
var _ url.Error
var _ = http.NoBody

type Queue struct {
	restClient gosdk.RestClientInterface
}

func NewQueue(restClient gosdk.RestClientInterface) *Queue {
	return &Queue{restClient}
}

// GET: Queries
func (t *Queue) GetAllJobs() (r *http.Response, err error) {
	queryParameters := url.Values{}

	return t.restClient.Get(
		`/v2/Queue/UseCase/GetAllJobs`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string JobId
func (t *Queue) GetJob(JobId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`jobId`, JobId)

	return t.restClient.Get(
		`/v2/Queue/UseCase/GetJob`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string CommandId
// @param array|null WithData QueueCommandErrors|QueueCommandMessages
func (t *Queue) GetQueueCommand(CommandId string, WithData *[]string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`commandId`, CommandId)
	if WithData != nil {
		for i := range *WithData {
			queryParameters.Add(`withData[]`, (*WithData)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Queue/UseCase/GetQueueCommand`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string QueueTaskId
// @param array|null WithData QueueTaskErrors|QueueTaskMessages
func (t *Queue) GetQueueTask(QueueTaskId string, WithData *[]string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`queueTaskId`, QueueTaskId)
	if WithData != nil {
		for i := range *WithData {
			queryParameters.Add(`withData[]`, (*WithData)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Queue/UseCase/GetQueueTask`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserId
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-100
// @param bool|null IsFinished true|false
// @param bool|null IsSuccess true|false
func (t *Queue) ListQueueCommandsForUser(UserId string, Page *int, ItemsPerPage *int, IsFinished *bool, IsSuccess *bool) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, UserId)
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
	}
	if IsFinished != nil {
		queryParameters.Add(`isFinished`, strconv.FormatBool(*IsFinished))
	}
	if IsSuccess != nil {
		queryParameters.Add(`isSuccess`, strconv.FormatBool(*IsSuccess))
	}

	return t.restClient.Get(
		`/v2/Queue/UseCase/ListQueueCommandsForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string JobId
func (t *Queue) DeleteJob(JobId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`jobId`, JobId)

	return t.restClient.Post(
		`/v2/Queue/UseCase/DeleteJob`,
		&queryParameters,
		nil,
		nil,
	)
}
