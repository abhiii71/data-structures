#include <iostream> 
using namespace std;


// Write Functions to Traverse and Print Array Elements:
void PrintIntArr(int arr[], int size) {
  for (int i= 0; i < size; i++) {
cout << arr[i] << " ";
  }
  cout << endl;
}

void PrintFloatArr(float arr[], int size){
  for (int i = 0; i < size; i++){
    cout << arr[i] << " ";
  }
  cout << endl;
}

void PrintCharArr(char arr[], int size){
  for (int i = 0; i < size; i++){
    cout << arr[i] << " ";
  } 
  cout << endl;
}

// Implement Functions for Insertion and Deletion in an Array: 
void InsertElement(int arr[], int &size,int element, int position) {
  for (int i = size; i > position; i--) {
    arr[i] = arr[i-1];
    size++;
  }
}

void deleteElement(int arr[], int &size, int position) {
  for (int i = position; i < size-1; i++) {
    arr[i] = arr[i+1];
  }
  size --;
} 

int main(){

// Create Arrays of Different Data Types:
// Integer Array: 
int intArr[5] = {1,2,3,4,5};

// Float Array:
float floatArr[5] = {1.1,2.2,3.3};

// Characte array{string}
char charArr[3] = {'a','b','c'};

  PrintIntArr(intArr, 5);
  PrintFloatArr(floatArr, 5);
  PrintCharArr(charArr, 3);

  // Implement Functions for Insertion and Deletion in an Array:
  int size = 5;
  int intArray[10] = {1,2,3,4,5};

  InsertElement(intArray, size, 10,2);
  PrintIntArr(intArray,size);

  // Delete element at position 3:
  deleteElement(intArray,size,3);
  PrintIntArr(intArray,size);

}
