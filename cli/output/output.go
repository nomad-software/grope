package output

// Line represents a matched line in a file.
type Line struct {
	Number string
	Text   string
}

// Match represents matched lines in a file.
type Match struct {
	File  string
	Lines []Line
}
