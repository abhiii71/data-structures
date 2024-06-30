#include <iostream>
#include <unordered_set>
using namespace std;

// Handling Duplicates and Unique Elements:
// Removing duplicates from a sorted array and returning new length
int removeDuplicates(int arr[], int n){
  if (n == 0 || n == 1)
    return n;

  int j = 0; // To store index of next unique element
  for(int i = 0; i < n-1; i++)
    if(arr[i] != arr[i+1])
      arr[j++] = arr[i];
  arr[j++] = arr[n-1];

    return j;
}


// Finding unique elements in an unsorted array using a set
void findUniqueElements(int arr[], int n){
  unordered_set<int> uniqueElements;
  for(int i = 0; i < n; i++){
    uniqueElements.insert(arr[i]);
  }
}



int main(){
  cout << " Questions: ";
    int sortedArr[] = {1, 1, 2, 2, 3, 4, 4, 5};
    int n = sizeof(sortedArr) / sizeof(sortedArr[0]);
    int newLength = removeDuplicates(sortedArr, n);

    cout << "Array after removing duplicates: ";
    for (int i = 0; i < newLength; i++) {
        cout << sortedArr[i] << " ";
    }
    cout << endl;
    cout << "New length: " << newLength << endl;

    // Example usage for findUniqueElements
    int unsortedArr[] = {4, 5, 1, 2, 2, 3, 4, 1};
    int m = sizeof(unsortedArr) / sizeof(unsortedArr[0]);
    findUniqueElements(unsortedArr, m);
  return 0;
}


