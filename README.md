# LRU Cache API

Welcome to the LRU (Least Recently Used) Cache API! This project implements a simple yet efficient LRU caching mechanism with a twist - each cache item comes with an expiration time. Built with Go and the elegant Gin framework, this API provides a high-performance, concurrency-safe cache layer suitable for a wide range of applications, from optimizing backend systems to speeding up web applications.

## Features

- **LRU Eviction Policy**: Ensures the cache only holds the most recently accessed items, automatically evicting the least recently used ones when the cache is full.
- **Expiration Time**: Each key-value pair in the cache can have an expiration time, after which it gets automatically evicted, ensuring fresh data and efficient memory usage.
- **High Performance**: Written in Go and leveraging the Gin framework, this API is designed for speed and efficiency.
- **Concurrency Safe**: Uses Go's concurrency primitives to ensure that the cache is safe to use in a multi-threaded environment.
- **Simple RESTful API**: Provides straightforward GET and SET endpoints for easy interaction with the cache.

## Getting Started

Dive into the world of caching with our LRU Cache API. Follow these steps to get your cache server up and running in no time!

### Prerequisites

Before you embark on this caching adventure, ensure you have the following installed on your machine:
- Go (version 1.14 or later recommended)
- A sense of adventure

### Installation

1. **Clone the repository**

    ```bash
    git clone https://github.com/anujaagarwal/LRU_Cache.git
    cd LRU_Cache
    ```

2. **Run the server**

    Fire up your terminal and launch the cache into action:

    ```bash
    go run main.go
    ```

    Your LRU Cache API is now listening for cache interactions!

### Usage

Interact with the cache using its RESTful endpoints:

- **Set a key-value pair**

    ```bash
    GET /set?key=<your-key>&value=<your-value>&duration=<expiration-in-seconds>
    ```

    Set your desired key-value pair in the cache, along with an optional expiration duration in seconds.

- **Get a value by key**

    ```bash
    GET /get?key=<your-key>
    ```

    Retrieve the value associated with your specified key. Remember, if the key has expired or doesn't exist, you'll embark on a quest to find nothing.

## Example

Setting a key-value pair with a 60-second expiration:

```bash
curl "http://localhost:8080/set?key=hero&value=batman&duration=60s"
```

Retrieving the value for the key `hero`:

```bash
curl "http://localhost:8080/get?key=hero"
```

## Contributing

Join the cache crusade! Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Final Notes

Embark on this caching journey and unlock the potential of efficient data retrieval. Should you have any queries or require assistance, feel free to raise an issue in the repository. Happy caching! 🚀