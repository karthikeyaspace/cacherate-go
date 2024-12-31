<!-- notes.md -->
Caching and Ratelimiting in go

caching is a technique to store the data in memory or disk so that the next time the same data is requested, it can be served from the cache instead of the original source. This reduces the load on the original source and improves the performance of the application.

caching techniques
- In-memory caching 
- Disk caching (like memcached)
- Distributed caching (like Redis)

caching response
- X-Cache header - HIT or MISS



Ratelimiting is a technique to limit the number of requests that can be made to a service in a given time frame. This is done to prevent abuse of the service and to ensure that the service remains available to all users.

ratelimiting techniques
- Token Bucket
- Leaky Bucket
- Fixed Window
- Sliding Window

ratelimit response
- X-RateLimit-Limit
- X-RateLimit-Remaining
- X-RateLimit-Reset
- 429 Too Many Requests