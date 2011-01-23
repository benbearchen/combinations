#include "../combinations.hpp"
#include "../combinations_init.hpp"
using namespace boost::combinations;

#include <iostream>
using namespace std;

template <typename Iter>
void out (Iter first, Iter middle, Iter last)
{
  for (Iter i = first; i != middle; ++i)
    cout << "\t" << *i;

  if (middle != last)
    {
      cout << "\t[" << *middle;
      for (Iter i = middle; ++i != last;)
	cout << "\t" << *i;
      cout << "]";
    }
  cout << endl;
}

int main ()
{
  char a[5] = {'a', 'b', 'c', 'd', 'e'};

  int k = 2, n = 5;

  cout << endl
       << "select 2 from 5, generate from smallest to largest:"
       << endl;
  // initialize to the smallest sequence
  init_permutation (a, a + k, a + n, true);
  
  // generate loop
  do
    {
      // use the sequence
      out (a, a + k, a + n);
    }
  // generate to next sequence, and judge whether end
  while (next_permutation (a, a + k, a + n));
  cout << "press ENTER to continue...";
  cin.get();


  k = n - k;

  cout << endl
       << "select 3 from 5, generate from largest to smallest:"
       << endl;
  // previous loop
  init_permutation (a, a + k, a + n, false);
  do
    {
      out (a, a + k, a + n);
    }
  while (prev_permutation (a, a + k, a + n));
  cout << "press ENTER to finish...";
  cin.get();
}
