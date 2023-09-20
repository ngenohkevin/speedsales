// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"context"
)

type Querier interface {
	CreateCodeTranslator(ctx context.Context, arg CreateCodeTranslatorParams) (CodeTranslator, error)
	CreateCustomer(ctx context.Context, arg CreateCustomerParams) (Customer, error)
	CreateDepartment(ctx context.Context, arg CreateDepartmentParams) (Department, error)
	CreateProducts(ctx context.Context, arg CreateProductsParams) (Product, error)
	CreateSales(ctx context.Context, arg CreateSalesParams) (Sale, error)
	CreateSales_till(ctx context.Context, arg CreateSales_tillParams) (SalesTill, error)
	CreateSupplier(ctx context.Context, arg CreateSupplierParams) (Supplier, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteCodeTranslator(ctx context.Context, masterCode string) error
	DeleteCustomer(ctx context.Context, customerID int64) error
	DeleteDepartment(ctx context.Context, departmentID int64) error
	DeleteProducts(ctx context.Context, productID int64) error
	DeleteSale(ctx context.Context, receiptNum int64) error
	DeleteSales_till(ctx context.Context, tillNum int64) error
	DeleteSupplier(ctx context.Context, supplierID int64) error
	DeleteUsers(ctx context.Context, userID int64) error
	GetCodeTranslator(ctx context.Context, masterCode string) (CodeTranslator, error)
	GetCustomer(ctx context.Context, customerID int64) (Customer, error)
	GetDepartment(ctx context.Context, departmentID int64) (Department, error)
	GetProducts(ctx context.Context, productID int64) (Product, error)
	GetSales(ctx context.Context, receiptNum int64) (Sale, error)
	GetSales_till(ctx context.Context, tillNum int64) (SalesTill, error)
	GetSupplier(ctx context.Context, supplierID int64) (Supplier, error)
	GetUser(ctx context.Context, userID int64) (User, error)
	ListCodeTranslator(ctx context.Context, arg ListCodeTranslatorParams) ([]CodeTranslator, error)
	ListCustomers(ctx context.Context, arg ListCustomersParams) ([]Customer, error)
	ListDepartment(ctx context.Context, arg ListDepartmentParams) ([]Department, error)
	ListProducts(ctx context.Context, arg ListProductsParams) ([]Product, error)
	ListSales(ctx context.Context, arg ListSalesParams) ([]Sale, error)
	ListSales_till(ctx context.Context, arg ListSales_tillParams) ([]SalesTill, error)
	ListSuppliers(ctx context.Context, arg ListSuppliersParams) ([]Supplier, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	UpdateCodeTranslator(ctx context.Context, arg UpdateCodeTranslatorParams) (CodeTranslator, error)
	UpdateCustomer(ctx context.Context, arg UpdateCustomerParams) (Customer, error)
	UpdateDepartment(ctx context.Context, arg UpdateDepartmentParams) (Department, error)
	UpdateProducts(ctx context.Context, arg UpdateProductsParams) (Product, error)
	UpdateSale(ctx context.Context, arg UpdateSaleParams) (Sale, error)
	UpdateSales_till(ctx context.Context, arg UpdateSales_tillParams) (SalesTill, error)
	UpdateSupplier(ctx context.Context, arg UpdateSupplierParams) (Supplier, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
