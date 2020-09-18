package model

import "database/sql"

type Product struct {
	ID          uint64  `json:"id"`
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

	rows, err := db.Query("SELECT * FROM product")
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

func (pd *ProductDao) FindById(id uint64) (*Product, error) {

	sqlStatement := `SELECT * FROM product WHERE id=$1`
	p := &Product{}

	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&p.ID, &p.Name, &p.Type, &p.Description, &p.Price)
	switch err {
	case nil:
		return p, nil
	case sql.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}
}

func (pd *ProductDao) Insert(p *Product) (*Product, error) {

	sqlStatement := `INSERT INTO product(name, type, description, price) VALUES($1, $2, $3, $4) RETURNING id`
	var id uint64 = 0

	err := db.QueryRow(sqlStatement, p.Name, p.Type, p.Description, p.Price).Scan(&id)
	if err != nil {
		return nil, err
	}

	p.ID = id
	return p, nil
}
