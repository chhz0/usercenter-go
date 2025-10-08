package options

type LogOptions struct {
	Level      string
	Encoding   string
	Caller     bool
	CallerSkip int
	Output     []string
}
