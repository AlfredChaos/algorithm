#include <iostream>
#include <string>
using namespace std;

class Soluction
{
public:
    string generateTheSTring(int n)
    {
        if (n % 2 == 0)
        {
            return string(n - 1, 'a') + 'b';
        }
        else
        {
            return string(n, 'a');
        }
    }
};

int main()
{
    Soluction so;
    cout << so.generateTheSTring(3) << endl;
}