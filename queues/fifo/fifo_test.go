package fifo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/dbgoytia/datastructures/queues/fifo"
)

var _ = Describe("fifo", func() {

	Describe("The FIFO module is able to", func() {

		Context("enqueue elemenets", func() {

			It("if there are no elements, first element gets to be head of the queue", func() {
				// Expected queue
				// [1]
				// queue := fifo.Queue{}
				queue := fifo.NewQueue()
				queue.Enqueue(1)
				Expect(queue.Entries[0]).To(Equal(1))
			})

			It("enqueues values to existing queues", func() {
				// Expected queue
				// [0, 1, 2, 3, 4, 5, 6, 7, 8]
				queue := fifo.NewQueue()
				for i := 0; i < 9; i++ {
					queue.Enqueue(i)
				}
				Expect(queue.Entries[len(queue.Entries)-1]).To(Equal(8))
			})
		})

		Context("dequeue elements", func() {

			It("knows when the queue doesn't have any elements", func() {
				// Expected queue is empty
				queue := fifo.NewQueue()
				_, err := queue.Dequeue()
				Expect(err.Error()).To(Equal("empty queue"))
			})

			It("removes an element from beginning of queue", func() {
				// Expected queue
				// [1, 2, 3, 4, 5, 6]
				queue := fifo.NewQueue()
				for i := 0; i < 7; i++ {
					queue.Enqueue(i)
				}
				queue.Dequeue()
				// Expect(queue.Entries[len(queue.Entries)-1]).To(Equal(8))
				expected := []int{1, 2, 3, 4, 5, 6}
				Expect(queue.Entries).To(Equal(expected))
			})
		})

		Context("Peeking the queue", func() {
			It("from an empty queue", func() {
				queue := fifo.NewQueue()
				_, err := queue.Peek()
				Expect(err.Error()).To(Equal("empty queue"))
			})

			It("from a queue with a lot of values", func() {
				queue := fifo.NewQueue()
				for i := 0; i < 7; i++ {
					queue.Enqueue(i)
				}
				Expect(queue.Peek()).To(Equal(0))
			})
		})

	})
})
