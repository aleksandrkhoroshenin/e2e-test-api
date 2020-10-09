package repository

import (
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

type CustomerRepositoryImpl struct {
	dbConnection *pg.DB
}

func NewCustomerRepositoryImpl(dbConnection *pg.DB) *CustomerRepositoryImpl {
	return &CustomerRepositoryImpl{
		dbConnection: dbConnection,
	}
}

func (c *CustomerRepositoryImpl) GetByPrefix(prefix []string) (customersInfo []*CustomerInfo, err error) {
	err = c.dbConnection.
		Model(&customersInfo).
		WhereGroup(func(query *orm.Query) (*orm.Query, error) {
			for _, p := range prefix {
				query = query.WhereOr("first_name like ?", p + "%")
			}
			return query, nil
		}).Select()

	return customersInfo, err
}

func (c *CustomerRepositoryImpl) DeleteByPrefix(prefix []string) (customers []CustomerInfo, err error) {
	_, err = c.dbConnection.Model(&customers).WhereGroup(func(query *orm.Query) (*orm.Query, error) {
		for _, p := range prefix {
			query = query.WhereOr("first_name like ?", p + "%")
		}
		return query, nil
	}).Returning("*").Delete()
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (c *CustomerRepositoryImpl) CreateCustomer(customer CustomerInfo) (result int, err error) {
	res, err := c.dbConnection.Model(&customer).WherePK().Insert()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected(), nil
}