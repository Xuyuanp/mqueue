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

import (
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
)

func TestQueue(t *testing.T) {
	convey.Convey("Given a new Queue", t, func() {
		queue := NewQueue(100)

		convey.Convey("It shouldn't be running", func() {
			convey.So(queue.Running(), convey.ShouldBeFalse)
		})

		convey.Convey("After Start()", func() {
			queue.Start()
			time.Sleep(time.Millisecond * 10)

			convey.Convey("It should be running", func() {
				convey.So(queue.Running(), convey.ShouldBeTrue)
			})

			convey.Convey("Add a task", func() {
				queue.AddTaskFunc(func(...interface{}) {
					convey.Convey("You should see me after a while", t, func() {
						convey.So(true, convey.ShouldBeTrue)
					})
				})
				time.Sleep(time.Millisecond * 10)
			})
		})

		convey.Convey("After Stop()", func() {
			queue.Stop()

			convey.Convey("It shouldn't be running", func() {
				convey.So(queue.Running(), convey.ShouldBeFalse)
			})

			convey.Convey("Add a task", func() {
				queue.AddTaskFunc(func(...interface{}) {
					convey.Convey("You shouldn't see me forever", t, func() {
						convey.So(false, convey.ShouldBeTrue)
					})
				})
				time.Sleep(time.Millisecond * 10)
			})
		})

		time.Sleep(time.Millisecond * 10)
	})
}
