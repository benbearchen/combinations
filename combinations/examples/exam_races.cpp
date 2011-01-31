#include "../combinations.hpp"
#include "../combinations_init.hpp"

#include <string>
#include <stdio.h>

int main()
{
  std::string races[] = {"Ancients", "Asgard", "Earth", "Furling", "Goa'uld",
			 "Nox", "Ori", "Tok'ra", "Tollan", "Unas"};
  int n = 10;
  int m = 3;

  boost::combinations::init_combination(races, races + m, races + n, true);
  // or
  std::sort(races, races + n);
  
  do
    {
      // visit [races, races + m), such as:
      for (int i = 0; i < m; ++i)
	printf("%15s,", races[i].c_str());
      printf("\n");
    }
  while (boost::combinations::next_combination(races, races + m, races + n));
}
