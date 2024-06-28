#include <algorithm>
#include <iostream>
#include <utility>
using namespace std;


// Merge sort:
void merge(int arr[], int left, int mid, int right){
  int n1 = mid - left + 1;
  int n2 = right - mid;

  // create temp arrays 
  int* L = new int[n1];
  int* R = new int[n2];

  // Copy data to temp arrays 
  for(int i = 0; i < n1; i++)
    L[i] = arr[left+i];
  for(int j = 0; j < n2; j++)
    R[j] = arr[mid+1+j];

  // Merge the temp arrays back into arr[left..right]
  int i = 0, j = 0, k = left;
  while(i < n1 && j < n2){
    if (L[i] <= R[j]){
      arr[k] = L[i];
      ++i;
    } else {
      arr[k] = R[j];
      ++j;
    }
    ++k;
  }

  // Copy the remaining elements of L[], if any 
  while (i < n1) {
    arr[k] =L[i];
    ++i;
    ++k;
  }

  // Copy the remaining elements of R[], if any 
  while(j < n2) {
    arr[k] = R[j];
    ++j;
    ++k;
  }

  // Clean up dynamically allocated memory 
  delete[] L;
  delete[] R;
}


void mergesort(int arr[], int left, int right){
  if (left < right){
    int mid = left + (right - left)/2;

    // sort first and second halves 
    mergesort(arr, left, mid);
    mergesort(arr, mid+1 , right);

    // Merge the sorted halves
    merge(arr,left, mid, right);
  }
}

// Quick Sort:
int  partition(int arr[], int low, int high){
  int pivot = arr[high]; // choosing the last element as pivot 
  int i = (low-1); // index of smaller element 


  for (int j = low; j <= high - 1; ++j){
    if (arr[j]< pivot){
      ++i;
      swap(arr[i], arr[j]);
    }
  }
  swap(arr[i+1], arr[high]);
  return (i+1);
}

void quicksort(int arr[], int low, int high) {
  if(low < high){
    int pi = partition(arr, low, high); // pi is partitioning index 

    // Recursively sort elements before and after partition
    quicksort(arr, low, pi-1);
    quicksort(arr, pi + 1, high);
  }
}


int main(){

  // Merge sort: 
  int arr[] = {71,69,7,21,33};
  int arr_size = sizeof(arr)/ sizeof(arr[0]);
  cout << "Given array: \n";
  for (int i = 0; i < arr_size; i++)
    cout << arr[i] << " ";
  cout << endl;

  mergesort(arr, 0, arr_size-1);
  
  cout << "After Merge sort: \n";
  for (int i = 0; i < arr_size; i++)
    cout << arr[i] << " ";
  cout << endl;


  // Quick Sort:
     int arr2[]  = {10, 7, 8, 9, 1, 5};
    int n = sizeof(arr2) / sizeof(arr2[0]);
    
    cout << "Given array is: ";
    for (int i = 0; i < n; ++i)
        cout << arr2[i] << " ";
    cout << endl;
    
    quicksort(arr2, 0, n - 1);
    
    cout << "Sorted array is: ";
    for (int i = 0; i < n; ++i)
        cout << arr2[i] << " ";
    cout << endl;

  return 0;

}
