// link: https://school.programmers.co.kr/learn/courses/30/lessons/120924?language=cpp

#include <string>
#include <vector>

using namespace std;

int solution(vector<int> common) {
    int answer = 0;
    bool isArithmetic = true;
    int difference = common[1] - common[0];
    
    for (int i = 1; i < common.size(); i++){
        if (common[i] - common[i-1] != difference) {
            isArithmetic = false;
            break;
        }
    }
    // Arithmetic sequence
    if (isArithmetic) {
        answer = common[common.size()-1] + difference;
    }
    // geometric sequence
    else {
        difference = common[1] / common[0];
        answer = common[common.size()-1] * difference;
    }
    
    return answer;
}
