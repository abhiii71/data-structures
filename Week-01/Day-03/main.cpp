#include <algorithm>
#include <iostream>
using namespace std;


/*
Matrix Rotations:
90 Degrees Clockwise Rotation
*/ 
void rotateMatrixClockwise(int matrix[3][3]) {
// Transpose matrix 
  for(int i = 0; i < 3; i++){
    for(int j = 0; j < 3; j++){
      int temp = matrix[i][j];
      matrix[i][j] = matrix[j][i];
      matrix[j][i] = temp;
      }
    }

// Now rotate the matix 

for (int i = 0; i < 3; i++){
int start = 0;
int end = 2;
while (start < end) {
int temp = matrix[i][start];
matrix[i][start] = matrix[i][end];
matrix[i][end] = temp;
start++;
end--;
}
}
}

void PrintMatrix(int matrix[3][3]){
for (int i = 0; i < 3; i++){
for (int j = 0; j < 3; j++){
cout << matrix[i][j] << " "; 
}
cout <<endl;
}
}
int main(){
// 2D array declaration
int matrix[3][3] = {
{1,2,3},
{4,5,6},
{7,8,9},
};

// Accessing elements  Row-wise Traversal:
cout << "Row wise traversal: \n";
for(int i = 0; i < 3; i++){
for (int j = 0; j < 3; j++){
cout << matrix[i][j] << " "; // Row-wise Traversal:
}
cout << endl;
} 
// Accessing elements  Col-wise Traversal:
cout << "Col wise traversal: \n";
for(int i = 0; i < 3; i++){
for (int j = 0; j < 3; j++){
cout << matrix[j][i] << " ";
}
cout << endl;
}
cout << "Original Matrix: \n";
PrintMatrix(matrix);

cout << "Matrix after rotated Clockwise: \n";
rotateMatrixClockwise(matrix);
PrintMatrix(matrix);

  return 0;
}

