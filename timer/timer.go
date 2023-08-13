package timer

import (
	"sync"

	"github.com/robfig/cron/v3"
)

/*
   @File: timer.go
   @Author: khaosles
   @Time: 2023/8/13 10:56
   @Desc:
*/

type Timer interface {
	AddTaskByFunc(taskName string, spec string, task func(), option ...cron.Option) (cron.EntryID, error)
	AddTaskByJob(taskName string, spec string, job cron.Job, option ...cron.Option) (cron.EntryID, error)
	FindCron(taskName string) (*cron.Cron, bool)
	StartTask(taskName string)
	StopTask(taskName string)
	Remove(taskName string, id int)
	Clear(taskName string)
	Close()
}

func NewTimerTask() Timer {
	return &timer{taskList: make(map[string]*cron.Cron)}
}

// timer
type timer struct {
	taskList map[string]*cron.Cron
	sync.Mutex
}

// AddTaskByFunc add task by function
func (t *timer) AddTaskByFunc(taskName string, spec string, task func(), option ...cron.Option) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.taskList[taskName]; !ok {
		t.taskList[taskName] = cron.New(option...)
	}
	id, err := t.taskList[taskName].AddFunc(spec, task)
	t.taskList[taskName].Start()
	return id, err
}

// AddTaskByJob add task by job
func (t *timer) AddTaskByJob(taskName string, spec string, job cron.Job, option ...cron.Option) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.taskList[taskName]; !ok {
		t.taskList[taskName] = cron.New(option...)
	}
	id, err := t.taskList[taskName].AddJob(spec, job)
	t.taskList[taskName].Start()
	return id, err
}

// FindCron find the cron by task name
func (t *timer) FindCron(taskName string) (*cron.Cron, bool) {
	t.Lock()
	defer t.Unlock()
	v, ok := t.taskList[taskName]
	return v, ok
}

// StartTask start task
func (t *timer) StartTask(taskName string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.taskList[taskName]; ok {
		v.Start()
	}
}

// StopTask stop task
func (t *timer) StopTask(taskName string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.taskList[taskName]; ok {
		v.Stop()
	}
}

// Remove delete task
func (t *timer) Remove(taskName string, id int) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.taskList[taskName]; ok {
		v.Remove(cron.EntryID(id))
	}
}

// Clear clear task
func (t *timer) Clear(taskName string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.taskList[taskName]; ok {
		v.Stop()
		delete(t.taskList, taskName)
	}
}

// Close release the timer
func (t *timer) Close() {
	t.Lock()
	defer t.Unlock()
	for _, v := range t.taskList {
		v.Stop()
	}
}
