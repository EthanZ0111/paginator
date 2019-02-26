package paginator

import "math"

type Paginator struct {
	PageSize   int
	CurPage    int
	TotalPage  int
	TotalCount int
	SortKey    interface{}
}

func (p *Paginator) GenerateIndexRange() (start int, end int) {
	// p.PageSize = 0
	start = 0
	end = p.TotalCount
	if p.PageSize <= 0 {
		p.PageSize = p.TotalCount
		// return
	}
	if p.CurPage <= 0 {
		p.CurPage = 1
	}
	p.TotalPage = int(math.Ceil(float64(p.TotalCount) / float64(p.PageSize)))
	if p.CurPage > p.TotalPage {
		p.CurPage = p.TotalPage
	}
	start = (p.CurPage-1)*p.PageSize + 1
	end = p.CurPage * p.PageSize //start + p.PageSize - 1
	if end >= p.TotalCount {
		end = p.TotalCount
		return
	}
	return
}
