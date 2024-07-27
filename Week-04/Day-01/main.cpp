#include <iostream>
#include <vector>
#include <list>
using namespace std;

class HashTable {
private:
  vector<list<pair<int, int>>> table;
  size_t tableSize;
  size_t hashFunction(int key){
    return key % tableSize;
 
  }
public:
  HashTable(size_t size): tableSize(size){
    table.resize(tableSize);
  }
  void insert(int key, int value){
    size_t index = hashFunction(key);
    auto& bucket = table[index];
    for(auto& pair : bucket) {
      if(pair.first == key){
        pair.second = value; // Update existing value
        return; 
      }
    }
    bucket.emplace_back(key,value); // Insert new key-value pair
  }

  bool search(int key, int& value){
    size_t index = hashFunction(key);
    auto& bucket = table[index];
    for(const auto& pair : bucket) {
      if(pair.first == key){
        value = pair.second;
        return true;
      }
    }
    return false; // key not found 
  }

  void remove(int key){
    size_t index = hashFunction(key);
    auto& bucket = table[index];
    for(auto it = bucket.begin(); it != bucket.end();++it){
      if(it->first == key){
        bucket.erase(it); //remove key-value pair
        return;
      }
    }
  }
};


int main(){
  HashTable hashTable(10); // create hash table with 10 slots 

  hashTable.insert(1, 100);
  hashTable.insert(2, 200);
  hashTable.insert(12, 300); // collisionwith key 2 
  
  int value; 

  if(hashTable.search(2,value)){
    cout << "Key 2 found with value: " << value << endl;
  } else {
    cout << "Key 2 not found" << endl;
  }

  hashTable.remove(2);
  if(hashTable.search(2, value)){
    cout << "Key 2 found with value: " << value << endl;
  } else {
    cout << "key not found\n";
  }
  return 0;
}


