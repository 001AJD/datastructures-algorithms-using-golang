# LRU cache

- A fixed size cache that evicts the least recently used item when the cache size is full
- Use Doubly linked list to store the actual data, provides O(1) insert, delete, shift operation. Will be used during inserting new entry in cache, deleting entry from cache
- Use hashmaps to access the node in 0(1) time complexity, will be used when accessing a node from linked list. Hashmap will store the node address

hashmap structure

```json
{
  index: nodeAddress
}
```

Linked List

```json
{
  index: int
  domain: string
  score: int
}
```

Algorithm

- prepopulate data in linked list during startup using insert commands
- max size of the cache = 5
- Whenever a fetch operation is issued
  - fetch data from linked list
  - store the recently fetched node's address in hashmap and move the node to the HEAD
- During inserts in cache
  - check if the max size of cache has been reached
  - if yes, remove the TAIL node
  - insert the new node at the HEAD
