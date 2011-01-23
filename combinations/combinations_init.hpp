// Generating All Combinations and Permutations algorithms
// implementation
//
// Copyright (C) 2004 - 2007, 2011 BenBear
// benbearchen at gmail dot com
//
// Under boost license

#ifndef __BOOST_ALGORITHM_COMBINATIONS_INIT_DEF__
#define __BOOST_ALGORITHM_COMBINATIONS_INIT_DEF__

#include <algorithm>

#include "detail/combi_utility.hpp"

namespace boost
{
  namespace combinations
  {
    //////////////////////////////////////////////////////////////////////
    // init_combination: init the (first, midle, last) to the min or the
    //                   max combination
    //////////////////////////////////////////////////////////////////////
    template <typename BiIter>
    void
    init_combination (BiIter first, BiIter middle, BiIter last, 
		      bool min = true)
    {
      detail::sort_combination (first, middle);
      detail::sort_combination (middle, last);
      if (min)
	detail::twice_merge (first, middle, middle, last);
      else
	detail::twice_merge (middle, last, first, middle);
    }

    //////////////////////////////////////////////////////////////////////
    // init_combination: init the (first, midle, last) to the min or the
    //                   max combination
    //////////////////////////////////////////////////////////////////////
    template <typename BiIter, typename Comp>
    void
    init_combination (BiIter first, BiIter middle, BiIter last, 
		      bool min, Comp comp)
    {
      detail::sort_combination (first, middle, comp);
      detail::sort_combination (middle, last, comp);
      if (min)
	detail::twice_merge (first, middle, middle, last, comp);
      else
	detail::twice_merge (middle, last, first, middle, comp);
    }

    //////////////////////////////////////////////////////////////////////
    // adjust_combination: make the (first, middle, last) to a right
    //                     combination. [first, middle) are the elements
    //                     selected in, [middle, last) are selected out
    //////////////////////////////////////////////////////////////////////
    template <typename BiIter>
    void
    adjust_combination (BiIter first, BiIter middle, BiIter last)
    {
      detail::sort_combination (first, middle);
      detail::sort_combination (middle, last);
    }

    //////////////////////////////////////////////////////////////////////
    // adjust_combination: make the (first, middle, last) to a right
    //                     combination. [first, middle) are the elements
    //                     selected in, [middle, last) are selected out
    //////////////////////////////////////////////////////////////////////
    template <typename BiIter, typename Comp>
    void
    adjust_combination (BiIter first, BiIter middle, BiIter last,
			Comp comp)
    {
      detail::sort_combination (first, middle, comp);
      detail::sort_combination (middle, last, comp);
    }

    template <typename BiIter>
    void
    init_permutation (BiIter first, BiIter middle, BiIter last, 
		      bool min = true)
    {
      detail::sort_combination (first, middle);
      detail::sort_combination (middle, last);
      if (min)
	detail::twice_merge (first, middle, middle, last);
      else
	{
	  detail::twice_merge (middle, last, first, middle);
	  std::reverse (first, middle);
	}
    }

    template <typename BiIter, typename Comp>
    void
    init_permutation (BiIter first, BiIter middle, BiIter last, 
		      bool min, Comp comp)
    {
      detail::sort_combination (first, middle, comp);
      detail::sort_combination (middle, last, comp);
      if (min)
	detail::twice_merge (first, middle, middle, last, comp);
      else
	{
	  detail::twice_merge (middle, last, first, middle, comp);
	  std::reverse (first, middle);
	}
    }

    template <typename BiIter>
    void
    adjust_permutation (BiIter first, BiIter middle, BiIter last)
    {
      detail::sort_combination (middle, last);
    }

    template <typename BiIter, typename Comp>
    void
    adjust_permutation (BiIter first, BiIter middle, BiIter last,
			Comp comp)
    {
      detail::sort_combination (middle, last, comp);
    }
  }
}
#endif
