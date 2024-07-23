#include <iostream>
#include <stack>
using namespace std;

#define N 5

bool isSafe(int maze[N][N], int x, int y) {
	return (x >= 0 && x < N && y >= 0 && y < N && maze[x][y] == 1);
}

bool solveMazeDFS(int maze[N][N]) {
	stack<pair<int, int>> s;
	s.push({0,0});

	int sol[N][N] = {0};

	while (!s.empty()){
		pair<int, int> pos = s.top();
		s.pop();

		int x = pos.first;
		int y = pos.second;

		if(x == N - 1 && y == N - 1) {
			sol[x][y] = 1;
			for(int i = 0; i < N; i++) {
				for(int j = 0; j < N; j++){
					cout << sol[i][j] << " ";
				}
				cout << endl;
			}
			return true;
		}

		if (isSafe(maze, x, y)){
				sol[x][y] = 1;

				if(isSafe(maze, x+1, y)) s.push({x+1, y});
				if(isSafe(maze, x, y+1)) s.push({x, y+1});
				if(isSafe(maze, x-1, y)) s.push({x-1, y});
				if(isSafe(maze, x, y-1)) s.push({x, y-1});
			}
	}
	cout << "No solution found" << endl;
	return false;
}

int main(){
	int maze[N][N] = {
		{1,0,0,0,0},
		{1,1,0,1,1},
		{0,1,1,0,0},
		{1,0,1,1,1},
		{1,1,1,0,1}
		};
	solveMazeDFS(maze);
	return 0;


}
