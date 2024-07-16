#include <iostream>
#include <stack>
using namespace std;


// Implement stack using arrays
class StackArray {
  int* arr;
  int top;
  int capacity;

public:
  StackArray(int size) {
    arr = new int[size];
    capacity = size;
    top = -1;
  }

  void push(int x){
    if(top == capacity - 1) {
      cout << " capacity full";
      return;
    }
    arr[++top] = x;
  }

  int pop(){
    if(top == -1){
      cout << "Stack underflow";
      return -1;
    }
    return arr[top--];
  }

  int peek(){
    if(top == -1){
      cout <<" stack is empty";
      return -1;
    }
    return arr[top];
  }

  bool isEmpty(){
    return top == -1;
  }
};


// Implement stack using LinkedList
class StackNode {
public:
  int data;
  StackNode* next;
  StackNode(int data): data(data), next(nullptr) {}
};

class StackLinkedList {
  StackNode* top;

public:
  StackLinkedList(): top(nullptr) {}


  void push(int x){
    StackNode* newNode = new StackNode(x);
    newNode->next = top;
    top = newNode;
  }


  int pop(){
    if(!top) {
      cout << "Stack Underflow\n";
      return -1;
    }
    int popped = top->data;
    StackNode* temp = top;
    top = top->next;
    delete temp;
    return popped;
  }


  int peek(){
    if(!top) {
      cout << "Stack is Empty\n";
      return -1;
    } 
    return top->data;
  }

  bool isEmpty(){
    return !top;
  }
};

// Check for Balanced Parentheses 
bool areParenthesesBalanced(string expr) {
  stack<char> s;
  for(char ch: expr) {
    if(ch == '(' || ch == '{' || ch == '[' ) {
      s.push(ch);
    } else {
      if(s.empty()) return false;
      char top = s.top();
      if((ch == ')' && top == '(') || (ch == '}' && top == '{')  || (ch == ']' && top == '[')){
        s.pop();
      } else {
        return false;
      }
    }
  }
  return s.empty();
}

int main(){
 
  // over array
  StackArray stack(5);
  stack.push(71);
  stack.push(22);
  stack.push(69);

  cout << "Peek element: " << stack.peek() << endl; 
  stack.pop();
  cout << "One item popped out!\n";
  cout << " Peek element after last element pop: "<< stack.peek() << endl;
  cout << "Is stack empty: "<< stack.isEmpty();
  
  // over Linked  list
  cout << " Stack over Linked List:\n";
  StackLinkedList stackN;
  stackN.push(13);
  stackN.push(33);
  stackN.push(17);
  cout << " Peek element: "<< stackN.peek() << endl;
  stackN.pop();
  cout << "Last element Removed!\n";
  cout << "Peek element: "<< stackN.peek() << endl;
  if (stackN.isEmpty()){
    cout << "stack is empty\n";
  } else {
    cout << "Stack is not empty\n";
  }

  // check parentheses
  cout << "Check Parentheses: \n";
  string expr  = "{()}[]";
  if(areParenthesesBalanced(expr)){
    cout << "Balanced\n";
  } else {
    cout << "Not Balanced\n";
  }

  return 0;
}
