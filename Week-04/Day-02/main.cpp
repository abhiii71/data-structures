#include <iostream>
#include <vector>
#include <cmath>
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

int main(){
  
  srand(time(0));
  
  testHashFunction();
  
  return 0;
}
