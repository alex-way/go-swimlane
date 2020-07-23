# Swimlane Go SDK

A Swimlane API client enabling Go programs to interact with Swimlane in a simple and uniform way

## Usage

```go
import "github.com/alex-way/go-swimlane"
```

Construct a new Swimlane client, then use the various services on the client to access different objects within Swimlane. For example to list all workspaces:

```go
swimClient, err := swimlane.NewClient("https://myswimlaneinstance.com", "myaccesstoken")
if err != nil {
  log.Fatalf("Failed to create client: %v", err)
}
workspaces, _, err := swimClient.ListWorkspaces()
```
