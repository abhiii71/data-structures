#include <iostream>
#include <vector>
#include <cmath>
#include <list>
using namespace std;

// Modular Hashing:
size_t modularHash(int key, size_t tableSize){
  return key % tableSize;
}

// Multiplication Method:
size_t multiplicationHash(int key, size_t tableSize, double A){
  return floor(tableSize * (key * A - floor(key * A)));
}

// Universal Hashing:
size_t universalHash(int key, size_t tableSize, int a, int b, int p){
  return ((a * key + b) % p) % tableSize;
}

void testHashFunction(){
  const size_t tableSize = 10;
  const double A = 0.6180339887; // (sqrt(5)-1)/2
  const int p = 31; // Prime number for universal hashing
  const int a = 1 + rand() % (p-1); // Random a in [1, p-1]
  const int b = rand() % p; // Random b in [0, p-1]

  vector<int> keys  = {1,2,12,22, 32, 42, 52, 62, 72, 82};

  cout << "Modular Hashing:\n";
  for(int key: keys){
    cout << "Key: " << key << "Hash: " << modularHash(key, tableSize) << endl; 
  }

  cout << "Multiplication Hashing:\n";
  for(int key: keys){
    cout << "Key: " << key << "Hash: " << multiplicationHash(key, tableSize, A) << endl;
  }

  cout << "Universal Hashing:\n";
  for(int key: keys){
    cout << "Key: " << key << "Hash: " << universalHash(key, tableSize, a, b, p) << endl;
  }
}

// Problem 1: Design a Custom Hash Function
size_t customStringHash(const string& str, size_t tableSize){
  size_t hash = 0;
  size_t prime = 31; // A small prime number 
  for(char c: str){
    hash =(hash * prime + c) % tableSize;
  }
  return hash;
}

void testCustomStringHash(){
  const size_t tableSize = 10;
  vector<string> keys = {"apple", "banana", "cherry", "date", "elderberry"};

  cout << "Custom String Hashing: \n";
  for(const string& key: keys) {
    cout << "Key: " << key << "Hash: " << customStringHash(key, tableSize) << endl;
  }
}

// Problem 2: Hash Table with Custom Hash Function
class CustomHashTable {
private:
  vector<list<pair<string, int>>> table;
  size_t tableSize;
  size_t customStringHash(const string& str){
    size_t hash = 0;
    size_t prime = 31; // A small prime number  
    for(char c: str){
      hash = (hash*prime+c) % tableSize;
    }
    return hash;
  }

public:
  CustomHashTable(size_t size): tableSize(size){
    table.resize(tableSize);
  }

  void insert(const string& key, int value){
    size_t index = customStringHash(key);
    auto& bucket = table[index];
    for(auto& pair: bucket){
      if(pair.first == key){
        pair.second = value; // Update existing value 
        return;
      }
    }
    bucket.emplace_back(key, value); // Insert new key value pair
  }
  bool search(const string& key, int& value){
    size_t index = customStringHash(key);
    auto& bucket = table[index];
    for(const auto& pair: bucket){
      if(pair.first == key){
        value = pair.second;
        return true;
      }
    }
    return false; // Key  not found
  }

  void remove(const string& key) {
    size_t index = customStringHash(key);
    auto& bucket = table[index];
    for(auto it = bucket.begin(); it != bucket.end(); ++it){
      if(it->first == key){
        bucket.erase(it);// Remove key-value pair
        return;
      }
    }
  }
};

void testCustomHashTable(){
  CustomHashTable hashTable(10); // create hash table with 10 slots

  hashTable.insert("apple", 100);
  hashTable.insert("banana", 200);
  hashTable.insert("cherry", 300); // Collision with key "banana"

  int value;
  if(hashTable.search("banana", value)){
    cout << "Key 'banana' found with value: " << value << endl;
  } else {
    cout << "Key 'banana' not found!";
  }

  hashTable.remove("banana");
  if(hashTable.search("banana", value)){
    cout << "Key 'banana' found with value: " << value << endl;
  } else {
    cout << "Key 'banana' not found" << endl;
  }
}

int main(){
  
  srand(time(0));
  
  testHashFunction();

  // Problem 1: Design a Custom Hash Function
  testCustomStringHash();


  // Problem 2: Hash Table with Custom Hash Function
  testCustomHashTable();

  return 0;
}
