package internal

var Arg = &Args{}

type Args struct {
	File      string `clop:"-f; --file" usage:"file name"`
	Directory string `clop:"-d; --dir" usage:"directory name"`
	Question  string `clop "-q; --question" usage:"question"`
	Category  string `clop:"-c; --category" usage:"category"`
}
