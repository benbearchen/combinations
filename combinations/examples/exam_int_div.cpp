#include "../combinations.hpp"
#include <vector>
#include <stdio.h>

int main()
{
  int n = 0;
  printf("Please input a integer(>0): ");
  if ((scanf("%d", &n) != 1) || (n < 1))
    {
      printf("ERROR: integer should >0! \n");
      return -1;
    }

  std::vector<int> a;
  for (int i = 2; n != 1; ++i)
    {
      while (n % i == 0)
	{
	  a.push_back(i);
	  n /= i;
	}
    }

  for (size_t i = 0; i <= a.size(); ++i)
    {
      do
	{
	  int m = 1;
	  for (size_t j = 0; j < i; ++j)
	    m *= a[j];
	  printf("%d\n", m);
	}
      while (boost::combinations::next_combination(a.begin(), a.begin() + i, a.end()));
    }
}
