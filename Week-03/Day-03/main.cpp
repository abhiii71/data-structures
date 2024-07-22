#include <iostream>
#include <iterator>
#define MAX 1000
using namespace std;

// 1. Implementing a Queue using Arrays:
class Queue {
  int front,rear;
  int size;
  int* arr;

public:
  Queue(int s) {
    front = rear = 0;
    size  = s;
    arr = new int[s];
  }

  void enqueue(int value){
    if(rear == size){
      cout << "Queue is full\n";
      return;
    }
    arr[rear++] = value;
  }

  void dequeue(){
    if(front == rear){
      cout << "Queue is empty\n";
      return;
    }
    for(int i = 0; i < rear -1; i++){
      arr[i] = arr[i+1];
    }
    rear--;
  }

  int frontElement(){
    if(front == rear){
      cout << "Queue is empty\n";
      return -1;
    }
    return arr[front];
  }

  bool isEmpty(){
    return (front == rear);
  }
};

int main() {
  Queue q(5);
  
  q.enqueue(21);
  q.enqueue(71);
  q.enqueue(33);
  q.enqueue(22);

  q.dequeue();
  cout << "Front Element: " << q.frontElement() << endl;
  q.dequeue();
  q.dequeue();
  q.dequeue();
  cout << "Is Queue Empty: " << (q.isEmpty() ? "Yes":"No") << endl;
  return 0;
}
