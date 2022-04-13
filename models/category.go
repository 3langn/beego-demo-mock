package models

type CategoryDto struct {
	Title string `json:"title"`
}

type Category struct {
	Base
	Title string  `json:"title"`
	Foods []*Food `json:"foods" gorm:"many2many:food_categories"`
}

//func (c *Category) TableName() string {
//	return "categories"
//}

func (c *Category) Create(dto []*Category) ([]*Category, error) {
	var categories []*Category
	for _, v := range dto {
		categories = append(categories, &Category{
			Title: v.Title,
		})
	}

	if err := GetDB().Save(categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *Category) GetById(id string) error {
	if err := GetDB().First(&c, "id = ?", id).Error; err != nil {
		return err
	}
	GetDB().Debug().Model(&c).Association("Foods").Find(&c.Foods)
	return nil
}

func (c *Category) Update(id string, dto Category) error {
	if err := GetDB().Model(Category{}).Where("id = ?", id).Updates(dto).Error; err != nil {
		return err
	}

	return nil
}

func (c *Category) Delete(id string) error {
	if err := GetDB().Model(Category{}).Where("id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (c *Category) GetAll() ([]Category, error) {
	var categories []Category
	if err := GetDB().Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *Category) GetAllCategory(sort string, offset int, limit int) ([]Category, error) {
	var categories []Category
	if err := GetDB().Paging(sort, offset, limit).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
