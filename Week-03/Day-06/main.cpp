#include <iostream>
#include <stack>
#include <deque> //  Includes the double-ended queue (deque) library, which provides the deque data structure.
#include <vector> //  Includes the vector library, which provides the vector data structure.
#include <climits>
using namespace std;

// 1. Implementing a Min Stack (Stack with O(1) Minimum Element Retrieval): 
class MinStack{
  stack<int>s;
  stack<int>minS;

public:

  void push(int x){
    s.push(x);
    if(minS.empty() || x < minS.top()){ // If x is less than or equal to the current minimum 
    // (the top of minS), push x onto minS. This ensures that minS always has the minimum element at the top
      minS.push(x);
    }
  }

  void pop(){
    if(s.top() == minS.top()){
      minS.pop();
    }
    s.pop();
  }

  int top(){
    return s.top(); 
  }

  //getMin
  int getMin(){
    return minS.top();
  }


};

// 2. Implementing a Sliding Window Maximum Using Deque:
vector<int> maxSlidingWindow(vector<int>& nums, int k){
  deque<int> dq;
  vector<int> result;

  for(int i = 0; i <nums.size(); i++){
    if(!dq.empty() && dq.front() == i-k){
      dq.pop_front();
    }

    while(!dq.empty() && nums[dq.back()] <= nums[i]){
      dq.pop_back();
    }
    dq.push_back(i);
    if(i >= k-1){
      result.push_back(nums[dq.front()]);
    }
  }
  return result;
}

int main(){


  // 1. Implementing a Min Stack (Stack with O(1) Minimum Element Retrieval):
  MinStack minStack;
  minStack.push(-2);
  minStack.push(0);
  minStack.push(-3);

  cout << "Min element: " << minStack.getMin() << endl;
  
  minStack.pop();

  cout << "Top Element: "<< minStack.top() << endl;
  cout << "Minimum Element: " << minStack.getMin() << endl;

  // 2. Implementing a Sliding Window Maximum Using Deque:

  vector<int> nums = {1, 3, -1, -3, 5, 3, 6, 7};
  int k = 3;
  vector<int> result = maxSlidingWindow(nums, k);

  cout << "Sliding window maximums: ";
  for(int num: result){
    cout << num << " ";
  }
  cout << endl;

  return 0;
}


