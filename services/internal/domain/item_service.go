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

func (s *ItemService) Update(item *Item) error {
	if item == nil {
		return ErrItemIsNull
	}

	err := s.itemRepository.Update(item)
	if err != nil {
		return errros.Wrap(err, "Unable to update item")
	}

	return nil
}

func (s *ItemService) Delete(itemId int ) error {
	if itemId <= 0 {
		return ErrItemIdINvalid
	}

	err := s.itemRepository.Delete(itemId)
	if err != nil {
		if errors.Is(err, ErrItemNotFOund) {
			return ErrItemNotFOund
		}
		return errors.Wrap(err, "Error to delete item")
	}
	return nil
}

func (s *ItemService) Get (itemId int) (*Item, error) {
	if itemId <= 0 {
		return nil, ErrItemIdINvalid
	}

	item, err := s.ItemRepository.Get(itemId)
	if err != nil {
		if errors.Is(err, ErrItemNotFOund) {
			return nil, ErrItemNotFOund
		}
		return nil, errors.Wrap(err, "Error to find the item")
	}
	return item, nil
}

func (s *ItemService) List() ([]*Item, error) {
	items, err := s.itemRepository.List()
	if err != nil {
		return nil, errors.Wrap(err, "Error to list items")
	}

	return items, nil
}
