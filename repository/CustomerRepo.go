package repository

import (
	"booking-website-be/database"
	"booking-website-be/model"
	"context"
	"fmt"
	"time"
)

type AccountRepo interface {
	CreateCusRepo(ctx context.Context, customer model.CreateCus) error
	ViewCusListRepo(ctx context.Context) ([]model.Customer, error)
	ViewCusDetailRepo(ctx context.Context, customer_id string) ([]model.Customer, error)
	UpdateCusRepo(ctx context.Context, customer_id string, customer model.UpdateCus) error
	DeleteCusRepo(ctx context.Context, customer_id string, deleteBy model.DeleteCus) error
}

type AccountSql struct {
	Sql *database.Sql
}

func NewAccountRepo(sql *database.Sql) AccountRepo {
	return &AccountSql{
		Sql: sql,
	}
}

//Customer

// create customer
func (db *AccountSql) CreateCusRepo(ctx context.Context, customer model.CreateCus) error {
	query := ` insert into customer (full_name, email, phone_number, address, nationality, date_of_birth,id_document, registration_date, note, createtime, createby)
				 values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) `
	current := time.Now()

	if _, err := db.Sql.Db.Exec(query, customer.FullName, customer.Email, customer.PhoneNumber, customer.Address, customer.Nationality, customer.DateOfBirth, customer.IDDocument, customer.RegistrationDate, customer.Note, current, customer.CreateBy); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// view list customers
func (db *AccountSql) ViewCusListRepo(ctx context.Context) ([]model.Customer, error) {
	data := []model.Customer{}
	query := `select 
				customer_id, 
				full_name,
				email, 
				phone_number, 
				address, 
				nationality, 
				date_of_birth,
				id_document, 
				registration_date, 
				note from customer`
	if err := db.Sql.Db.Select(&data, query); err != nil {
		return []model.Customer{}, err
	}

	return data, nil
}

// view customer information
func (db *AccountSql) ViewCusDetailRepo(ctx context.Context, customer_id string) ([]model.Customer, error) {
	data := []model.Customer{}
	query := `	select 
				customer_id, 
				full_name,
				email, 
				phone_number, 
				address, 
				nationality, 
				date_of_birth,
				id_document, 
				registration_date, 
				note 
				from customer where customer_id = $1`
	if err := db.Sql.Db.Select(&data, query, customer_id); err != nil {
		return []model.Customer{}, err
	}

	return data, nil

}

// update customer information
func (db *AccountSql) UpdateCusRepo(ctx context.Context, customer_id string, customer model.UpdateCus) error {
	query := `update customer
				set full_name = $1,
				email = $2, 
				phone_number = $3, 
				address = $4, 
				nationality = $5, 
				date_of_birth = $6,
				id_document = $7, 
				registration_date = $8, 
				updatetime = $9,
				updateby = $10,
				note = $11 where customer_id = $12`
	current := time.Now()
	result, err := db.Sql.Db.Exec(query,
		customer.FullName,
		customer.Email,
		customer.PhoneNumber,
		customer.Address,
		customer.Nationality,
		customer.DateOfBirth,
		customer.IDDocument,
		customer.RegistrationDate,
		current,
		customer.UpdateBy,
		customer.Note,
		customer_id)
	if err != nil {
		return err
	}

	rowwAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowwAffected == 0 {
		return fmt.Errorf("column is not exist")
	}

	return nil
}

// delete customer
func (db *AccountSql) DeleteCusRepo(ctx context.Context, customer_id string, customer model.DeleteCus) error {
	query := `update customer
			set deletetime = $1,
				deleteby = $2
			where customer_id = $3`

	current := time.Now()

	result, err := db.Sql.Db.Exec(query, current, customer.DeleteBy, customer_id)
	if err != nil {
		return err
	}

	rowwAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowwAffected == 0 {
		return fmt.Errorf("column is not exist")
	}

	return nil
}
