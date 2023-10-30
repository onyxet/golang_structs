package file_reader

type FileReader interface {
	ReadFile(string) error
	Search() ([]string, error)
}
