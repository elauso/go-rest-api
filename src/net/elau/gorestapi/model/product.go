package model

type Product struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
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

func (pd *ProductDao) Insert(p *Product) (*Product, error) {

	sqlStatement := `insert into product (name, type, description, price) values ($1, $2, $3, $4) returning id`
	var id uint = 0

	err := db.QueryRow(sqlStatement, p.Name, p.Type, p.Description, p.Price).Scan(&id)
	if err != nil {
		return nil, err
	}

	p.ID = id
	return p, nil
}
