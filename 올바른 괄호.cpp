// link: https://school.programmers.co.kr/learn/courses/30/lessons/12909?language=cpp

#include <string>
#include <iostream>
#include <stack>

using namespace std;

bool solution(string s)
{
    bool answer = true;
    int count=0;
    stack<char> stack;
    
// #1 using stack
    for (int i=0; i<s.length(); i++){
    	if (s[i]=='(') stack.push(s[i]);
        else {
            if (!stack.empty())
                stack.pop();
            else return false;
        }
    }
    answer = stack.empty();
    
// #2 using for loop
    
//     for (int i=0; i<s.length(); i++){
//       (s[i]=='(') ? count++ : count--;
//       if (count < 0 && s[i]==')') return false;
//     }
//     if (count != 0) answer = false;
    
    return answer;
}
