package category

type CRepository interface {
	Create(c *Category) (*Category, error)
	FindAll() ([]*Category, error)
	FindByID(id uint) (*Category, error)
	FindBySlug(slug string) (*Category, error)
}
