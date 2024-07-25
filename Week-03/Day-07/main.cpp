#include <iostream>
#include <stack.h>
#include <queue.h>
using namespace std;

void stackOp(){
  stack<int> s;
  s.push(7);
  s.push(10);
  
  cout << s.top() << endl;
  cout << s.top() << endl;
  s.pop();

  s.push(3);
  s.push(5);
  cout << s.top() << endl;
  s.pop();

  cout << s.size() << endl;
  cout << s.top() << endl;
  s.push(8);
  cout << s.top() << endl;
  s.pop();
  cout << s.top() << endl;
  s.pop();

}

// reverse the queue
void reverseQ(){
  queue<int> q;
  stack<int> s;

  q.enqueue(1);
  q.enqueue(2);
  q.enqueue(3);

  while(!q.IsEmpty()) {
    s.push(q.dequeue());
  }

  while(!s.empty()){
    q.enqueue(s.pop());
  }
  cout << q << endl;

}

int main(){

  stackOp();
  // reverseQueue:
  reverseQ();
  return 0;

}
