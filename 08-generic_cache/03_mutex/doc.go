// Package cache exposes a generic cache that can be used to
//   - store data by key
//   - retrieve data by key
//   - delete data by key
//
// The cache has no memory limit. All items will persist forever.
// The cache is thread-safe.
// The cache stores copies of user values, but it can be used with references.
// The most common syntax for using the cache is:
//
//	c := cache.New[K,V]()
//	...
//	v, found := c.Read(key)
//	if !found {
//	  v = ...
//	  c.Upsert(k, v)
//	}
package cache
