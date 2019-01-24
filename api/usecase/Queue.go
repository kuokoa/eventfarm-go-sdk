/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"fmt"
	"github.com/eventfarm/go-sdk/rest"
	"net/http"
	"net/url"
	"strconv"
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

// @param string JobId

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

// @param string CommandId
// @param array|null WithData QueueCommandErrors|QueueCommandMessages

type GetQueueCommandParameters struct {
	CommandId string
	WithData  *[]string
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

// @param string QueueTaskId
// @param array|null WithData QueueTaskErrors|QueueTaskMessages

type GetQueueTaskParameters struct {
	QueueTaskId string
	WithData    *[]string
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

// @param string UserId
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-100
// @param bool|null IsFinished true|false
// @param bool|null IsSuccess true|false

type ListQueueCommandsForUserParameters struct {
	UserId       string
	Page         *int64
	ItemsPerPage *int64
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
// @param string JobId

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
