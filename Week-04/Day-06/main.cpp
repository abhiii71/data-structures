#include <iostream>
#include <unordered_map>
#include <list>
using namespace std;

// LRU (Least Recently Used) Cache 
class LRUCache {
private:
  int capacity;
  list<pair<int, int>> cacheList;
  unordered_map<int, list<pair<int, int>>::iterator> cacheMap;

public:
  LRUCache(int cap): capacity(cap) {}

  int get(int key){
    if(cacheMap.find(key) == cacheMap.end()){
      return -1; // key not found
    }
    auto it = cacheMap[key];
    int value = it->second;
    cacheList.erase(it);
    cacheList.push_front({key, value});
    cacheMap[key] = cacheList.begin();
    return value;
  }

  void put(int key, int value){
    if(cacheMap.find(key) != cacheMap.end()){
      auto it = cacheMap[key];
      cacheList.erase(it);
    } else if (cacheList.size() == capacity){
      auto last = cacheList.back();
      cacheMap.erase(last.first);
      cacheList.pop_back();
    }
    cacheList.push_front({key, value});
    cacheMap[key] = cacheList.begin();
  }
};

// LFU (Least Frequently Used) Cache
class LFUCache {
private:
  int capacity;
  int minFreq;
  struct Node {
    int key, value, freq;
    Node(int k, int v, int f): key(k), value(v), freq(f) {}
  };

  unordered_map<int, list<Node>::iterator> cache;
  unordered_map<int, list<Node>> freqList;

public:
  LFUCache(int cap): capacity(cap), minFreq(0) {}

  int get(int key){
    if(capacity == 0 || cache.find(key) == cache.end()) {
      return -1;
    }

    auto it = cache[key];
    int value = it->value;
    int freq = it->freq;
    freqList[freq].erase(it);
    if(freqList[freq].empty()){
      freqList.erase(freq);
      if(minFreq == freq) minFreq++;
    }
    freqList[freq+1].push_front(Node (key, value, freq+1));
    cache[key] = freqList[freq+1].begin();
    return value;
  }

  void put(int key, int value){
    if(capacity == 0) return;

    if(cache.find(key) != cache.end()){
      auto it = cache[key];
      int freq = it->freq;
      freqList[freq].erase(it);
      if(freqList[freq].empty()) {
        freqList.erase(freq);
        if(minFreq == freq) minFreq++;
      }
      freqList[freq+1].push_front(Node(key, value, freq+1));
      cache[key] = freqList[freq + 1].begin();
    } else {
      if(cache.size() == capacity) {
        auto it = freqList[minFreq].back();
        cache.erase(it.key);
        freqList[minFreq].pop_back();
        if(freqList[minFreq].empty()){
          freqList.erase(minFreq);
        }
      }
      freqList[1].push_front(Node(key, value, 1));
      cache[key] = freqList[1].begin();
      minFreq = 1;
    }
  }
};


int main(){

  // LRU (Least Recently Used) Cache 
  cout << "LRU (Least Recently Used) Cache ";
  LRUCache cache(2);
  cache.put(1,1);
  cache.put(2,2);
  cout << cache.get(1) << endl; // returns 1
  cache.put(3,3); // evicate key 2
  cout << cache.get(2) << endl; // returns -1 (not found)
  cache.put(4,4);
  cout << cache.get(1) << endl; // returns -1 (not found)
  cout << cache.get(3) << endl; // returns 3
  cout << cache.get(4) << endl; // returns 4

  // LFU (Least Frequently Used) Cache

  cout << " LFU (Least Frequently Used) Cache";
  LFUCache cache2(2);
  cache2.put(1, 1);
  cache2.put(2, 2);
  cout << cache2.get(1) << endl; // return 1 
  cache2.put(3, 3); // evicts key 2 
  cout << cache2.get(2) << endl; // return -1 (not found)
  cout << cache2.get(3) << endl; // return 3 
  cache2.put(4,4); // evicts 1 
  cout << cache2.get(1) << endl; // return -1(not found)
  cout << cache2.get(3) << endl; // return  3 
  cout << cache2.get(4) << endl; // return 4
  
  return 0;
}
