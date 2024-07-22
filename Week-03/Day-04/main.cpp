#include <clocale>
#include <cstddef>
#include <iostream>
#include <iterator>
using namespace std;

#define MAX 1000 //  ensuring that the priority queue can now hold up to 1000 elements.

// 1. Implementing Priority Queue using array: 
class PriorityQueue {
  int arr[MAX]; // array to store the elements of the queue
  int size; // current number of elements in the queue
  
public:

  PriorityQueue(): size(0) {}

  // The enqueue function adds a new number to our row of seats,
  // making sure to keep the numbers in order from smallest to largest.
  void enqueue(int value){
    if(value == MAX){
      cout << "Queue is full\n";
      return;
    }
    int i;
    for(i = size-1;(i >= 0 && arr[i] > value); i--){
      arr[i+1] = arr[i];
    }
    arr[i+1] = value;
    size++;
  }

  void dequeue(){
    if(size == 0){
      cout << "Queue is empty\n";
      return;
    }
    size--;
  }

  int front(){
    if(size == 0){
      cout << "Queue is empty\n";
      return -1;
    }
    return arr[0];
  }

  bool isEmpty(){
    return size == 0;
  }
};

// 2. Implementing a Priority Queue using Linked Lists:
struct Node{
  int data;
  int priority;
  Node* next;
  Node(int d, int p): data(d), priority(p), next(nullptr) {}
};

class PriorityQueueList {
  Node* front;

public:
  PriorityQueueList(): front(nullptr) {}

  void enqueue(int value, int priority){
    Node* temp = new Node(value, priority);
    if(!front || front->priority > priority){
      temp->next = front;
      front = temp;
    }else{
      Node* curr = front;
      while(curr->next && curr->next->priority <= priority){
        curr = curr->next;
      }
      temp->next = curr->next;
      curr->next = temp;
    } 
  }

  void dequeue() {
    if(!front){
      cout << "Queue is empty\n";
      return;
    }
    Node* temp = front;
    front = front->next;
    delete front;
  }

  int frontElement(){
    if(!front){
      cout << "Queue is empty\n";
      return -1;
    }
    return front->data;
  }

  bool isEmpty(){
    return (front == nullptr);
  }
};

// 3. Implementing a Deque using Arrays:
class Deque{
  int arr[MAX];
  int front, rear, size;

public:
  Deque(int s){
    front = -1;
    rear = 0;
    size = s;
  }

  bool isFull(){
    return ((front == 0 && rear == size-1) || front == rear+1);
  }

  bool isEmpty(){
    return (front == -1);
  }

  void insertFront(int value){
    if(isFull()){
      cout << "Deque is full\n";
      return;
    }
    if(front == -1){
      front = 0;
      rear = 0;
    } else if (front == 0){
      front = size-1;
    } else {
      front = front - 1;
    } 
    arr[front] = value;
  }

  void insertRear(int value){
    if (isFull()) {
      cout << "Dequeue is full\n";
      return;
    }
    if(front == rear){
      front = -1;
      rear = -1;
    } else if (front == size - 1){
      front = 0;
    } else {
      front = front + 1;
    }
  }

  void deleteFront(){
    if(isEmpty()){
      cout << "Queue is empty\n";
      return;
    }

    if(front == rear) {
      front = -1;
      rear = -1;
    } else if(front == size-1) {
      front = 0;
    } else {
      front = front + 1; 
    }
  }

  void deleteRear(){
    if(isEmpty()){
      cout << "Dequeue is empty\n";
      return;
    }

    if(front == rear){
      front = -1;
      rear = -1;
    } else if (rear == 0){
      rear = size - 1;
    } else {
      rear = rear - 1;
    }
  }

  int getFront(){
    if(isEmpty()){
      cout << "Dequeue is empty\n";
      return -1;
    }
    return arr[front];
  }

  int getRear(){
    if(isEmpty()){
      cout << "Dequeue is empty\n";
      return -1;
    }
    return arr[rear];
  }
};

