/*
 * Copyright 2014 Xuyuan Pang <xuyuanp # gmail.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mqueue

type Task interface {
	Do(v ...interface{})
}

type TaskFunc func(...interface{})

func (t TaskFunc) Do(v ...interface{}) {
	t(v...)
}

type Queue struct {
	V       []interface{}
	ch      chan Task
	MaxSize int
	running bool
	init    bool
}

func NewQueue(size int) *Queue {
	q := &Queue{
		V:       make([]interface{}, 0),
		MaxSize: size,
		running: false,
		init:    false,
	}

	return q
}

func (q *Queue) Init() {
	if q.init {
		return
	}
	q.init = true
	q.ch = make(chan Task, q.MaxSize)
	go func() {
		for {
			select {
			case t, ok := <-q.ch:
				if !ok {
					return
				}
				t.Do(q.V...)
			}
		}
	}()
}

func (q *Queue) Stop() {
	q.running = false
}

func (q *Queue) Start() {
	q.Init()
	q.running = true
}

func (q *Queue) AddTask(ts ...Task) {
	if q.running {
		for _, t := range ts {
			q.ch <- t
		}
	}
}

func (q *Queue) AddTaskFunc(ts ...TaskFunc) {
	if q.running {
		for _, t := range ts {
			q.ch <- t
		}
	}
}

func (q *Queue) Running() bool {
	return q.running
}

func (q *Queue) Destroy() {
	q.Stop()
	if q.init {
		close(q.ch)
		q.init = false
	}
}
