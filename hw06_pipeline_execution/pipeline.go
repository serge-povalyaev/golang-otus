package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	resultChan := in
	for _, stage := range stages {
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

		resultChan = stage(stageChan)
	}

	return resultChan
}
