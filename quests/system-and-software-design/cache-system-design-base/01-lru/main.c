#include <stdlib.h>
#include <string.h>

#define MAX_KEY 10001

typedef struct Node {
    int key;
    int value;
    struct Node *prev;
    struct Node *next;
} Node;

typedef struct {
    int capacity;
    int count;
    Node *head;
    Node *tail;
    Node **nodeMap;
    Node *nodePool;
    Node *freeNode;
} LRUCache;

LRUCache *lRUCacheCreate(int capacity) {
    LRUCache *cache = malloc(sizeof(LRUCache));
    cache->nodeMap = calloc(MAX_KEY, sizeof(Node *));
    cache->nodePool = malloc(sizeof(Node) * capacity);

    cache->capacity = capacity;
    cache->count = 0;
    cache->head = NULL;
    cache->tail = NULL;
    return cache;
}

void moveToHead(LRUCache *obj, Node *node) {
    // Node is already at start
    if (node == obj->head) {
        return;
    }

    if (node->prev) {
        node->prev->next = node->next;
    }
    if (node->next) {
        node->next->prev = node->prev;
    }

    // Node is last one
    if (node == obj->tail) {
        obj->tail = node->prev;
    }

    node->next = obj->head;
    node->prev = NULL;

    if (obj->head) {
        obj->head->prev = node;
    }
    obj->head = node;

    if (obj->tail == NULL) {
        obj->tail = node;
    }
}

int lRUCacheGet(LRUCache *obj, int key) {
    Node *node = obj->nodeMap[key];

    if (node == NULL) {
        return -1;
    }

    moveToHead(obj, node);
    return node->value;
}

void lRUCachePut(LRUCache *obj, int key, int value) {
    Node *node = obj->nodeMap[key];

    // Key is already there
    if (node != NULL) {
        node->value = value;
        moveToHead(obj, node);
    } else {
        // Key has not been used
        if (obj->count < obj->capacity) {
            // There is enough space
            node = &obj->nodePool[obj->count++];
            node->key = key;
            node->value = value;

            node->next = obj->head;
            node->prev = NULL;

            if (obj->head) {
                obj->head->prev = node;
            }
            obj->head = node;

            if (!obj->tail) {
                obj->tail = node;
            }
            obj->nodeMap[key] = node;
        } else {
            Node *lru = obj->tail;
            obj->nodeMap[lru->key] = NULL;
            lru->key = key;
            lru->value = value;

            obj->nodeMap[key] = lru;
            moveToHead(obj, lru);
        }
    }
}

void lRUCacheFree(LRUCache *obj) {
    if (obj) {
        if (obj->nodeMap)
            free(obj->nodeMap);
        if (obj->nodePool)
            free(obj->nodePool);
        free(obj);
    }
}

/**
 * Your LRUCache struct will be instantiated and called as such:
 * LRUCache* obj = lRUCacheCreate(capacity);
 * int param_1 = lRUCacheGet(obj, key);

 * lRUCachePut(obj, key, value);

 * lRUCacheFree(obj);
*/
