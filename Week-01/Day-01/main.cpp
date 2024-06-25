#include <algorithm>
#include <cstdio>
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
void InsertElement(int arr[], int &size,int capacity,int element, int position) {
 
  if (position < 0 || position > size) {
    cout << "Invalid position!" << endl;
    return;
  }
  if (size >= capacity){
    cout << "Array is at full capacity!" << endl;
    return;
  }
// shift element to the right from the end to the position 
  for (int i = size; i > position; i--) {
    arr[i] = arr[i-1];
  }

  // Insert the new element at the specified position
arr[position] = element;
  size++;
}

void deleteElement(int arr[], int &size, int position) {
  if (position < 0 || position >= size) {
    cout << "Invalid position!" << endl;
    return ;
  }
// shift elements to the left from the position 
  for (int i = position; i < size-1; i++) {
    arr[i] = arr[i+1];
  }

  // Reduce the size of the array 
  size --;
}

// Find Maximum and Minimum Element in an Array: 
void minMaxArr(int  arr[], int &size){
  int min = arr[0];
  // For min: 
  for (int i = 0; i < size; i++){
    if (min > arr[i]){
      min = arr[i];
    }
  }
  // For max:
  int max = arr[0];
  for (int i = 0; i < size; i++){
    if (max < arr[i]){
      max = arr[i];
    }
  }
  cout << "Min element in Array: "<<  min << "\n";
  cout << "Max element in Array: "<<  max << "\n";
}

// Rotate an Array (Left or Right):
void leftRotate(int arr[], int n, int k){
  
  // To handle cases where k is greater than n 
  k = k%n;


  // Temporary array stored elements to be rotated
  int temp[k]; 

  // Copy first k element to the temp array:
  for (int i = 0; i < k; i++){
    temp[i] = arr[i];
  }

  //shift remaining elements of arr[] to the left 
  for (int i = k; i < n; i++){
    arr[i-k] = arr[i];
  }

  // Copy elements from temp[] back to arr[] to complete rotation
  for (int i = 0; i < k; i++){
    arr[n-k+i] = temp[i];
  }
}




void rightRotate(int arr[], int n, int k){
  k = k%n;

  int temp[k];

  for (int i = n-k; i < n ;i++){
    temp[i-(n-k)] = arr[i];
 }
  for(int i= n-1; i >= k; i--){
    arr[i] = arr[i-k];
  }

  for (int i = 0; i < k;i++){
    arr[i] = temp[i];
  }
}















int main(){

// Create Arrays of Different Data Types:
// Integer Array: 
int intArr[5] = {1,2,3,4,5};

// Float Array:
float floatArr[3] = {1.1,2.2,3.3};

// Characte array{string}
char charArr[3] = {'a','b','c'};

  PrintIntArr(intArr, 5);
  PrintFloatArr(floatArr, 5);
  PrintCharArr(charArr, 3);

  // Implement Functions for Insertion and Deletion in an Array:
  int size = 5;
  int capacity = 10;
  int intArray[10] = {1,2,3,4,5};

  InsertElement(intArray, size,capacity, 10,2);
  PrintIntArr(intArray,size);

  // Delete element at position 3:
  deleteElement(intArray,size,3);
  PrintIntArr(intArray,size);

  // min max:

  int arr[10] = {71,69,22,7,33};
  size = 5;
  minMaxArr(arr,size);

  // Rotation:
  cout << "Original Array: ";
  PrintIntArr(arr,size);

  int n = 5; 
  int k = 2;
  // Left Rotation:
  cout << "Left Rotated Array: ";
  leftRotate(arr, n, k);
  PrintIntArr(arr,size);

  //  Right Rotation:

  cout << "Original Array: ";
  PrintIntArr(arr,size);
  cout << "Right Rotation: ";
  rightRotate(arr,n,k);
  PrintIntArr(arr, size);
}
