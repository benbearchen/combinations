#include "../combinations.hpp"
#include "../combinations_init.hpp"
using namespace boost::combinations;

#include "compare.hpp"

#include <iostream>
#include <cstring>
#include <cstdio>
using namespace std;

void
my_throw (const char* msg, int line)
{
  size_t len = strlen (msg);
  char* str = new char[len + 20];
  strncpy (str, msg, len);
  sprintf (str + len, " at line: %d", line);
  throw (const char*)str;
}

template <typename Iter>
void
print_seq (Iter first, Iter last)
{
  if (first != last)
    {
      cout << *first;
      while (++first != last)
	cout << "," << *first;
    }
  cout << endl;
}

template <typename Iter>
int
get_number (Iter first, Iter middle, Iter last)
{
  int cn = 0;
  init_permutation (first, middle, last, true);

  do
    {
      if (!test_sequence_grow (middle, last))
	my_throw ("test order failed", __LINE__);

      print_seq (first, middle);
      ++cn;
    }
  while (next_permutation (first, middle, last));
  cout << "next count: " << cn << endl;

  int cp = 0;
  init_permutation (first, middle, last, false);
  
  do
    {
      if (!test_sequence_grow (middle, last))
	my_throw ("test order failed", __LINE__);

      print_seq (first, middle);
      ++cp;
    }
  while (prev_permutation (first, middle, last));
  cout << "prev cout: " << cp << endl;
  
  if (cn != cp)
    my_throw ("next_permutation != prev_permutation", __LINE__);

  return cn;
}

int main ()
{
  try
    {
      int a[5] = {1, 2, 3, 4, 5};
      int g = get_number (a, a + 2, a + 5);
      cout << "test is ok!" << endl;
    }
  catch (const char* except)
    {
      cout << "Exception:"
	   << except
	   << endl;
    }
}
