// Generating All Combinations and Permutations algorithms
// implementation
//
// Copyright (C) 2004 - 2007, 2011 BenBear
// benbearchen at gmail dot com
//
// Under boost license

#ifndef __BOOST_ALGORITHM_COMBINATIONS_DETAIL_COMBINATION_HPP__
#define __BOOST_ALGORITHM_COMBINATIONS_DETAIL_COMBINATION_HPP__

#include <algorithm>

#include "combi_utility.hpp"

namespace boost
{
  namespace combinations
  {
    namespace detail
    {
      /////////////////////////////////////////////////////////////////////
      // do_next_combination: get next combination.
      //
      // [first1, last1): the elements selected in \\
      // [first2, last2): the elements selected out
      /////////////////////////////////////////////////////////////////////
      template <typename BiIter>
      bool
      do_next_combination (BiIter first1, BiIter last1, 
			   BiIter first2, BiIter last2)
      {
	if ((first1 == last1) || (first2 == last2))
	  return false;
      
	BiIter qmax = last2;
	--qmax;
	BiIter pout1 = std::lower_bound (first1, last1, *qmax);
	bool fin = (pout1 == first1);
	BiIter left1, left2;
	if (!fin)
	  {
	    BiIter pout = pout1;
	    --pout;
	    BiIter qin = std::upper_bound (first2, last2, *pout);
	    std::iter_swap (pout, qin);
	    left1 = pout;
	    ++left1;
	    left2 = qin;
	    ++left2;
	  }
	else
	  {
	    left1 = first1;
	    left2 = first2;
	  }
	combination_merge_right (left1, last1, left2, last2);
	return !fin;
      }
    
      /////////////////////////////////////////////////////////////////////
      // do_next_combination: get next combination.
      //
      // [first1, last1): the elements selected in \\
      // [first2, last2): the elements selected out
      /////////////////////////////////////////////////////////////////////
      template <typename BiIter, typename Comp>
      bool
      do_next_combination (BiIter first1, BiIter last1, 
			   BiIter first2, BiIter last2,
			   Comp comp)
      {
	if ((first1 == last1) || (first2 == last2))
	  return false;
      
	BiIter qmax = last2;
	--qmax;
	BiIter pout1 = std::lower_bound (first1, last1, *qmax, comp);
	bool fin = (pout1 == first1);
	BiIter left1, left2;
	if (!fin)
	  {
	    BiIter pout = pout1;
	    --pout;
	    BiIter qin = std::upper_bound (first2, last2, *pout, comp);
	    std::iter_swap (pout, qin);
	    left1 = pout;
	    ++left1;
	    left2 = qin;
	    ++left2;
	  }
	else
	  {
	    left1 = first1;
	    left2 = first2;
	  }
	combination_merge_right (left1, last1, left2, last2);
	return !fin;
      }
    }
  }
}

#endif
