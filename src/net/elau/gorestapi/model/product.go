package model

type Product struct {
	ID          uint
	Name        string
	Type        string
	Description string
	Price       float32
}

type ProductDao struct{}

func NewProductDao() *ProductDao {
	return &ProductDao{}
}

func (pd *ProductDao) List() ([]*Product, error) {

	rows, err := db.Query("select * from product")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	plist := make([]*Product, 0)
	for rows.Next() {
		p := new(Product)
		err := rows.Scan(&p.ID, &p.Name, &p.Type, &p.Description, &p.Price)
		if err != nil {
			return nil, err
		}
		plist = append(plist, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return plist, nil
}
