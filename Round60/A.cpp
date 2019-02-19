#include <iostream>
#include <string>
#include <cstdlib>

using namespace std;

typedef struct Group {
    int start;
    int end;
    float mean;
}group;

int main() {

    string input;
    getline (cin, input);
    int num = atol(input.c_str());
    int array[num];

    string inputs;
    getline (cin,inputs);

    int j = 0;
    for (int i = 0 ; i < inputs.length() ; i++){

        if (inputs[i] == ' ') {}
        else {
            array[j] = inputs[i]-48;
            j++;
        }
    }

    float best = 0;
    group p;
    for (int i = 0 ; i < (sizeof(array)/sizeof(int))-1 ; i++) {
        for (int j = i+1 ; j < sizeof(array)/sizeof(int) ; j++){
            float sum = 0;
            for (int k = i ; k <= j ; k++){
                sum += array[k];
            }
            sum = sum/(j-i+1);
            if (best == sum) {
                if (j-i > p.end-p.start){
                    p.start = i;
                    p.end = j;
                    p.mean = sum;
                    best = sum;
                }
            }
            else if (best < sum){
                p.start = i;
                p.end = j;
                p.mean = sum;
                best = sum;
            }
        }
    }
    cout<<p.end<<p.start<<endl;
    cout<<p.end-p.start+1<<endl;

    return 0;
}
