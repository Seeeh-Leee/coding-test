// link : https://school.programmers.co.kr/learn/courses/30/lessons/12939

#include <string>
#include <vector>

using namespace std;

string solution(string s) {
    string answer = "";
    int start = 0, end = 0, toInt = 0;
    vector<int> v;
    string tmp;
    
    for(int i=0;i<s.length();i++){
        if(s[i]!= ' '){
            tmp += s[i];
        }else{
            v.push_back(stoi(tmp));
            tmp.clear();
        }
    }
    v.push_back(stoi(tmp));
    
    start = v[0]; //min
    end = v[0]; //max
    
    for (int i=1; i< v.size(); i++) {
        if(v[i] < start) start = v[i];
        if(v[i] > end) end = v[i];
    }
    
    answer = to_string(start) + " " + to_string(end); 
    cout << answer;
    
    return answer;
}
