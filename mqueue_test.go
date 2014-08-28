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
		var queue = NewQueue(100)
		convey.Convey("It shouldn't be running", func() {
			convey.So(queue.Running(), convey.ShouldBeFalse)
		})
		convey.Convey("It should be running after Run()", func() {
			queue.Run()
			convey.So(queue.Running(), convey.ShouldBeTrue)
		})
		convey.Convey("It shouldn't be running after Stop()", func() {
			queue.Stop()
			time.Sleep(time.Millisecond * 10)
			convey.So(queue.Running(), convey.ShouldBeFalse)
		})

		queue.Add(func() {
			convey.Convey("You shouldn't see me", func() {
				convey.So(false, convey.ShouldBeTrue)
			})
		})

		queue.Run()
		time.Sleep(time.Millisecond * 10)
		queue.Add(func() {
			convey.Convey("You should see me", func() {
				convey.So(true, convey.ShouldBeTrue)
			})
		})
		time.Sleep(time.Millisecond * 10)
		queue.Stop()
		time.Sleep(time.Millisecond * 10)
	})
}
