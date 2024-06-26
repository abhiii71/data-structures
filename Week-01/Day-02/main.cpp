#include <iostream>
#include <algorithm>

using namespace std;

// Linear Search:
int linear_Search(int arr[], int size, int element){
  for (int i = 0; i < size; i++){
    if (arr[i] == element){
      return 0;
    }
  }
  return -1;
}

// Binary Search
int binary_Search(int arr[], int size, int element){
  
  int start = 0, end = size-1;
 
  while (start <= end){
      int middle = start+(end-start)/2;

    if (arr[middle] == element) {
      cout << "Element found at index: " << middle << endl; 
      return middle; 
    }
    if (arr[middle] < element){
      start = middle + 1;
    } else {
        end = middle -1;
      }
    }
  cout << "Element not found! ";
  return  -1;
}

// Bubble Sort: 
void bubble_Sort(int arr[], int size){
  for (int i = 0; i < size-1; i++){
    for (int j = 0; j < size-i-1; j++){
      if (arr[j] > arr[j+1]){
        int temp = arr[j];
        arr[j] = arr[j+1];
        arr[j+1] = temp;
      }
    }
  }
}

// Selection Sort:
void selection_Sort(int arr[], int size){
  for (int i = 0; i < size-1; i++){
    int minIdx = i;
    for (int j = i+1; j < size; j++){
      if (arr[j] < arr[minIdx]) {
        minIdx = j;
      }
    }
    int temp = arr[minIdx];
    arr[minIdx]= arr[i];
    arr[i] = temp;
  }
}

// Insertion Sort:
void insertion_Sort(int arr[], int size){
  for(int i = 1; i < size; i++){
    int key = arr[i];
    int j = i-1;
    while (j >= 0 && arr[j] > key) {
      arr[j+1] = arr[j];
      j--;
    }
  arr[j+1] = key;
  }
}

// Printing Array: 
void PrintArr(int arr[], int size){
  for (int i = 0; i < size; i ++){
    cout << arr[i] << " ";
  }
  cout << endl;
}


int main(){
  
  int arr[5] = {1,2,3,4,5};
  int element = 3;

  
  // LinearSearch:
  cout << "Linear Search : found(0) || not found(-1): " << linear_Search(arr,5, element)<< "\n";

  // BinarySearch:
  cout << "Binary Search: " << binary_Search(arr,5,3) << endl;

  // Bubble Sort:
  int arr2[] = {71,21,7,69,3};
  int size = sizeof(arr2)/sizeof(arr2[0]);
  cout << "Original Array: ";
  PrintArr(arr2, size);
  bubble_Sort(arr2, size);
  cout << "After Bubble Sort: ";
  PrintArr(arr2, size);
  cout << endl;

  // Selection Sort:
  int arr3[] = {12,22,71,3,1,33}; 
  cout << "Original Array: ";
  size = sizeof(arr3)/sizeof(arr3[0]);
  PrintArr(arr3, size);
  selection_Sort(arr3, size); 
  cout << "After Selection Sort: ";
  PrintArr(arr3, size);
  cout << endl;

  // Insertion Sort: 
  int arr4[] = {12,22,71,3,1,33}; 
  cout << "Original Array: ";
  size = sizeof(arr4)/sizeof(arr4[0]);
  PrintArr(arr4, size);
  insertion_Sort(arr4, size); 
  cout << "After Insertion Sort: ";
  PrintArr(arr4, size);
  cout << endl;
}

