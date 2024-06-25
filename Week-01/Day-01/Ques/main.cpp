#include <iostream>
using namespace  std;



// Traversal and Print: 
void PrintArr(int arr[], int size){
  for (int i = 0; i < size; i++){
    cout << arr[i] << " ";
  }
  cout << endl;
}

// Insert an Element at a Specific Position:
void InsertAtPosition(int arr[], int &size, int position, int element,int capacity) { // &size = passed by reference (means value reflect outside of the function)
  if (position < 0 || position > size || size >= capacity){ 
    cout << "Invalid position or array is full" << endl;
    return;
  }
  // shift elements to the right from the end of the position
  for (int i = size; i > position; i--) {
    arr[i] = arr[i-1];
  }

  // Insert the new element 
  arr[position] = element;
  size++;
}

// Delete an Element from a Specific Position:
void DeleteAtPosition (int arr[],int &size, int position, int element, int capacity) {
  if (position < 0 || position > size || size >= capacity) {
    cout << "Invalid position or array is full" << endl;
    return;
  }

  for (int i = size; i > position-1; i++){
arr[position] = arr[position+1];
  }
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

// [1, 2, 3, 4, 5]
// [3, 4, 5, 1, 2]
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



// [1, 2, 3, 4, 5]
// [4, 5, 1, 2, 3]
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

// Search in a Rotated Sorted Array:
int Search(int arr[], int &size, int element){
  for (int i = 0; i < size; i++){
    if (arr[i] == element){
      return 0; 
    } 
  }
  return -1;
}

// Delete from specific:
void Delete(int arr[], int &size, int index){
  for (int i = index; i < size-1; i++){
    arr[i] = arr[i+1];
  }
  size--;
}




int main(){
  cout << "Day -1 Array Questions: \n";

  // Traversal and Print: 
  int intArr[6] = {1,2,3,4,5};

  int size = 5;

  PrintArr(intArr,size);

  // Insert an Element at a Specific Position:
  int position = 3;
  int element = 71;
  int capacity = 6;
  InsertAtPosition(intArr, size, position, element,capacity);
  PrintArr(intArr, size);
  cout << "size increated: "<< size << "\n";

  // min max:

  int arr[10] = {71,69,22,7,33};
  size = 5;
  minMaxArr(arr,size);

  // Rotation:
  cout << "Original Array: ";
  PrintArr(arr,size);

  int n = 5; 
  int k = 2;
  // Left Rotation:
  cout << "Left Rotated Array: ";
  leftRotate(arr, n, k);
  PrintArr(arr,size);

  //  Right Rotation:

  cout << "Original Array: ";
  PrintArr(arr,size);
  cout << "Right Rotation: ";
  rightRotate(arr,n,k);
  PrintArr(arr, size);

  // Search in a Rotated Sorted Array:
  element = 7;
  cout <<"Is element found in the Array: " << Search(arr, size, element);
}


