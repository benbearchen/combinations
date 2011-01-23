// Generating All Combinations and Permutations algorithms
// implementation
//
// Copyright (C) 2004 - 2007, 2011 BenBear
// benbearchen at gmail dot com
//
// Under boost license

#ifndef __BOOST_ALGORITHM_COMBINATIONS_DEF__
#define __BOOST_ALGORITHM_COMBINATIONS_DEF__

#include <algorithm>

#include "detail/combination.hpp"

namespace boost
{
  namespace combinations
  {
    /////////////////////////////////////////////////////////////////////
    // next_combination: get next combination.
    //
    // [first, middle): the elements selected in \\
    // [middle, last): the elements selected out
    /////////////////////////////////////////////////////////////////////
    template <typename BiIter>
    inline bool
    next_combination (BiIter first, BiIter middle, BiIter last)
    {
      return detail::do_next_combination (first, middle, middle, last);
    }
    
    /////////////////////////////////////////////////////////////////////
    // next_combination: get next combination.
    //
    // [first, middle): the elements selected in \\
    // [middle, last): the elements selected out
    /////////////////////////////////////////////////////////////////////
    template <typename BiIter, typename Comp>
    inline bool
    next_combination (BiIter first, BiIter middle, BiIter last, 
		      Comp comp)
    {
      return detail::do_next_combination (first, middle, 
					  middle, last, 
					  comp);
    }
    
    /////////////////////////////////////////////////////////////////////
    // prev_combination: get prev combination.
    //
    // [first, middle): the elements selected in \\
    // [middle, last): the elements selected out
    /////////////////////////////////////////////////////////////////////
    template <typename BiIter>
    inline bool 
    prev_combination (BiIter first, BiIter middle, BiIter last)
    {
      return detail::do_next_combination (middle, last, first, middle);
    }
    
    /////////////////////////////////////////////////////////////////////
    // prev_combination: get prev combination.
    //
    // [first, middle): the elements selected in \\
    // [middle, last): the elements selected out
    /////////////////////////////////////////////////////////////////////
    template <typename BiIter, typename Comp>
    inline bool 
    prev_combination (BiIter first, BiIter middle, BiIter last, 
		      Comp comp)
    {
      return detail::do_next_combination (middle, last, 
					  first, middle, 
					  comp);
    }

    template <typename BiIter>
    inline bool
    next_permutation (BiIter first, BiIter middle, BiIter last)
    {
      if (first == middle)
	return false;
      
      std::reverse (middle, last);
      return std::next_permutation (first, last);
    }

    template <typename BiIter, typename Comp>
    inline bool
    next_permutation (BiIter first, BiIter middle, BiIter last, 
		      Comp comp)
    {
      if (first == middle)
	return false;

      std::reverse (middle, last);
      return std::next_permutation (first, last, comp);
    }

    template <typename BiIter>
    inline bool 
    prev_permutation (BiIter first, BiIter middle, BiIter last)
    {
      if (first == middle)
	return false;

      bool ret = std::prev_permutation (first, last);
      std::reverse (middle, last);

      return ret;
    }

    template <typename BiIter, typename Comp>
    inline bool 
    prev_permutation (BiIter first, BiIter middle, BiIter last, 
		      Comp comp)
    {
      if (first == middle)
	return false;

      bool ret = std::prev_permutation (first, last, comp);
      std::reverse (middle, last);

      return ret;
    }
  }
}

#endif
