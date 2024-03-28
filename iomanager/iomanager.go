package iomanager

type IOManager interface {
	ReadLinesFromFile() ([]string, error)
	WriteToJson(data interface{}) error
}