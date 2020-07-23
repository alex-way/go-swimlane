package swimlane

import (
	"errors"
	"time"

	"github.com/go-resty/resty/v2"
)

type Workspace struct {
	Type           string         `json:"$type"`
	CreatedDate    time.Time      `json:"createdDate"`
	CreatedByUser  CreatedByUser  `json:"createdByUser"`
	ModifiedDate   time.Time      `json:"modifiedDate"`
	ModifiedByUser ModifiedByUser `json:"modifiedByUser"`
	Dashboards     []string       `json:"dashboards"`
	Applications   []string       `json:"applications"`
	Permissions    Permissions    `json:"permissions"`
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	Disabled       bool           `json:"disabled"`
}

type CreatedByUser struct {
	Type string `json:"$type"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ModifiedByUser struct {
	Type string `json:"$type"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Permissions struct {
	Type string `json:"$type,omitempty"`
}

// ListWorkspaces returns a list of Workspace structs of all workspaces
func (c *Client) ListWorkspaces() ([]*Workspace, *resty.Response, error) {
	url := c.parseURL("/api/workspaces")
	var workspaces []*Workspace
	resp, err := c.R().SetResult(&workspaces).Get(url)
	if err != nil {
		return []*Workspace{}, resp, err
	}
	if resp.StatusCode() != 200 {
		err := "The get request was unsuccessful"
		return []*Workspace{}, resp, errors.New(err)
	}
	return workspaces, resp, nil
}

func (c *Client) GetWorkspace(ID string) (*Workspace, *resty.Response, error) {
	url := c.parseURL("/api/workspaces/" + ID)
	var workspace *Workspace
	resp, err := c.R().SetResult(&workspace).Get(url)
	if err != nil {
		return workspace, resp, err
	}
	if resp.StatusCode() != 200 {
		err := "The get request was unsuccessful"
		return workspace, resp, errors.New(err)
	}
	return workspace, resp, nil
}

type CreateWorkspaceOptions struct {
	Name         string      `json:"name"`
	Applications []string    `json:"applications"`
	Dashboards   []string    `json:"dashboards"`
	Permissions  Permissions `json:"permissions,omitempty"`
}

func (c *Client) CreateWorkspace(body *CreateWorkspaceOptions) (*Workspace, *resty.Response, error) {
	url := c.parseURL("/api/workspaces")
	var createdWorkspace *Workspace
	resp, err := c.R().SetResult(&createdWorkspace).SetBody(body).Post(url)
	if err != nil {
		return createdWorkspace, resp, err
	}
	if resp.StatusCode() != 200 {
		err := "The post request was unsuccessful."
		return createdWorkspace, resp, errors.New(err)
	}
	return createdWorkspace, resp, nil
}

func (c *Client) DeleteWorkspace(ID string) (*resty.Response, error) {
	url := c.parseURL("/api/workspaces/" + ID)
	resp, err := c.R().Delete(url)
	if err != nil {
		return resp, err
	}
	if resp.IsError() {
		err := "The delete request was unsuccessful."
		return resp, errors.New(err)
	}
	return resp, nil
}
