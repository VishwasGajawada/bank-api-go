package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
}

type PostgressStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgressStore, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgressStore{
		db: db,
	}, nil
}

func (s *PostgressStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgressStore) createAccountTable() error {
	query := `create table if not exists account (
		id serial primary key,
		first_name varchar(50),
		last_name varchar(50),
		account_number serial,
		balance serial,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgressStore) CreateAccount(acc *Account) error {
	query := `insert into account (first_name, last_name, number, balance, created_at)
	values ($1, $2, $3, $4, $5)`
	resp, err := s.db.Query(
		query,
		acc.FirstName, acc.LastName, acc.AccountNumber, acc.Balance, acc.CreatedAt)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)
	return nil
}

func (s *PostgressStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgressStore) DeleteAccount(id int) error {
	_, err := s.db.Query("delete from account where id = $1", id)
	return err
}

func (s *PostgressStore) GetAccountByID(id int) (*Account, error) {
	rows, err := s.db.Query("select * from account where id = $1", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoAccount(rows)
	}
	return nil, fmt.Errorf("account %d not found", id)
}

func (s *PostgressStore) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query("select * from account")
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next() {
		account, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.AccountNumber,
		&account.Balance,
		&account.CreatedAt)

	return account, err
}
