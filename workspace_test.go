package swimlane

import (
	"testing"
)

func TestListWorkspaces(t *testing.T) {
	client, _ := NewTestClient()
	_, _, err := client.ListWorkspaces()
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateGetDeleteWorkspace(t *testing.T) {
	client, _ := NewTestClient()
	workspaceOptions := CreateWorkspaceOptions{
		Name: "deleteme",
	}
	createdWorkspace, _, err := client.CreateWorkspace(&workspaceOptions)
	if err != nil {
		t.Fatal(err)
	}

	_, _, err = client.GetWorkspace(createdWorkspace.ID)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.DeleteWorkspace(createdWorkspace.ID)
	if err != nil {
		t.Fatal(err)
	}
}
