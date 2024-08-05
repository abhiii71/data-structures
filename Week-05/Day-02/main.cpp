#include <iostream>
#include <algorithm>
#include <queue> 

using namespace std; 

struct TreeNode {
  int val;
  TreeNode* left;
  TreeNode* right;
  TreeNode(int x): val(x), left(nullptr), right(nullptr) {}
};

// Height of a Tree:
int calculateHeight(TreeNode* root){
  if(root == nullptr) return -1; // retunrn -1 for height to count edges
  int leftHeight = calculateHeight(root->left);
  int rightHeight = calculateHeight(root->right);
  return max(leftHeight, rightHeight) + 1;
}

// Depth of a Node:
int calculateDepth(TreeNode* root, TreeNode* target){
  if(root == nullptr) return -1; // return -1 if target is not found 
  if(root == target) return 0; // depth of the node is 0
  
  int leftDepth = calculateDepth(root->left, target);
  if(leftDepth != -1) return leftDepth + 1;

  int rightDepth = calculateDepth(root->right, target);
  if(rightDepth != -1) return rightDepth + 1;

  return -1;// target not found
}

// Rotations in AVL Trees:
// Right rotate
TreeNode* rightRotate(TreeNode* y){
  TreeNode* x = y->left;
  TreeNode* T2 = x->right;

  x->left = y;
  y->left = T2;

  return x;
}
// Left rotate
TreeNode* leftRotate(TreeNode* x){
  TreeNode* y = x->right;
  TreeNode* T2 = y->left;

  y->left = x;
  x->right = T2;

  return y;
}

// Get the balance factor for AVL tree 
int getBalance(TreeNode* N){
  if(N == nullptr) return 0;
  return calculateHeight(N->left) - calculateHeight(N->right);
}

// Insert function for AVL tree
TreeNode* insertAVL(TreeNode* node, int key){
  if(node == nullptr) return new TreeNode(key);

  if(key < node->val) node->left = insertAVL(node->left, key);
  else if(key > node->val) node->right = insertAVL(node->right, key);
  else return node;

  int balance = getBalance(node);

  // left left case 
  if(balance > 1 && key < node->left->val) return rightRotate(node);

  // right right case 
  if(balance < -1 && key > node->right->val) return leftRotate(node);

  // left right case 
  if(balance > 1 && key > node->left->val) {
    node->left = leftRotate(node->left);
    return rightRotate(node);
  }
   
  // Right left case 
  if(balance < -1 && key < node->right->val) {
    node->right = rightRotate(node->right);
    return leftRotate(node);
  } 
  return node;
}

// In-Order traversal function to display the tree 
void inOrder(TreeNode* root){
  if(root != nullptr){
    inOrder(root->left);
    cout << root->val << " ";
    inOrder(root->right);
  }
}

int main(){
  TreeNode* root = new TreeNode(1);
  root->left = new TreeNode(2);
  root->right = new TreeNode(3);
  root->left->left = new TreeNode(4);
  root->left->right = new TreeNode(5); 


  cout << "Height of the tree: " << calculateHeight(root) << endl;


  cout << "Depth of a Node: \n";
  
  // specify the target node 
  TreeNode* target = root->left->right; //  Node with value 5
  
  // Calculate the depth of the target node
  int depth = calculateDepth(root, target); 

  // Output the depth of the target node 
  if(depth != -1){
    cout << "The depth of the target node with value " << target->val << " is: " << depth << endl;
  } else {
    cout << "The target node with value " << target->val << "was not found in the tree." << endl;
  }

  // Rotations in AVL Trees:
  TreeNode* avlroot = nullptr;
  avlroot = insertAVL(avlroot, 10);
  avlroot = insertAVL(avlroot, 20);
  avlroot = insertAVL(avlroot, 30);
  avlroot = insertAVL(avlroot, 40);
  avlroot = insertAVL(avlroot, 50);
  avlroot = insertAVL(avlroot, 25);

  cout << "In-Order traversal of the constructed AVL tree is: \n";
  inOrder(root);
  cout << endl;

  // clean up allocated memory (for completeness)
  delete root->left->left; 
  delete root->left->right;
  delete root->left;
  delete root->right;
  delete root;

  return 0;

}
