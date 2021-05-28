package paginator

//...
const (
	DefaultPerPage = 10
	MaxPerPage     = 500
)

//Paginate  分页
type Paginate struct {
	Paginate struct {
		Page int `json:"Page"` //查询第几页
		Per  int `json:"Per"`  //每页数量
	} `json:"Paginate"`
	DisableLimit bool `json:"DisableLimit"` //是否关闭限制
}

//OffSet ...
func (p Paginate) OffSet() int {
	return (p.GetPage() - 1) * p.GetPer()
}

//SetPage ...
func (p Paginate) SetPage(page int) {
	p.Paginate.Page = page
}

//GetPage ...
func (p Paginate) GetPage() int {
	if p.Paginate.Page <= 0 {
		p.Paginate.Page = 1
	}
	return p.Paginate.Page
}

//GetPer ...
func (p Paginate) GetPer() int {
	if p.Paginate.Per <= 0 {
		p.Paginate.Per = DefaultPerPage
	}
	if !p.DisableLimit {
		if p.Paginate.Per > MaxPerPage {
			p.Paginate.Per = MaxPerPage
		}
	}
	return p.Paginate.Per
}

//Paginator ...
func Paginator(page int, per int) Paginate {
	p := Paginate{}
	p.Paginate.Page = page
	p.Paginate.Per = per
	return p
}

//NoLimit ...
func NoLimit(page int, per int) Paginate {
	p := Paginator(page, per)
	p.DisableLimit = true
	return p
}
