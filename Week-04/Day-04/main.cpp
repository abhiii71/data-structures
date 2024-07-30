#include <iostream>
#include <vector>
#include <utility>
using namespace std;

// Implementing Open Addressing in a Hash Table


// Linear Probing Implementation:
class LinearProbingHashTable{
private:
  vector<pair<int, int>> table;
  size_t tableSize;
  vector<bool> occupied;

  size_t hashFunction(int key) {
    return key % tableSize;
  }

public:
  LinearProbingHashTable(size_t size): tableSize(size){
    table.resize(tableSize, {-1, -1});
    occupied.resize(tableSize, false);
  }

  void insert(int key, int value){
    size_t index = hashFunction(key);
    size_t originalIndex = index;
    while(occupied[index]){
      index = (index+1)%tableSize;
      if(index == originalIndex){
        cout << "Hash Table is full\n";
        return;
      }
    }
    table[index] = {key, value};
    occupied[index] = true;
  }

  bool search(int key, int& value){
    size_t index = hashFunction(key);
    size_t originalIndex = index;
    while(occupied[index]){
      if(table[index].first == key){
        value = table[index].second;
        return true;
      }
      index = (index+1)%tableSize;
      if(index == originalIndex){
        return false;
      }
    }
    return false;
  }

  void remove(int key){
    size_t index = hashFunction(key);
    size_t originalIndex = index;
    while(occupied[index]){
      if(table[index].first == key){
        table[index] = {-1, -1};
        occupied[index] = false;
        return;
      }
      index = (index+1)%tableSize;
      if(index == originalIndex){
        return;
      }
    }
  }
};

// Quadratic Probing Implementation:
class QuadaraticProbingHashTable {
private:
  vector<pair<int, int>> table;
  size_t tableSize;
  vector<bool> occupied;

  size_t hashFunction(int key){
    return key % tableSize;
  }

public:
  QuadaraticProbingHashTable(size_t size): tableSize(size) {
    table.resize(tableSize, {-1, -1});
    occupied.resize(tableSize, false);
  }

  void insert(int key, int value){
    size_t index = hashFunction(key);
    size_t i = 0;
    while(occupied[(index+i*i)%tableSize]){
      i++;
      if(i == tableSize){
        cout << "Hash table is full/n";
        return;
      }
    }
    index = (index+i*i) % tableSize;
    table[index] = {key, value};
    occupied[index] = true;
  }

  bool search(int key, int& value){
    size_t index = hashFunction(key);
    size_t i = 0;
    while(occupied[(index+i*i) % tableSize]){
      index = (index+i*i) % tableSize;
      if(table[index].first == key){
        value = table[index].second;
        return true;
      }
      i++;
      if(i == tableSize){
        return false;
      }
    }
    return false;
  }

  void remove(int key){
    size_t index = hashFunction(key);
    size_t i = 0;
    
    while(occupied[(index+i*i) % tableSize]){
      index = (index+i*i)% tableSize;
      if(table[index].first == key){
        table[index] = {-1, -1};
        occupied[index] = false; 
        return;
      }
      i++;
      if(i == tableSize){
        return;
      }
    }
  }
};

// Double Hashing Implementation:
class DoubleHashingHashTable {
private:
  vector<pair<int, int>> table;
  size_t tableSize;
  vector<bool>occupied;

  size_t hashFunction1(int key){
    return key % tableSize;
  }

  size_t hashFunction2(int key){
    return (key / tableSize) % tableSize;
  }
  
public:
  DoubleHashingHashTable(size_t size): tableSize(size){
    table.resize(tableSize, {-1, -1});
    occupied.resize(tableSize, false);
  }

  void insert(int key,int value){
    size_t index1 = hashFunction1(key);
    size_t index2 = hashFunction2(key);
    size_t i = 0;
    while(occupied[(index1+i* index2) % tableSize]){
      i++;
      if(i == tableSize){
        cout << "Hash Function is full\n";
        return;
      }
    }
    index1 = (index1 + i * index2) % tableSize;
    table[index1] = {key, value};
    occupied[index1] = true;
  }

  bool search(int key, int& value){
    size_t index1 = hashFunction1(key);
    size_t index2 = hashFunction2(key);
    size_t i = 0;
    while(occupied[(index1 + i * index2) % tableSize]){
      index1 = (index1 + i * index2) % tableSize;
      if(table[index1].first == key){
        value = table[index1].second;
        return true;
      }
      i++;
      if(i == tableSize){
        return false;
      }
    }
    return false;
  }

  void remove(int key){
    size_t index1 = hashFunction1(key);
    size_t index2 = hashFunction2(key);
    size_t i = 0;
    while(occupied[(index1 + i * index2)% tableSize]) {
      index1 = (index1 + i * index2) % tableSize;
      if(table[index1].first == key){
        table[index1] = {-1, -1};
        occupied[index1] = false;
        return;
      }
      i++;
      if(i == tableSize){
        return;
      }
    }
  }

};

int main(){

  // Linear Probing Implementation:
  LinearProbingHashTable hashTable(10);
  hashTable.insert(1,100);
  hashTable.insert(11, 200); // Collision with key 1 
  hashTable.insert(21, 300); // Colision with key 1 and 11
  
  int value;

  if(hashTable.search(11, value)){
    cout << " Key 11 found with value: " << value << endl;
  } else {
    cout << "Key not found" << endl;
  }

  hashTable.remove(11);

 
  if (hashTable.search(11, value)) {
    std::cout << "Key 11 found with value: " << value << std::endl;
  } else {
    std::cout << "Key 11 not found" << std::endl;
  }


  // Quadratic Probing Implementation:
  QuadaraticProbingHashTable hashTablee(10);
  hashTablee.insert(1, 100);
  hashTablee.insert(11, 200);
  hashTablee.insert(21, 300);

  int valuee;

  if(hashTablee.search(11, valuee)){
    cout << "Key 11 found with value: " << valuee << endl;
  } else {
    cout << "Key 11 not found." << endl;
  }

  hashTablee.remove(11);

  if(hashTable.search(11, valuee)){
    cout << "Key 11 found with value: " << valuee << endl;
  } else {
    cout << "Key 11 not found" << endl;
  }

  // Double Hashing Implementation:

  DoubleHashingHashTable hashTableee(10);
  
  hashTableee.insert(11, 200); // Collision with key 1
  hashTableee.insert(21, 300); // Collision with key 1 and 11
  hashTableee.insert(12, 500); // Collision with key 2
  hashTableee.insert(22, 600); // Collision with key 2 and 12
  hashTableee.insert(3, 700);
  hashTableee.insert(13, 800); // Collision with key 3
  hashTableee.insert(23, 900); // Collision with key 3 and 13

  int Value;

  if (hashTableee.search(2, Value)) {
        std::cout << "Key 2 found with value: " << Value << std::endl;
    } else {
        std::cout << "Key 2 not found" << std::endl;
    }

    if (hashTableee.search(23, Value)) {
        std::cout << "Key 23 found with value: " << Value << std::endl;
    } else {
        std::cout << "Key 23 not found" << std::endl;
    }


  if(hashTableee.search(11, Value)){
    cout << "Key 11 found with value:  " << value << endl;
  } else {
    cout << "Key 11 not found" << endl;
  }

  hashTableee.remove(11);

  if(hashTableee.search(11, Value)){
    cout << "Key 11 found with value:  " << value << endl;
  } else {
    cout << "Key 11 not found" << endl;
  }


  return 0;
}
