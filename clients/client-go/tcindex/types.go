// This source code file is AUTO-GENERATED by github.com/taskcluster/jsonschema2go

package tcindex

import (
	"encoding/json"

	tcclient "github.com/taskcluster/taskcluster/clients/client-go/v17"
)

type (
	// Representation of an indexed task.
	IndexedTaskResponse struct {

		// Data that was reported with the task. This is an arbitrary JSON object.
		//
		// Additional properties allowed
		Data json.RawMessage `json:"data"`

		// Date at which this entry expires from the task index.
		Expires tcclient.Time `json:"expires"`

		// Namespace of the indexed task, used to find the indexed task in the index.
		//
		// Max length: 255
		Namespace string `json:"namespace"`

		// If multiple tasks are indexed with the same `namespace` the task with the
		// highest `rank` will be stored and returned in later requests. If two tasks
		// has the same `rank` the latest task will be stored.
		Rank float64 `json:"rank"`

		// Unique task identifier, this is UUID encoded as
		// [URL-safe base64](http://tools.ietf.org/html/rfc4648#section-5) and
		// stripped of `=` padding.
		//
		// Syntax:     ^[A-Za-z0-9_-]{8}[Q-T][A-Za-z0-9_-][CGKOSWaeimquy26-][A-Za-z0-9_-]{10}[AQgw]$
		TaskID string `json:"taskId"`
	}

	// Representation of the index entry to insert.
	InsertTaskRequest struct {

		// This is an arbitrary JSON object. Feel free to put whatever data you want
		// here, but do limit it, you'll get errors if you store more than 32KB.
		// So stay well, below that limit.
		//
		// Additional properties allowed
		Data json.RawMessage `json:"data"`

		// Date at which this entry expires from the task index.
		Expires tcclient.Time `json:"expires"`

		// If multiple tasks are indexed with the same `namespace` the task with the
		// highest `rank` will be stored and returned in later requests. If two tasks
		// has the same `rank` the latest task will be stored.
		Rank float64 `json:"rank"`

		// Unique task identifier, this is UUID encoded as
		// [URL-safe base64](http://tools.ietf.org/html/rfc4648#section-5) and
		// stripped of `=` padding.
		//
		// Syntax:     ^[A-Za-z0-9_-]{8}[Q-T][A-Za-z0-9_-][CGKOSWaeimquy26-][A-Za-z0-9_-]{10}[AQgw]$
		TaskID string `json:"taskId"`
	}

	// Response from a request to list namespaces within a given namespace.
	ListNamespacesResponse struct {

		// A continuation token is returned if there are more results than listed
		// here. You can optionally provide the token in the request payload to
		// load the additional results.
		ContinuationToken string `json:"continuationToken,omitempty"`

		// List of namespaces.
		Namespaces []Namespace `json:"namespaces"`
	}

	// Representation of an indexed task.
	ListTasksResponse struct {

		// A continuation token is returned if there are more results than listed
		// here. You can optionally provide the token in the request payload to
		// load the additional results.
		ContinuationToken string `json:"continuationToken,omitempty"`

		// List of tasks.
		Tasks []Task `json:"tasks"`
	}

	// Representation of a namespace that contains indexed tasks.
	Namespace struct {

		// Date at which this entry, and by implication all entries below it,
		// expires from the task index.
		Expires tcclient.Time `json:"expires"`

		// Name of namespace within it's parent namespace.
		Name string `json:"name"`

		// Fully qualified name of the namespace, you can use this to list
		// namespaces or tasks under this namespace.
		//
		// Max length: 255
		Namespace string `json:"namespace"`
	}

	// Representation of a task.
	Task struct {

		// Data that was reported with the task. This is an arbitrary JSON
		// object.
		//
		// Additional properties allowed
		Data json.RawMessage `json:"data"`

		// Date at which this entry expires from the task index.
		Expires tcclient.Time `json:"expires"`

		// Index path of the task.
		//
		// Max length: 255
		Namespace string `json:"namespace"`

		// If multiple tasks are indexed with the same `namespace` the task
		// with the highest `rank` will be stored and returned in later
		// requests. If two tasks has the same `rank` the latest task will be
		// stored.
		Rank float64 `json:"rank"`

		// Unique task identifier for the task currently indexed at `namespace`.
		//
		// Syntax:     ^[A-Za-z0-9_-]{8}[Q-T][A-Za-z0-9_-][CGKOSWaeimquy26-][A-Za-z0-9_-]{10}[AQgw]$
		TaskID string `json:"taskId"`
	}
)
