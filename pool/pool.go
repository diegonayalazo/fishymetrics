package pool

import "sync"

// Pool is a worker group that runs a number of tasks at a
// configured concurrency.
type Pool struct {
	Tasks []*Task

	concurrency int
	tasksChan   chan *Task
	wg          sync.WaitGroup
}

// MoonshotPool is a worker group that runs a number of tasks at a
// configured concurrency.
type MoonshotPool struct {
	Tasks []*MoonshotTask

	concurrency int
	tasksChan   chan *MoonshotTask
	wg          sync.WaitGroup
}

// NewPool initializes a new pool with the given tasks and
// at the given concurrency.
func NewPool(tasks []*Task, concurrency int) *Pool {
	return &Pool{
		Tasks:       tasks,
		concurrency: concurrency,
		tasksChan:   make(chan *Task),
	}
}

// NewMoonshotPool initializes a new pool with the given tasks and
// at the given concurrency.
func NewMoonshotPool(tasks []*MoonshotTask, concurrency int) *MoonshotPool {
	return &MoonshotPool{
		Tasks:       tasks,
		concurrency: concurrency,
		tasksChan:   make(chan *MoonshotTask),
	}
}

// Run runs all work within the pool and blocks until it's
// finished.
func (p *Pool) Run() {
	for i := 0; i < p.concurrency; i++ {
		go p.work()
	}

	p.wg.Add(len(p.Tasks))
	for _, task := range p.Tasks {
		p.tasksChan <- task
	}

	// all workers return
	close(p.tasksChan)

	p.wg.Wait()
}

// The work loop for any single goroutine.
func (p *Pool) work() {
	for task := range p.tasksChan {
		task.Run(&p.wg)
	}
}

// Run runs all work within the pool and blocks until it's
// finished.
func (p *MoonshotPool) Run() {
	for i := 0; i < p.concurrency; i++ {
		go p.work()
	}

	p.wg.Add(len(p.Tasks))
	for _, task := range p.Tasks {
		p.tasksChan <- task
	}

	// all workers return
	close(p.tasksChan)

	p.wg.Wait()
}

// The work loop for any single goroutine.
func (p *MoonshotPool) work() {
	for task := range p.tasksChan {
		task.Run(&p.wg)
	}
}
