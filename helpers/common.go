package helpers

import beego "github.com/beego/beego/v2/server/web"

func IsEmpty(s string) bool {
	return len(s) == 0
}

func IsNotEmpty(s string) bool {
	return len(s) > 0
}

func Remove(e interface{}, slice ...interface{}) {
	for i, v := range slice {
		if v == e {
			slice = append(slice[:i], slice[i+1:]...)
			return
		}
	}
}

func GetSortOffsetLimit(c *beego.Controller) (sort string, offset, limit int) {

	if v, err := c.GetInt("limit"); err == nil {
		limit = v
	}
	if v, err := c.GetInt("page"); err == nil {
		offset = v
	}
	if v := c.GetString("sort"); v != "" {
		sort = v
	}
	return
}
