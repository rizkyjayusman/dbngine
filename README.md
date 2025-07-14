# DBEngine

## Overview
This project is a prototype of a lightweight MySQL-like database engine built using Golang. It is designed to demonstrate fundamental database concepts such as query parsing, indexing, transaction management, and data storage. This project is a learning-oriented prototype of database engine, like MySQL but simplified to focus on core functionalities. The goal is to gain hands-on experience in database architecture, storage management, query processing, and transaction handling.

## Features
✅ **Storage Engine** (Basic file-based row storage)

✅ **SQL Query Processing** (`SELECT`, `INSERT`, `UPDATE`, `DELETE`)

✅ **Indexing** (B+ Tree for primary keys)

✅ **Query Execution** (Sequential scan + index lookup)

✅ **Transaction Support** (ACID, Write-Ahead Logging)

✅ **Simple CLI for Running Queries**

### Out of Scope

❌ **Joins & Complex Queries** (Only simple queries for now)

❌ **Advanced Query Optimizations** (No cost-based optimizer)

❌ **Replication & Sharding** (Single-node only)

❌ **User Management & Security** (No authentication)

## Installation

### Prerequisites

Ensure you have Go installed on your system. You can download it from Go's official website.

**Clone the Repository**
```
git clone https://github.com/yourusername/mysql-like-db.git
cd dbengine
```

**Build and Run**

```
go build -o dbengine
./dbengine
```

## Usage

### Start the Database Engine

Run the executable:
```
./dbengine
```

### Execute SQL Queries

Once running, you can interact with the database using SQL-like commands.

**Creating a Table**
```
CREATE TABLE users (
id INT PRIMARY KEY,
name TEXT,
email TEXT UNIQUE
);
```

**Inserting Data**
```
INSERT INTO users (id, name, email) VALUES (1, 'John Doe', 'john@example.com');
```

**Querying Data**
```
SELECT * FROM users;
```

**Updating Data**

```
UPDATE users SET name = 'Jane Doe' WHERE id = 1;
```

**Deleting Data**

```
DELETE FROM users WHERE id = 1;
```

## Architecture

![image info](./docs/dbengine.png)

### User Interface (CLI/API)

**Main Components:**

- **Command Line Interface (CLI)** → Input query from user.
- **Query Input Handler** → Send query input to Lexer.

### SQL Parser

1. **Lexing & Parsing**
- **Lexer** → Break query into tokens.
- **Tokenizer** → Populate token based on its type (keyword, identifier, operator).
- **Query Parser** → Analysis tokens and construct it as AST nodes.
2. **Query Optimization**
- **Query Rewriter** → Simplify the query if necessary.
- **Execution Plan Generator** → Plan the best ways to execute the query.
- **Cost Estimator (Basic)** → Measure performance of how query executed.

**Query Executor**

- **Query Executor** → Execute query based on execution plan.
- **Result Formatter** → Print the query result CLI formatted.

**Storage Engine**

1. **Data Management**
- **Table Manager** → Managing table metadata information.
- **Row & Column Storage** → Store data in row-based/column-based format.
2. Indexing System
- **B-Tree / Hash Index** → Better performance for searching.
3. **Transaction & Concurrency**
- **Write-Ahead Logging (WAL)** → Store log before write to disk.
- **Locking Mechanism** → Handle paralel transaction.
4. **Data Persistence**
- **File Storage System (Basic)** → Store data on binary file.
- **Buffer Pool** → Cache for prevent direct access to disk (direct acess necessary if the cache was none).

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests.

## License

This project is licensed under the MIT License.
