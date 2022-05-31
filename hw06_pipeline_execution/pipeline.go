package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	resultChan := stageDone(done, in)
	for _, stage := range stages {
		resultChan = stageDone(done, stage(resultChan))
	}

	return resultChan
}

func stageDone(done In, resultChan In) Out {
	stageChan := make(Bi)
	go func(inChan In, outChan Bi) {
		defer close(outChan)
		for {
			select {
			case <-done:
				return
			case result, ok := <-inChan:
				if !ok {
					return
				}
				outChan <- result
			}
		}
	}(resultChan, stageChan)

	return stageChan
}
