package helper

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math"
	"strconv"
)

type Pagination struct {
	Limit        int         `json:"limit,omitempty;query:limit"`
	CurrentPage  int         `json:"currentPage,omitempty;query:page"`
	Sort         string      `json:"sort,omitempty;query:sort"`
	Order        string      `json:"order,omitempty;query:order"`
	TotalElement int64       `json:"totalElement"`
	TotalPage    int         `json:"totalPage"`
	NextPage     int         `json:"nextPage"`
	PrevPage     int         `json:"prevPage"`
	Content      interface{} `json:"-"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.CurrentPage == 0 {
		p.CurrentPage = 1
	}
	return p.CurrentPage
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "id_cust"
	}
	return p.Sort
}

func (p *Pagination) GetOrder() string {
	if p.Order == "" {
		p.Order = "asc"
	}
	return p.Order
}
func PaginateQuery(value interface{}, pagination *Pagination, totalRows int64) func(db *gorm.DB) *gorm.DB {
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPage = totalPages
	pagination.NextPage = pagination.GetPage() + 1
	pagination.PrevPage = pagination.GetPage() - 1
	pagination.Content = value
	pagination.TotalElement = totalRows

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort() + " " + pagination.GetOrder())
	}
}

func GeneratePaginationFromRequest(c *gin.Context) Pagination {
	// Initializing default
	limit := 10
	page := 1
	sort := "transactions.created_at"
	order := "asc"
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]

		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break
		case "order":
			order = queryValue
			break
		}
	}
	return Pagination{
		Limit:       limit,
		CurrentPage: page,
		Sort:        sort,
		Order:       order,
	}
}