// 4. Implementing a Deque using Linked Lists:

struct DequeNode {
  int data;
  DequeNode* prev;
  DequeNode* next;
  DequeNode(int val): data(val), prev(nullptr), next(nullptr) {}
};

class DequeLinkedList {
  DequeNode* front;
  DequeNode* rear;

public:
  DequeLinkedList():  front(nullptr), rear(nullptr) {}

  bool isEmpty(){
    return (front == nullptr);
  }

  void insertFront(int value){
    DequeNode* temp = new DequeNode(value);
    if(front == nullptr){
      front = rear = temp;
    } else {
      temp->next = front;
      front->prev = temp;
      front = temp;
    }
  }

  void insertRear(int value){
    DequeNode* temp = new DequeNode(value);
    if(rear == nullptr){
      front = rear = temp;
    } else {
      temp->prev = rear;
      rear->next = temp;
      rear = temp;
    }
  }

  void deleteFront(){
    if(isEmpty()){
      cout << "Deque is empty\n";
      return;
    }
    DequeNode* temp = front;
    front = front->next;
    if(front == nullptr){
      rear = nullptr;
    } else {
      front->prev = nullptr;
    }
    delete temp;
  }

  void deleteRear(){
    if(isEmpty()){
      cout << "Deque is empty\n";
      return;
    }
    DequeNode* temp = rear;
    rear = rear->prev;
    if(rear  == nullptr){
      front = nullptr;
    }else {
      rear->next = nullptr;
    }
    delete temp;
  }

  int getFront(){
    if(isEmpty()){
      cout << "Deque is empty\n";
      return -1;
    }
    return front->data;
  }

  int getRear(){
  if(isEmpty()){
      cout << "Deque is empty\n";
      return -1;
    }
    return rear->data;
  }
};

int main(){
  // 1. Implementing a Priority Queue using Arrays:
  cout << "Priority Queue Using Arrays:\n";
  PriorityQueue pq;
  pq.enqueue(21);
  pq.enqueue(71);
//  pq.enqueue(22);
//  pq.enqueue(33);
//  pq.enqueue(69);

  cout << "Front Element: " <<pq.front() << endl;
  
  pq.dequeue();
  pq.dequeue();

  cout << "Front Element: " << pq.front() << endl;

  // 2. Implementing a Priority Queue using Linked Lists:
  cout << "Priority Queue Using Linkedlist:\n";
  PriorityQueueList pql;
  pql.enqueue(33,1);
  pql.enqueue(12,3);
  pql.enqueue(14,2);

  cout << "Front element: " << pql.frontElement() << endl;

 // pql.dequeue();
 // pql.dequeue();
  cout << "Front element: " << pql.frontElement() << endl;

 // pql.dequeue();
//  pql.dequeue();

  cout << "Front Element: " << pql.frontElement() << endl;

  // 3. Implementing a Deque using Arrays:
  cout << "Implement Dequeue Using Arrays:\n";
  Deque dq(5);
  dq.insertFront(223);
  dq.insertFront(749);
  dq.insertRear(334);
  dq.insertRear(456);

  cout << "Front Element: " << dq.getFront() << endl;
  cout << "Rear Element: " << dq.getRear() << endl;

  dq.deleteRear();
  dq.deleteFront();
   
  cout << "Front Element: " << dq.getFront() << endl;
  cout << "Rear Element: " << dq.getRear() << endl;

  // 4. Implementing a Deque using Linked Lists:
  cout << "Implement Deque using Linked List:\n";

  DequeLinkedList dql;
  dql.insertRear(10);
  dql.insertRear(20);
  dql.insertFront(30);
  dql.insertRear(40);
  cout << "Front element: " << dql.getFront() << endl;
  cout << "Rear Elemet" << dql.getRear() << endl;

  dql.deleteRear();
  dql.deleteFront();

  cout << "Front element: " << dql.getFront() << endl;
  cout << "Rear Elemet" << dql.getRear() << endl;

  

  return 0;
}

