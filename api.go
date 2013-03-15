package tbreadline

type Terminal struct {
	terminal
}

var (
	HistoryLength = -1
)

func (T *Terminal) Close() {
	T.close()
}

func (T *Terminal) Reset() error {
	return T.reset()
}

func (T *Terminal) ReadLine() *string {
	line := "hello"
	return &line
}

func AddHistory(s string) error {
	return nil
}

func ReadHistoryFile(s string) error {
	return nil
}

func WriteHistoryFile(s string) error {
	return nil
}

func SetCompleterDelims(break_chars string) {
}

func CompletionMatches(text string, cbk func(text string, state int) string) []string {
	match := []string{"ls", "cat"}
	return match
}

func SetAttemptedCompletionFunction(cbk func(text string, start, end int) string) {
}
