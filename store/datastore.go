// Copyright 2019 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package store

import (
	"context"
	"errors"

	"github.com/mendersoftware/workflows/model"
)

var (
	ErrWorkflowNotFound      = errors.New("Workflow not found")
	ErrWorkflowMissingName   = errors.New("Workflow missing name")
	ErrWorkflowAlreadyExists = errors.New("Workflow already exists")
)

// DataStoreMongoInterface for DataStore  services
type DataStore interface {
	InsertWorkflows(workflow ...model.Workflow) (int, error)
	GetWorkflowByName(workflowName string) (*model.Workflow, error)
	GetWorkflows() []model.Workflow
	InsertJob(ctx context.Context, job *model.Job) (*model.Job, error)
	GetJobs(ctx context.Context) <-chan *model.Job
	AquireJob(ctx context.Context, job *model.Job) (*model.Job, error)
	UpdateJobAddResult(ctx context.Context, job *model.Job, result *model.TaskResult) error
	UpdateJobStatus(ctx context.Context, job *model.Job, status int) error
	GetJobByNameAndID(ctx context.Context, name string, ID string) (*model.Job, error)
	Shutdown()
}
