#include <algorithm>
#include <cctype>
#include <iostream>
#include <stack>
#include <string>
using namespace std;

// 1. Converting Infix Expressions to Postfix and Prefix:
// a. Infix to Postfix:
int precedence(char op) { // to determine  precedence  of operators
  if(op == '+' || op == '-') return 1;
  if(op == '*' || op == '/') return 2;
  if(op == '^') return 3;
  return 0;
}

// function to convert infix to postfix
string infixToPostfix(string infix){
  stack<char> s; // stack to hold operators
  string postfix = ""; // resulting postfix expression 
  
  for(char ch: infix){ // process each character in the infix expression  
    if(isalnum(ch)) { // if the character  is an operand (letter/number)
      postfix += ch; // Add it directly to the postfix expression 
    } else if (ch == '('){ // if the charcter is an opening parenthesis 
      s.push(ch); // push it onto the stack 
    } else if(ch == ')') { // if the character is a closing parenthesis
      // Pop from the stack to the postfix expression  until an  opening parenthesis is encountered
       while (!s.empty() && s.top() !=  '(') {
        postfix += s.top();
        s.pop();
      }
      s.pop(); // Remove the opening parenthesis from the stack 
    } else { // if the character  is an operator 
      // Po//pop from the stack to the postfix  expression  based on operator precedence 
      while (!s.empty() && precedence(s.top()) >= precedence(ch)) {
        postfix += s.top();
        s.pop();
      }
      s.push(ch); // push the current operator onto the stack
    }
  }

  // Pop all remaining operators from the stack to the postfix  expression
  while(!s.empty()){
    postfix += s.top();
    s.pop();
  }
  return postfix; // Return the resulting postfix expression
}

// b. Infix to Prefix:
string infixToPrefix(string infix) {
  reverse(infix.begin(), infix.end());
  for (int i = 0; i < infix.size();++i) {
    if(infix[i] == '(') infix[i] = ')';
    else if(infix[i] == ')') infix[i] = '(';
  } 
  string postfix = infixToPrefix(infix);
  reverse(postfix.begin(), postfix.end());
  return postfix;
}

// 2. Evaluating Postfix and Prefix Expressions:
// a. Postfix Evaluation:
int evalutatePostfix(string postfix) {
  stack<int> s;
  for(char ch: postfix) {
    if(isdigit(ch)) {
      s.push(ch - '0');
    } else {
      int val1 = s.top(); s.pop();
      int val2 = s.top(); s.pop();
      switch (ch) {
        case '+': s.push(val2 + val1); break;
        case '-': s.push(val2 - val1); break;
        case '*': s.push(val2 * val1); break;
        case '/': s.push(val2 / val1); break;
      }
    }
  }
  return s.top();
}

// b. Prefix Evaluation:
int evaluatePrefix(string prefix) {
  stack<int> s;
  for(int i = prefix.size() - 1; i >= 0; --i) {
    if(isdigit(prefix[i])) {
      s.push(prefix[i] - '0');
    } else {
      int val1 = s.top(); s.pop();
      int val2 = s.top(); s.pop();
      switch (prefix[i]) {
        case '+': s.push(val1+val2); break;
        case '-': s.push(val1-val2); break;
        case '*': s.push(val1*val2); break;
        case '/': s.push(val1/val2); break;
      }
    }
  }
  return s.top();
}

// 3. Implement a Stack to Reverse a String:
string reverseString(string str) {
  stack<char> s;
  for(char ch: str){
    s.push(ch);
  }
  string reversed = "";
  while (!s.empty()) {
    reversed += s.top();
    s.pop();
  }
  return reversed;
}

int main() {

  // Converting Infix Expressions to Postfix and Prefix:
  // a.
  string infix = "A+B*C-D/E";
  cout << "Postfix: " << infixToPostfix(infix) << endl;

  // b. 
  string infix1 = "A+B*C-D/E";
  cout << "Prefix: " << infixToPrefix(infix1) << endl;

  // 2. Evaluating Postfix and Prefix Expressions:
  // a. Postfix Evaluation:
  string postfix = "231*+9-";
  cout << "Postfix Evaluation: " << evalutatePostfix(postfix) << endl;
  
  // b. Prefix Evaluation:
  string prefix = "-+*2319";
  cout << "Prefix Evaluation: " << evaluatePrefix(prefix) << endl;

  // 3. Implement a Stack to Reverse a String:

  string str = "Hello, World!";
  cout << "Reversed String: " << reverseString(str); 
  return 0;
}
