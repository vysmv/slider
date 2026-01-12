package store

type SlideStore interface {
	Total() int
	Content(k int) (string, error) // k starts from 1
}
