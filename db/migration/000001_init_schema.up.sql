-- Create Suppliers table
CREATE TABLE Suppliers (
                           SupplierID SERIAL PRIMARY KEY,
                           Name VARCHAR(255) NOT NULL,
                           Address VARCHAR(255),
                           ContactNumber VARCHAR(20),
                           Email VARCHAR(255)
);

-- Create Products table
CREATE TABLE Products (
                          ProductID SERIAL PRIMARY KEY,
                          Name VARCHAR(255) NOT NULL,
                          Description TEXT,
                          Category VARCHAR(255),
                          SupplierID INT,
                          Cost DECIMAL(10, 2),
                          SellingPrice DECIMAL(10, 2),
                          Quantity INT,
                          FOREIGN KEY (SupplierID) REFERENCES Suppliers(SupplierID)
);

-- Create Customers table
CREATE TABLE Customers (
                           CustomerID SERIAL PRIMARY KEY,
                           Name VARCHAR(255) NOT NULL,
                           Address VARCHAR(255),
                           ContactNumber VARCHAR(20),
                           Email VARCHAR(255)
);

-- Create Sales table
CREATE TABLE Sales (
                       SaleID SERIAL PRIMARY KEY,
                       ProductID INT,
                       CustomerID INT,
                       Quantity INT,
                       SaleDate DATE,
                       TotalPrice DECIMAL(10, 2),
                       FOREIGN KEY (ProductID) REFERENCES Products(ProductID),
                       FOREIGN KEY (CustomerID) REFERENCES Customers(CustomerID)
);
