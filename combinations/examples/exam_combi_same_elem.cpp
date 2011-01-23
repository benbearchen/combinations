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
       << "<2 of 5>, <smallest to largest>:"
       << endl;
  // initialize
  init_combination (a, a + k, a + n, true);
  
  // generate loop
  do
    {
      // use the sequence
      out (a, a + k, a + n);
    }
  // generate to next sequence, and judge whether end
  while (next_combination (a, a + k, a + n));
  cout << "press ENTER to continue...";
  cin.get();

  ////////////////////////////////////////////////////

  a[1] = a[0];
  
  cout << endl
       << "now make it has same elements:"
       << endl;
  out (a, a + n, a + n);

  cout << "then do generating again:"
       << endl;

  init_combination (a, a + k, a + n, true);
  do
    {
      out (a, a + k, a + n);
    }
  while (next_combination (a, a + k, a + n));
  cout << "press ENTER to continue...";
  cin.get();

  ////////////////////////////////////////////////////

  a[3] = a[2];

  cout << endl
       << "now make more elements have the same value:"
       << endl;
  out (a, a + n, a + n);

  cout << "then do generating again"
       << endl;
  init_combination (a, a + k, a + n, true);
  do
    {
      out (a, a + k, a + n);
    }
  while (next_combination (a, a + k, a + n));
  cout << "press ENTER to Finish...";
  cin.get();
}
