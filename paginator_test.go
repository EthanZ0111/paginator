package paginator

import (
	"fmt"
	"reflect"
	"testing"
)

type pagniatorTest struct {
	start    int
	end      int
	pgInput  Paginator
	pgOutput Paginator
}

var pgTestCase []*pagniatorTest = []*pagniatorTest{
	//0
	&pagniatorTest{
		pgInput:  Paginator{TotalCount: 11, PageSize: 3, CurPage: 1},
		pgOutput: Paginator{TotalCount: 11, PageSize: 3, CurPage: 1, TotalPage: 4},
		start:    1, end: 3,
	},
	//1
	&pagniatorTest{
		pgInput:  Paginator{TotalCount: 11, PageSize: 3, CurPage: 0},
		pgOutput: Paginator{TotalCount: 11, PageSize: 3, CurPage: 1, TotalPage: 4},
		start:    1, end: 3,
	},
	//2
	&pagniatorTest{
		pgInput:  Paginator{TotalCount: 11, PageSize: 3, CurPage: 4},
		pgOutput: Paginator{TotalCount: 11, PageSize: 3, CurPage: 4, TotalPage: 4},
		start:    10, end: 11,
	},
	//3
	&pagniatorTest{
		pgInput:  Paginator{TotalCount: 11, PageSize: 0, CurPage: 1},
		pgOutput: Paginator{TotalCount: 11, PageSize: 11, CurPage: 1, TotalPage: 1},
		start:    1, end: 11,
	},
	//4
	&pagniatorTest{
		pgInput:  Paginator{TotalCount: 11, PageSize: 3, CurPage: 5},
		pgOutput: Paginator{TotalCount: 11, PageSize: 3, CurPage: 4, TotalPage: 4},
		start:    10, end: 11,
	},
	//5
	&pagniatorTest{
		pgInput:  Paginator{TotalCount: 12, PageSize: 3, CurPage: 4},
		pgOutput: Paginator{TotalCount: 12, PageSize: 3, CurPage: 4, TotalPage: 4},
		start:    10, end: 12,
	},
	//6
	&pagniatorTest{
		pgInput:  Paginator{TotalCount: 12, PageSize: 20, CurPage: 1},
		pgOutput: Paginator{TotalCount: 12, PageSize: 20, CurPage: 1, TotalPage: 1},
		start:    1, end: 12,
	},
	//7
	&pagniatorTest{
		pgInput:  Paginator{TotalCount: 12, PageSize: 20, CurPage: 0},
		pgOutput: Paginator{TotalCount: 12, PageSize: 20, CurPage: 1, TotalPage: 1},
		start:    1, end: 12,
	},
}

func testPaginatorReturn(start, end int, pgTestCase pagniatorTest, pgOutput Paginator) error {
	if start != pgTestCase.start || end != pgTestCase.end {
		return fmt.Errorf("Start and end %d-%d is not correct in test case %#v", start, end, pgTestCase)
	}
	if !reflect.DeepEqual(pgOutput, pgTestCase.pgOutput) {
		return fmt.Errorf("Return paginator is not correct %#v in test case %#v", pgOutput, pgTestCase)
	}
	return nil
}

func TestPaginator(t *testing.T) {
	for idx, v := range pgTestCase {
		start, end := v.pgInput.GenerateIndexRange()
		if err := testPaginatorReturn(start, end, *v, v.pgOutput); err != nil {
			t.Error("Error test case ", idx, err)
			return
		}
	}
}
