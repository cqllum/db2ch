![db2ch](https://i.imgur.com/Vq4PBVd.png)


# db2ch

`db2ch` is an open-source framework designed to simplify Change Data Capture (CDC) to ClickHouse without the need for complex setups like Debezium, Kafka, or ClickHouse Materialized Views. This project supports real-time data replication from PostgreSQL, MySQL, and MSSQL to ClickHouse.


## Features

- **Real-time Replication**: Efficiently stream data changes from PostgreSQL, MySQL, and MSSQL to ClickHouse in real-time.
- **Simplicity**: No need for additional tools like Debezium, Kafka, or ClickHouse Materialized Engines.
- **High Throughput**: Built with Go, ensuring high performance and low latency.
- **Control and Monitor**: HTTP APIs for controlling and monitoring replication processes.
- **Automatic Table Creation**: `db2ch` will automatically map source data types to Clickhouse data types and create tables.
- **Primary Key Based Backfill Capability**: `db2ch` also includes a capability to perform a backfil to bring across historic data.

## Project Structure
```
db2ch/
├── main.go
├── config/
│ ├── config.json
├── api/
│ ├── handler.go
│ └── routes.go
├── replication/
│ ├── replication.go
│ ├── clickhouse.go
│ ├── mysql.go
│ └── mssql.go
└── go.mod
```
### Folder Descriptions
- `config/` - The purpose of this directory is define configuration details (connection details, data mapping, http port, etc).
- `api/` - Contains the `routes` for HTTP API endpoints and a handler for core api functionality.
- `replication/` - The framework behind listening to data and producing data.

## Getting Started

### Prerequisites

- Go 1.16+ installed
- PostgreSQL, MySQL, and MSSQL databases set up
- ClickHouse set up

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/cqllum/db2ch.git
    cd db2ch
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```
