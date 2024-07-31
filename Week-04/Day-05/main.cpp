#include <iostream>
#include <vector>
#include <list> 
#include <utility>
using namespace std;


// Implementing Dynamic Resizing (Rehashing) for Hash Tables
class HashTable {
private:
  vector<pair<int, int>> table;
  size_t tableSize;
  size_t numElements;
  vector<bool> occupied;

  size_t hashFunction(int key){
    return key % tableSize;
  }

  void rehash(){
    size_t newTableSize = tableSize * 2;
    vector<pair<int, int>> newTable(newTableSize, {-1, -1});
    vector<bool> newOccupied(newTableSize, false);

    for(size_t i = 0; i < tableSize; ++i){
      if(occupied[i]){
        size_t newIndex = table[i].first % newTableSize;
        while(newOccupied[newIndex]){
          newIndex = (newIndex + 1) % newTableSize;
        }
        newTable[newIndex] = table[i];
        newOccupied[newIndex] = true;
      }
    }
    table = newTable;
    tableSize = newTableSize;
    occupied = newOccupied;
  }

public:
  HashTable(size_t size): tableSize(size), numElements(0){
    table.resize(tableSize, {-1, -1});
    occupied.resize(tableSize, false);
  }

  void insert(int key, int value){
    if (numElements > tableSize * 0.75){
      rehash();
    }

    size_t index = hashFunction(key);
    while(occupied[index]){
      index = (index + 1) % tableSize; 
    }
    table[index] = {key, value};
    occupied[index] = true;
  
    ++numElements;
  }

  bool search(int key, int& value){
    size_t index = hashFunction(key);
    while (occupied[index]){
      if(table[index].first == key){
        value = table[index].second;
        return true;
      }
      index = (index + 1) % tableSize;
    }
    return false; // Key not found
  }

  void remove(int key){
    size_t index = hashFunction(key);
    while(occupied[index]){
      if(table[index].first == key){
        table[index] = {-1, -1};
        occupied[index] = false;
        --numElements;
        return;
      }
      index = (index+1)%tableSize;
    } 
  }
};


int main(){

  HashTable hashTable(10); // Create hash table with 10 slots 
 
  hashTable.insert(1, 100); 
  hashTable.insert(21, 400); 
  hashTable.insert(31, 300);
  hashTable.insert(11, 500);

  int value;
  if(hashTable.search(11, value)){
    cout << "Key 11 found with value: "<< value << endl;
  } else {
    cout << "Key not found." << endl;
  }

  hashTable.remove(21);
  if(hashTable.search(21, value)){
    cout << "Key 21 found with value" << value << endl;
  } else {
    cout << "Key not found" << endl;
  }

  return 0;
}
