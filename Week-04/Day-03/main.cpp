#include <iostream>
#include <vector>
#include <list>
#include <utility>
using namespace std;

// Implementing Chaining in a Hash Table:
class ChainingHashTable {
private:
  vector<list<pair<int, int>>> table;
  size_t tableSize;

  size_t hashFunction(int key){
    return key%tableSize;
  }

public:
  ChainingHashTable(size_t size): tableSize(size){
    table.resize(tableSize);
  }

  void insert(int  key, int value){
    size_t index = hashFunction(key);
    auto& bucket = table[index];
    for(auto& pair: bucket){
      if(pair.first == key){
        pair.second = value;// Update existing value
        return;
      }
    }
    bucket.emplace_back(key, value); // Insert new Key-value pair  
  }

  bool search(int key, int& value){
    size_t index = hashFunction(key);
    auto& bucket = table[index];
    for(const auto& pair: bucket){
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
    for(auto it = bucket.begin(); it != bucket.end(); ++it){
      if(it->first == key){
        bucket.erase(it); // Remove key-value pair
        return;
      }
    }
  }
};

// Problem 1: Design a Hash Table with Chaining
void testChainingHashTable(){
  ChainingHashTable hashTable(10);

  hashTable.insert(15, 100);
  hashTable.insert(25, 200); // Collision with key 15
  hashTable.insert(35, 300); // Collision with keys 15 and 25 

  int value;
  if(hashTable.search(25, value)){
    cout << "Key 25 found with value: " <<  value << endl;
  } else {
    cout << "Key 25 not found" << endl;
  }

  hashTable.remove(25);
  if(hashTable.search(25, value)){
    cout << "Key 25  found with value: " << value << endl;
  } else {
    cout << "Key 25 not found" << endl;
  }
}

int main(){

  ChainingHashTable hashTable(10); // create hash table with 10 slots
  hashTable.insert(1, 100);
  hashTable.insert(11,200); //collision with key 1 
  hashTable.insert(21, 300); // Collision with key 1 and 11 
  int value;
  
  if(hashTable.search(11, value)){
    cout << "Key 11 found with value: " << value << endl;
  } else {
    cout << "Key 11 not found" << endl;
  }

  hashTable.remove(11);
  if(hashTable.search(11, value)){
  cout << "Key 11 found with value: " << value << endl;
  } else {
  cout << "key 11 not found" << endl;
  }

  // Problem 1: Design a Hash Table with Chaining

  testChainingHashTable();

  return 0;
}
