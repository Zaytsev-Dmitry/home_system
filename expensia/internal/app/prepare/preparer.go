package prepare

type Preparer interface {
	Prepare(input interface{}) (interface{}, error)
}
