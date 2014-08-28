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

type Worker func()

func (w Worker) Do() {
	w()
}

type Queue struct {
	ch      chan Worker
	sw      chan bool
	MaxSize int
	running bool
}

func NewQueue(size int) *Queue {
	return &Queue{
		MaxSize: size,
		running: false,
	}
}

func (q *Queue) Run() {
	if q.running {
		return
	}
	q.running = true
	q.ch = make(chan Worker, q.MaxSize)
	go func() {
		for {
			w, ok := <-q.ch
			if !ok {
				return
			}
			w.Do()
		}
	}()
}

func (q *Queue) Stop() {
	if q.running {
		q.running = false
		close(q.ch)
	}
}

func (q *Queue) Add(w Worker) {
	if q.running {
		q.ch <- w
	}
}

func (q *Queue) Running() bool {
	return q.running
}
