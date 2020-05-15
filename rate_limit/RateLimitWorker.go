package rate_limit

type RateLimitWorker interface {
	// total number of work to be done
	WorkSize() int

	// get input for index
	GetInput(index int) (interface{}, error)

	// do work for one input
	Work(interface{}) (interface{}, error)

	//process resp
	ProcessResp(resp interface{}, output interface{})

	//Handle the output
	HandleOutput()

	//post Handling the output
	Close()
}
