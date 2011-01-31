#include "../combinations.hpp"
#include "../combinations_init.hpp"

#include <string>
#include <stdio.h>

int main()
{
  int a[9] = {1, 2, 3, 4, 5, 6, 7, 8, 9};
  int n = 9;
  int m = 4;
  boost::combinations::init_permutation(a, a + m, a + n, false);
  // or
  boost::combinations::init_permutation(a, a + m, a + n, true);
  boost::combinations::prev_permutation(a, a + m, a + n);
  // or
  std::sort(a, a + n);
  std::reverse(a, a + n);
  std::reverse(a + m, a + n);
  
  do
    {
      if (a[0] + a[1] * a[2] - a[3] == 32)
	printf("%d+%d*%d-%d=32\n", a[0], a[1], a[2], a[3]);
    }
  while (boost::combinations::prev_permutation(a, a + m, a + n));
}
