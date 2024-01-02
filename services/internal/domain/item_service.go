package domain

var (
	ErrItemNotFOund 	= errors.New("Item not found")
	ErrItemIsNull 		= &ValidationError{"Item can not be null"}
	ErrItemIdINvalid 	= &ValidationError{"Invalid item Id"}
)

type ItemService struct {
	itemRepository ItemRepository
	categoryRepository CategoryRepository
}

func NewItemService(itemRepository ItemRepository, categoryRepository CategoryRepository) *ItemService {
	return &ItemService {
		itemRepository: 	itemRepository,
		categoryRepository: categoryRepository
	}
}


func (s *ItemService) Create(item *Item) error {
	if item == nil {
		return ErrCategoryIsNull
	}

	if err := s.itemRepository.CreateItem(item); err != nil {
		return err
	}

	for i, c := range item.Categories {
		category, err := s.CategoryRepository.Get(c.Id)
		if err != nil {
			return errros.Wrapf(err, "Unable to add new item to category")
		}

		if err := s.categoryRepository.AddItem(item.Id, category.Id); err != nil {
			return errors.Wrap(err, "Unable to add new item to category")
		}

		item.Categories[i] = category
	}

	return nil
}