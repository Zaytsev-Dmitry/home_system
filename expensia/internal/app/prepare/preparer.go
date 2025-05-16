package prepare

type Preparer[In any, Out any] interface {
	Prepare(input In) (Out, error)
}
