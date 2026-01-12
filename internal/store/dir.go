package store

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

type DirStore struct {
	dir   string
	total int
}

func NewDirStore(dir string) (*DirStore, error) {

	nums, err := check(dir)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return &DirStore{dir: dir, total: len(nums)}, nil
}

func check(dir string) ([]int, error) {
	
	info, err := os.Stat(dir)
	if err != nil {
		return nil, fmt.Errorf("slides dir error: %w", err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("slides path is not a directory: %s", dir)
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read dir error: %w", err)
	}
	if len(entries) == 0 {
		return nil, fmt.Errorf("slides dir is empty")
	}

	nums := make([]int, 0, len(entries))

	for _, e := range entries {
		if e.IsDir() {
			return nil, fmt.Errorf("slides dir must contain only files named 1..n (found directory: %s)", e.Name())
		}
		n, convErr := strconv.Atoi(e.Name())
		if convErr != nil || n <= 0 {
			return nil, fmt.Errorf("invalid slide filename %q (must be a positive number)", e.Name())
		}
		nums = append(nums, n)
	}

	sort.Ints(nums)

	if nums[0] != 1 {
		return nil, fmt.Errorf("slide 1 not found")
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {
			return nil, fmt.Errorf(
				"slides must be numbered 1..N without gaps; expected %d, got %d",
				nums[i-1]+1, nums[i],
			)
		}
	}

	return nums, nil
}

func (s *DirStore) Total() int {
	return s.total
}

func (s *DirStore) Content(k int) (string, error) {
	name := strconv.Itoa(k)
	p := filepath.Join(s.dir, name)

	b, err := os.ReadFile(p)
	if err != nil {
		return "", fmt.Errorf("read slide %d: %w", k, err)
	}
	return string(b), nil
}
