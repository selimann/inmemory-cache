# In-Memory Sharded Cache in Go

This project is a simple implementation of an **in-memory cache** in Go, built with **sharding** and **concurrent access support** using `sync.RWMutex`.

The goal is to demonstrate a practical approach to handling concurrent reads/writes in a cache-like structure, while distributing data across shards to reduce lock contention.

---

## Features

* **Sharded design** – keys are distributed across multiple shards using an FNV hash function.
* **Thread-safe** – protects shared state with `sync.RWMutex`.
* **Methods** – provides `Set` and `Get` methods.
* **Concurrent usage example** – shows how the cache behaves under concurrent goroutines.

---

## Code Structure

* `InMemoryCache` – main cache structure holding multiple shards.
* `Shard` – contains a `map[string]string` and a mutex for synchronization.
* `hasher()` – computes shard index from a key using FNV hashing.
* `Set()` / `Get()` – public API methods to interact with the cache.
* `main()` – example usage with goroutines and `sync.WaitGroup`.

