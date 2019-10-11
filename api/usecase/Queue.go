/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/kuokoa/eventfarm-go-sdk/rest"
)

type Queue struct {
	restClient rest.RestClientInterface
}

func NewQueue(restClient rest.RestClientInterface) *Queue {
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

type GetJobParameters struct {
	JobId string
}

func (t *Queue) GetJob(p *GetJobParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`jobId`, p.JobId)

	return t.restClient.Get(
		`/v2/Queue/UseCase/GetJob`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetQueueCommandParameters struct {
	CommandId string
	WithData  *[]string // QueueCommandErrors | QueueCommandMessages
}

func (t *Queue) GetQueueCommand(p *GetQueueCommandParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`commandId`, p.CommandId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Queue/UseCase/GetQueueCommand`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetQueueTaskParameters struct {
	QueueTaskId string
	WithData    *[]string // QueueTaskErrors | QueueTaskMessages
}

func (t *Queue) GetQueueTask(p *GetQueueTaskParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`queueTaskId`, p.QueueTaskId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Queue/UseCase/GetQueueTask`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListQueueCommandsForUserParameters struct {
	UserId       string
	Page         *int64 // >= 1
	ItemsPerPage *int64 // 1-100
	IsFinished   *bool
	IsSuccess    *bool
}

func (t *Queue) ListQueueCommandsForUser(p *ListQueueCommandsForUserParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}
	if p.IsFinished != nil {
		queryParameters.Add(`isFinished`, strconv.FormatBool(*p.IsFinished))
	}
	if p.IsSuccess != nil {
		queryParameters.Add(`isSuccess`, strconv.FormatBool(*p.IsSuccess))
	}

	return t.restClient.Get(
		`/v2/Queue/UseCase/ListQueueCommandsForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type DeleteJobParameters struct {
	JobId string
}

func (t *Queue) DeleteJob(p *DeleteJobParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`jobId`, p.JobId)

	return t.restClient.Post(
		`/v2/Queue/UseCase/DeleteJob`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Queue) DeleteJobWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Queue/UseCase/DeleteJob`,
		data,
		nil,
		nil,
	)
}
