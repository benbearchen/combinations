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
  init_combination (first, middle, last, true);

  test_sequence_serial<Iter> nfm(true);
  do
    {
      if (!test_sequence_grow (first, middle))
	my_throw ("test order failed", __LINE__);
      nfm.push_sequence (first, middle);
      if (!test_sequence_grow (middle, last))
	my_throw ("test order failed", __LINE__);

      print_seq (first, middle);
      ++cn;
    }
  while (next_combination (first, middle, last));
  cout << "next cout: " << cn << endl;

  int cp = 0;
  init_combination (first, middle, last, false);
  
  test_sequence_serial<Iter> pfm(false);
  do
    {
      if (!test_sequence_grow (first, middle))
	my_throw ("test order failed", __LINE__);
      pfm.push_sequence (first, middle);
      if (!test_sequence_grow (middle, last))
	my_throw ("test order failed", __LINE__);

      print_seq (first, middle);
      ++cp;
    }
  while (prev_combination (first, middle, last));
  cout << "prev cout: " << cp << endl;
  
  if (cn != cp)
    my_throw ("next_combination != prev_combination", __LINE__);

  return cn;
}

int main ()
{
  try
    {
      int a[5] = {1, 2, 3, 4, 5};
      int g = get_number (a, a + 2, a + 5);
      int t = test_get_combi_count (a, a + 2, a + 5);

      if (g != t)
	{
	  cout << g << " != " << t << endl;
	  my_throw ("get number != test get number", __LINE__);
	}
      cout << "test is ok!" << endl;
    }
  catch (const char* except)
    {
      cout << "Exception:"
	   << except
	   << endl;
    }
}
