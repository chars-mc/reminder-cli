package client

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// BackendHTTPClient defines the methods of the http client
type BackendHTTPClient interface {
	Create(title, message string, duration time.Duration) ([]byte, error)
	Edit(id, title, message string, duration time.Duration) ([]byte, error)
	Fetch(ids []string) ([]byte, error)
	Delete(ids []string) error
	Healthy(host string) bool
}

// Switch contains the http clients, the backend url
// and a map with the commands and functions
type Switch struct {
	client        BackendHTTPClient
	backendAPIURL string
	commands      map[string]func() func(string) error
}

// NewSwitch returns a new Switch
func NewSwitch(uri string) Switch {
	httpClient := NewHTTPClient(uri)
	s := Switch{
		client:        httpClient,
		backendAPIURL: uri,
	}
	s.commands = map[string]func() func(string) error{
		"create": s.create,
		"edit":   s.edit,
		"fetch":  s.fetch,
		"delete": s.delete,
		"health": s.health,
	}
	return s
}

// Switch executes a function if a command exists, if it doesn't returns an error
func (s Switch) Switch() error {
	cmdName := os.Args[1]
	cmd, ok := s.commands[cmdName]
	if !ok {
		return fmt.Errorf("invalid command '%s'", cmdName)
	}
	return cmd()(cmdName)
}

// create adds a new reminder
func (s Switch) create() func(string) error {
	return func(cmd string) error {
		fmt.Println("create reminder")
		return nil
	}
}

// edit modify an existing reminder
func (s Switch) edit() func(string) error {
	return func(cmd string) error {
		fmt.Println("edit reminder")
		return nil
	}
}

// fetch gets all the reminders
func (s Switch) fetch() func(string) error {
	return func(cmd string) error {
		fmt.Println("fetch reminders")
		return nil
	}
}

// delete deletes a reminder
func (s Switch) delete() func(string) error {
	return func(cmd string) error {
		fmt.Println("delete reminder")
		return nil
	}
}

// health checks if the notifier server is running
func (s Switch) health() func(string) error {
	return func(cmd string) error {
		fmt.Println("calling health")
		return nil
	}
}

// Help prints a help message for commands
func (s Switch) Help() {
	sb := strings.Builder{}
	for name := range s.commands {
		sb.WriteString(name)
		sb.WriteString("\t--help\n")
	}
	fmt.Printf("Usage of: %s:\n<command> [<args>]\n%s", os.Args[0], sb.String())
}
