package generate

const (
	FileSyntax  = "syntax"
	FilePackage = "package"
	FileImport  = "import"
)

func NewProto() (proto Proto) {
	return Proto{}
}

// FileState
type Proto struct {
	Syntax   string
	Package  string
	Imports  []string
	Options  []Option
	Services []Service
	Messages []Message
}

type Option struct {
	Name  string
	Value string
}

type Service struct {
	Name    string
	Options []Option
	RPCs    []RPC
}

type RPC struct {
	Name    string
	Input   string
	Output  string
	Options []Option
}

type Message struct {
	Name   string
	Fields []Field
}

type Field struct {
	Name     string
	DataType string
	Position string
	Options  []Option
}

type KV struct {
	Key string
	Val string
}
