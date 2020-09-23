package builder

// FakeCommand implements the Command interface and is used for mocking on testing scenarios
type FakeCommand struct {
	ArgsAppend        []string
	ArgsInteractive   []string
	ArgsExec          []string
	CalledAppendArgs  bool
	CalledString      bool
	CalledLookPath    bool
	CalledInteractive bool
	CalledExec        bool
}

// AppendArgs mocked function for testing
func (f *FakeCommand) AppendArgs(args ...string) {
	f.ArgsAppend = append(f.ArgsAppend, args...)
	f.CalledAppendArgs = true
}

// String mocked function for testing
func (f *FakeCommand) String() string {
	f.CalledString = true
	return ""
}

// LookPath returns if the command exists
func (f *FakeCommand) LookPath() (err error) {
	f.CalledLookPath = true
	return
}

// Interactive will send the command to an interactive execution.
func (f *FakeCommand) Interactive(args ...string) (err error) {
	f.CalledInteractive = true
	f.ArgsInteractive = args
	return
}

// Exec will send the command to shell execution.
func (f *FakeCommand) Exec(args ...string) (outStr string, err error) {
	f.CalledExec = true
	f.ArgsExec = args
	return
}
