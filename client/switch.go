package client

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// idsFlag contains the ids of the reminders
type idsFlag []string

// String return a string with all the reminder's id separated by ','
func (list idsFlag) String() string {
	return strings.Join(list, ",")
}

// Set adds an id to the list of ids
func (list *idsFlag) Set(v string) error {
	*list = append(*list, v)
	return nil
}

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
		createCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		t, m, d := s.reminderFlags(createCmd)

		if err := s.checkArgs(3); err != nil {
			return err
		}
		if err := s.parseCmd(createCmd); err != nil {
			return err
		}

		res, err := s.client.Create(*t, *m, *d)
		if err != nil {
			return wrapError("could not create reminder", err)
		}
		fmt.Printf("Reminder created succesfully:\n%s", string(res))
		return nil
	}
}

// edit modify an existing reminder
func (s Switch) edit() func(string) error {
	return func(cmd string) error {
		ids := idsFlag{}
		editCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		editCmd.Var(&ids, "id", "The ID (int) of the reminder to edit")
		t, m, d := s.reminderFlags(editCmd)

		if err := s.checkArgs(2); err != nil {
			return err
		}
		if err := s.parseCmd(editCmd); err != nil {
			return err
		}

		lastID := ids[len(ids)-1]
		res, err := s.client.Edit(lastID, *t, *m, *d)
		if err != nil {
			return wrapError("Could not edit reminder", err)
		}
		fmt.Printf("Reminder edited succesfully:\n%s", string(res))
		return nil
	}
}

// fetch gets all the reminders
func (s Switch) fetch() func(string) error {
	return func(cmd string) error {
		ids := idsFlag{}
		fetchCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		fetchCmd.Var(&ids, "id", "The ID (int) of the reminder to fetch")

		if err := s.checkArgs(1); err != nil {
			return err
		}
		if err := s.parseCmd(fetchCmd); err != nil {
			return err
		}

		res, err := s.client.Fetch(ids)
		if err != nil {
			return wrapError("could not fetch reminder(s)", err)
		}
		fmt.Printf("reminders fetched succesfully:\n%s", string(res))
		return nil
	}
}

// delete deletes a reminder
func (s Switch) delete() func(string) error {
	return func(cmd string) error {
		ids := idsFlag{}
		deleteCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		deleteCmd.Var(&ids, "id", "List of reminder IDs (int) to delete")

		if err := s.checkArgs(1); err != nil {
			return err
		}
		if err := s.parseCmd(deleteCmd); err != nil {
			return err
		}

		err := s.client.Delete(ids)
		if err != nil {
			return wrapError("could not delete reminder(s)", err)
		}
		fmt.Printf("succesfully deleted record(s):\n%v\n", ids)
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

// reminderFlags sets the title, message and duration of a reminder on a FlagSet
func (s Switch) reminderFlags(f *flag.FlagSet) (*string, *string, *time.Duration) {
	t, m, d := "", "", time.Duration(0)
	f.StringVar(&t, "title", "", "Reminder title")
	f.StringVar(&t, "t", "", "Reminder title")
	f.StringVar(&m, "message", "", "Reminder message")
	f.StringVar(&m, "m", "", "Reminder message")
	f.StringVar(&m, "duration", "", "Reminder time")
	f.StringVar(&m, "d", "", "Reminder time")

	return &t, &m, &d
}

// checkArgs checks if there are enough arguments on os.Args
func (s Switch) checkArgs(minArgs int) error {
	if len(os.Args) == 3 && os.Args[2] == "--help" {
		return nil
	}
	if len(os.Args)-2 < minArgs {
		fmt.Printf(
			"incorrect use of %s\n%s %s --help",
			os.Args[1], os.Args[0], os.Args[1],
		)
		return fmt.Errorf(
			"%s expects at least %d arg(s), %d provided",
			os.Args[1], minArgs, len(os.Args)-2,
		)
	}
	return nil
}

// parseCmd parses the arguments from os.Args[2:]
func (s Switch) parseCmd(cmd *flag.FlagSet) error {
	err := cmd.Parse(os.Args[2:])
	if err != nil {
		return wrapError("could not parse '"+cmd.Name()+"' flags", err)
	}
	return nil
}
