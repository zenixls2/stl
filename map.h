#ifdef __cplusplus
extern "C" {
#endif /* __cplusplus  */
	typedef void* Map;
	typedef void* MapIterator;
	Map MapInit(void);
	void MapFree(Map);
	void MapAdd(Map m, int key, uintptr_t value);
	uintptr_t MapGet(Map m, int key);
	MapIterator MapLowerBound(Map m, int key);
	MapIterator MapUpperBound(Map m, int key);
	uintptr_t MapIteratorGet(MapIterator iter);
	int MapIteratorGetKey(MapIterator iter);
	MapIterator MapIteratorNext(MapIterator iter);
	MapIterator MapIteratorPrev(MapIterator iter);
#ifdef __cplusplus
}
#endif /* __cplusplus */
