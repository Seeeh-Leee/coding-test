//link : https://school.programmers.co.kr/learn/courses/30/lessons/120808 

#include <string>
#include <vector>

using namespace std;
int gcdf(int big, int small);

vector<int> solution(int numer1, int denom1, int numer2, int denom2) {
    vector<int> answer;
    int n0, d0, gcd =0;

    n0 = numer1 * denom2 + numer2 * denom1;
    d0 = denom1 * denom2;
    
    if (d0 > n0) gcd = gcdf(d0,n0);
    else gcd = gcdf(n0,d0);    
    
    
    if (gcd != 0){
        n0 = n0 / gcd;
        d0 = d0 / gcd;
    }
    
    answer.push_back(n0); 
    answer.push_back(d0); 
        
    return answer;
}

int gcdf(int big, int small)
{
    int remainder;
    
    while(small)
    {
        remainder = big % small;
        big = small;
        small = remainder;
    }
    return big;
}
