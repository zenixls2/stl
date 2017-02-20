#include <map>
#include <iterator>
#include <utility>
#include <iostream>
#include "map.h"

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wc++11-extensions"

#define MAP_TYPE std::map<int, uintptr_t>
typedef MAP_TYPE::iterator iter_type;

Map MapInit()
{
	return (Map)(new MAP_TYPE());
}

void MapFree(Map m)
{
	std::map<int, uintptr_t>* mptr =
		static_cast<MAP_TYPE*>(m);
	delete mptr;
}

void MapAdd(Map m, int key, uintptr_t value)
{
	static_cast<MAP_TYPE*>(m)
		->insert(std::make_pair(key, value));
}

uintptr_t MapGet(Map m, int key)
{
	return (static_cast<MAP_TYPE*>(m)
			->find(key))
			->second;
}

typedef struct _iterator {
	iter_type iter;
	iter_type begin;
	iter_type end;
} _iterator;

MapIterator MapLowerBound(Map m, int key)
{
	auto ret = new _iterator();
	auto container = static_cast<MAP_TYPE*>(m);
	ret->iter = container->lower_bound(key);
	ret->begin = container->begin();
	ret->end = container->end();
	return static_cast<MapIterator>(ret);
}

MapIterator MapUpperBound(Map m, int key)
{
	auto ret = new _iterator();
	auto container = static_cast<MAP_TYPE*>(m);
	ret->iter = container->upper_bound(key);
	ret->begin = container->begin();
	ret->end = container->end();
	return static_cast<MapIterator>(ret);
}

uintptr_t MapIteratorGet(MapIterator iter)
{
	return (static_cast<_iterator*>(iter))->iter->second;
}

int MapIteratorGetKey(MapIterator iter)
{
	return (static_cast<_iterator*>(iter))->iter->first;
}

MapIterator MapIteratorNext(MapIterator iter)
{
	auto ret = (static_cast<_iterator*>(iter));
	ret->iter++;
	if (ret->iter == ret->end)
		return NULL;
	return iter;
}

MapIterator MapIteratorPrev(MapIterator iter)
{
	auto ret = (static_cast<_iterator*>(iter));
	if (ret->iter == ret->begin)
		return NULL;
	ret->iter--;
	return iter;
}

#pragma clang diagnostic pop
