GoVault - A Distributed Key-Value Store

GoVault is a scalable, distributed key-value store built in Golang. It supports sharding, replication, and multi-database management, making it a lightweight alternative to Redis for specific use cases.

🚀 Features

✅ Multi-Database Support (DB 0, DB 1, etc.) 

✅ Sharding with Consistent Hashing (Distributes data across multiple nodes)

✅ Leader-Follower Replication (Ensures high availability)

✅ TTL-based Expiry (Keys expire automatically)

✅ Write-Ahead Logging (AOF Persistence)

✅ Raft Consensus for Leader Election (Ensures fault tolerance)

✅ HTTP API for Data Operations

📌 Installation

Prerequisites

Go 1.23+ installed

Clone the Repository

git clone https://github.com/yourusername/govault.git
cd govault

Run the Server

go run main.go

🔹 API Endpoints

Method

Endpoint

Description

GET

/get/{key}

Get a value by key

POST

/set

Set a key-value pair

DELETE

/delete/{key}

Delete a key

POST

/set_with_ttl

Set a key with TTL (expiration)

POST

/select_db

Switch active database (DB 0, etc.)

Example Requests

Set a Key

curl -X POST http://localhost:8080/set -d '{"key": "foo", "value": "bar"}' -H "Content-Type: application/json"

Get a Key

curl -X GET http://localhost:8080/get/foo

Delete a Key

curl -X DELETE http://localhost:8080/delete/foo

🔹 Scaling GoVault

1️⃣ Sharding (Distributing Data Across Nodes)

Uses consistent hashing (HashRing) to distribute keys.

Efficient key remapping when adding/removing nodes.

2️⃣ Replication (Ensuring High Availability)

Implements Leader-Follower replication.

Uses Raft Consensus for automatic leader election.

Followers replicate data from the leader.

3️⃣ Persistence & Durability

Write-Ahead Logging (WAL) ensures data is not lost.

Append-Only File (AOF) Logging enables recovery after crashes.

📌 Future Enhancements

🔹 Cluster Mode for distributed storage

🔹 Support for gRPC API

🔹 Redis Protocol Compatibility

📜 License

MIT License

💡 Contributing

We welcome contributions! Please submit a pull request or open an issue.

🔗 Connect with Us

📧 Email: hhimalayaa55@gmail.com🐙 GitHub: [github.com/hhimalayaa/govault]

