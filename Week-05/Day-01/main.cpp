#include <iostream>
#include <queue>
using namespace std;

// In-Order Traversal:
struct TreeNode {
  int val;
  TreeNode* left;
  TreeNode* right;
  TreeNode(int x): val(x), left(nullptr), right(nullptr) {}
};

//In-Order Traversal:
void inOrder(TreeNode* root){
  if(root == nullptr) return;

  inOrder(root->left);
  cout << root->val << " ";
  inOrder(root->right);
}

//Pre-Order Traversal:
void preOrder(TreeNode* root){
  if(root == nullptr) return;

  cout << root->val << " ";
  preOrder(root->left);
  preOrder(root->right);
}

// Post-Order Traversal:
void postOrder(TreeNode* root){
  if(root == nullptr) return;

  postOrder(root->left);
  postOrder(root->right);
  cout << root->val << " ";
}

// Level-Order Traversal:
void levelOrder(TreeNode* root){
  if(root == nullptr) return;

  queue<TreeNode*> q;
  q.push(root);

  while(!q.empty()){
    TreeNode* node = q.front();
    q.pop();

    cout << node->val << " ";

    if(node->left) q.push(node->left);
    if(node->right) q.push(node->right);
  }


}
int main(){
  TreeNode* root = new TreeNode(1);
  root->left = new TreeNode(2);
  root->right = new TreeNode(3);
  root->left->left = new TreeNode(4);
  root->left->right = new TreeNode(5);

  cout << "In-Order Traversal: \n";
  inOrder(root);
  cout << endl;

  cout << "Pre-Order Traversal: \n";
  preOrder(root);
  cout << endl;

  cout << "Post-Order Traversal: \n";
  postOrder(root);
  cout << endl;

  cout << "Level-Order Traversal: \n";
  levelOrder(root);
  cout << endl;



}
