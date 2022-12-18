package category

type IService interface {
	CreateCategory(input *CategoryCreateInput) (*Category, error)
	UpdateCategory(input *CategoryUpdateInput, categoryId int) (*Category, error)
	GetCategoryById(categoryId int) (*Category, error)
	DeleteCategory(categoryId int) error
	GetCategorys() ([]*Category, error)
}

type service struct {
	repo IRepository
}

func NewCategoryService(repository IRepository) IService {
	return &service{repo: repository}
}

func (s service) CreateCategory(input *CategoryCreateInput) (*Category, error) {
	var category *Category

	category = &Category{
		Type: input.Type,
	}

	category, err := s.repo.CreateCategory(category)
	if err != nil {
		return category, err
	}

	return CategoryCreateResponse(category), nil
}

func (s *service) GetCategoryById(categoryId int) (*Category, error) {
	category, err := s.repo.GetCategoryById(categoryId)
	if err != nil {
		return category, err
	}

	return category, nil
}

func (s *service) UpdateCategory(input *CategoryUpdateInput, categoryId int) (*Category, error) {
	category, err := s.GetCategoryById(categoryId)
	if err != nil {
		return category, err
	}

	category.Type = input.Type

	category = &Category{
		ID:                category.ID,
		Type:              category.Type,
		SoldProductAmount: category.ID * 5,
		UpdatedAt:         category.UpdatedAt,
	}

	return category, nil
}

func (s *service) DeleteCategory(categoryId int) error {
	if _, err := s.GetCategoryById(categoryId); err != nil {
		return err
	}

	if err := s.repo.DeleteCategory(categoryId); err != nil {
		return err
	}

	return nil
}

func (s service) GetCategorys() ([]*Category, error) {
	categorys, err := s.repo.GetCategorys()
	if err != nil {
		return categorys, err
	}

	return categorys, nil
}
