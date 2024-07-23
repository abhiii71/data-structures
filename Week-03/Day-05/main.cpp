#include <iostream>
#include <stack>
#include <queue> //  Includes the queue library for using the queue data structure.
#include <vector>
using namespace std;


// 1. Implementing DFS using Stacks:
void DFS(vector<vector<int>>& graph, int start){
// Defines a function named DFS that takes a reference to a 2D vector graph 
// and an integer start as parameters. The function does not return any value (void).
	vector<bool>  visited(graph.size(), false);
	// vector<bool> visited(graph.size(), false);: Creates a vector of boolean values named visited 
	// with the same size as the number of nodes in the graph. Initially, all values are set to false,
	// indicating that no nodes have been visited.

	stack<int> s;
	// stack<int> s;: Creates a stack of integers named s.
	
	s.push(start);
	// s.push(start);: Pushes the starting node onto the stack.

	while(!s.empty()) { // Continues the loop as long as the stack is not empty.
		int node = s.top(); // Retrieves the node at the top of the stack without removing it.
		s.pop(); // Removes the node from the top of the stack.

		if(!visited[node]) { // Checks if the node has not been visited.
			cout << node << " "; //  Prints the current node to the console.
			visited[node] = true; // Marks the node as visited.
		}

		for(int i = graph[node].size() - 1; i >= 0; i--){ 
		//  Iterates over the adjacency list of the current node in reverse order. 
		//  This ensures that nodes are visited in the correct order when they are pushed 
		//  onto the stack.


			if(!visited[graph[node][i]]) { // Checks if the adjacent node has not been visited
				s.push(graph[node][i]);// Pushes the adjacent node onto the stack.
			}
		}
	}
}

// 2. Implementing BFS using Queues:
void BFS(vector<vector<int>>& graph, int start){
// This function takes a reference to a graph represented as an adjacency list (vector<vector<int>>) 
// and an integer start indicating the starting node for BFS.
	vector<bool> visited(graph.size(), false);
	// A boolean vector visited to keep track of visited nodes. Initialized to false for all nodes.
	queue<int> q; // A queue q is used to manage the BFS process.
	visited[start] = true;
	q.push(start); // The starting node is marked as visited and pushed onto the queue.

	while(!q.empty()){
	// While the queue is not empty, the node at the front of the queue is processed.
		int node = q.front();
		cout << node << " ";
		q.pop(); // The node is printed and removed from the queue.

		// This loop iterates through the neighbors of the current node and processes them 
		// if they haven't been visited yet. 
		for(int i = 0; i < graph[node].size();i++){
			if(!visited[graph[node][i]]) {
				// graph[node][i] accesses the i-th neighbor of the current node.
				// visited[graph[node][i]] checks if this neighbor has been visited.
				// !visited[graph[node][i]] means the neighbor has not been visited.
				visited[graph[node][i]] = true; 
				// Mark this neighbor as visited to prevent it from being processed again.
				q.push(graph[node][i]);
				// Add this neighbor to the queue to ensure it will be processed in 
				// subsequent steps of the BFS algorithm.
			}
		}
	}
}


// main function: 
int main(){

	// 1. Implementing DFS using Stacks:
	vector<vector<int>> graph = { 
	// Initializes a 2D vector representing the adjacency list of the graph. Each index represents 
	// a node, and the inner vectors represent the nodes it is connected to.
		{1,2}, // Node 0 is connected to nodes 1 and 2.
		{0,3,4}, // Node 1 is connected to nodes 0, 3, and 4.
		{0,4}, // Node 2 is connected to nodes 0 and 4.
		{1,5}, // Node 3 is connected to nodes 1 and 5.
		{1,2,5}, // Node 4 is connected to nodes 1, 2, and 5.
		{3,4} // Node 5 is connected to nodes 3 and 4.
	};
	cout << "DFS starting from Node 0: "; // Prints a message indicating the start of DFS from node 0.
	DFS(graph, 0); // Calls the DFS function, passing the graph and the starting node (0) as argument.
	cout << endl; 

	// 2. Implementing BFS using Queues:
	cout << "BFS starting from Node 0: ";
	BFS(graph, 0);
	cout << endl;
	return 0;
}
