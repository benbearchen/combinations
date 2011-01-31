// Generating All Combinations and Permutations algorithms
// implementation
//
// Copyright (C) 2007, BenBear
// benbearchen at gmail dot com
// 
// Under boost license

#ifndef __BOOST_ALGORITHM_COMBINATIONS_DETAIL_COMBI_UTILITY_DEF__
#define __BOOST_ALGORITHM_COMBINATIONS_DETAIL_COMBI_UTILITY_DEF__

namespace boost
{
  namespace combinations
  {
    namespace detail
    {
      ////////////////////////////////////////////////////////////////////
      // twice_merge: merge [first1, last1) and [first2, last2), then
      //                fill the min element to [first1, last1), left to
      //                [first2, last2)
      ////////////////////////////////////////////////////////////////////
      template <typename Iter1, typename Iter2>
      void
      twice_merge (Iter1 first1, Iter1 last1, 
		   Iter2 first2, Iter2 last2)
      {
	typedef typename std::iterator_traits<Iter1>::value_type value_type;
	typedef typename std::iterator_traits<Iter1>::difference_type diff_t;
      
	if ((first1 == last1) || (first2 == last2))
	  return;
      
	diff_t len1 = std::distance(first1, last1);
	diff_t len2 = std::distance(first2, last2);
	diff_t lent = len1 + len2;
	value_type *tmp = new value_type[lent];
      
	std::merge (first1, last1, first2, last2, tmp);
	std::copy (tmp, tmp + len1, first1);
	std::copy (tmp + len1, tmp + lent, first2);
      
	delete[] tmp;
      }
    
      ////////////////////////////////////////////////////////////////////
      // twice_merge: merge [first1, last1) and [first2, last2), then
      //                fill the min element to [first1, last1), left to
      //                [first2, last2)
      ////////////////////////////////////////////////////////////////////
      template <typename Iter1, typename Iter2, typename Comp>
      void
      twice_merge (Iter1 first1, Iter1 last1, 
		   Iter2 first2, Iter2 last2, 
		   Comp comp)
      {
	typedef typename std::iterator_traits<Iter1>::value_type value_type;
	typedef typename std::iterator_traits<Iter1>::difference_type diff_t;
      
	if ((first1 == last1) || (first2 == last2))
	  return;
      
	diff_t len1 = std::distance(first1, last1);
	diff_t len2 = std::distance(first2, last2);
	diff_t lent = len1 + len2;
	value_type *tmp = new value_type[lent];
      
	std::merge (first1, last1, first2, last2, tmp, comp);
	std::copy (tmp, tmp + len1, first1);
	std::copy (tmp + len1, tmp + lent, first2);
      
	delete[] tmp;
      }
    
      ////////////////////////////////////////////////////////////////////
      // twice_merge: merge [first1, last1) and [first2, last2), then
      //                fill the min element to [first1, last1), left to
      //                [first2, last2), using buffer
      ////////////////////////////////////////////////////////////////////
      template <typename Iter1, typename Iter2, 
		typename Iter3, typename Comp>
      void
      twice_merge (Iter1 first1, Iter1 last1, 
		   Iter2 first2, Iter2 last2,
		   Iter3 buffer, Comp comp)
      {
	typedef typename std::iterator_traits<Iter1>::value_type value_type;
	typedef typename std::iterator_traits<Iter1>::difference_type diff_t;
      
	if ((first1 == last1) || (first2 == last2))
	  return;
      
	diff_t len1 = std::distance(first1, last1);
	diff_t len2 = std::distance(first2, last2);
	diff_t lent = len1 + len2;
      
	std::merge (first1, last1, first2, last2, buffer, comp);
	std::copy (buffer, buffer + len1, first1);
	std::copy (buffer + len1, buffer + lent, first2);
      }
    
      ///////////////////////////////////////////////////////////////////
      // sort_combination: merge sort the [first, last)
      ///////////////////////////////////////////////////////////////////
      template <typename BiIter, typename Comp>
      void
      sort_combination (BiIter first, BiIter last, Comp comp)
      {
	typedef typename std::iterator_traits<BiIter>::difference_type diff_t;
	diff_t len = std::distance(first, last);
	if (len <= 1)
	  return;
      
	if (len == 2)
	  {
	    if (comp(*--last, *first))
	      std::iter_swap (first, last);
	  }
	else
	  {
	    BiIter middle = first;
	    std::advance(middle, len / 2);
	    sort_combination (first, middle, comp);
	    sort_combination (middle, last, comp);
	    twice_merge (first, middle, middle, last, comp);
	  }
      }
    
      ///////////////////////////////////////////////////////////////////
      // sort_combination: merge sort the [first, last)
      ///////////////////////////////////////////////////////////////////
      template <typename BiIter, typename BufferIter, typename Comp>
      void
      sort_combination (BiIter first, BiIter last, BufferIter buffer, Comp comp)
      {
	typedef typename std::iterator_traits<BiIter>::difference_type diff_t;
	diff_t len = std::distance(first, last);
	if (len <= 1)
	  return;
      
	if (len == 2)
	  {
	    if (comp(*--last, *first))
	      std::iter_swap (first, last);
	  }
	else
	  {
	    BiIter middle = first;
	    std::advance(middle, len / 2);
	    sort_combination (first, middle, buffer, comp);
	    sort_combination (middle, last, buffer, comp);
	    twice_merge (first, middle, middle, last, 
			 buffer, comp);
	  }
      }
    
      ///////////////////////////////////////////////////////////////////
      // sort_combination: merge sort the [first, last)
      ///////////////////////////////////////////////////////////////////
      template <typename BiIter>
      void
      sort_combination (BiIter first, BiIter last)
      {
	typedef typename std::iterator_traits<BiIter>::difference_type diff_t;
	diff_t len = std::distance(first, last);
	if (len <= 1)
	  return;
      
	if (len == 2)
	  {
	    if (*--last < *first)
	      std::iter_swap (first, last);
	  }
	else
	  {
	    BiIter middle = first;
	    std::advance(middle, len / 2);
	    sort_combination (first, middle);
	    sort_combination (middle, last);
	    twice_merge (first, middle, middle, last);
	  }
      }
    
      ///////////////////////////////////////////////////////////////////
      // combination_merge_right: merge the two right parts in
      //            combination.
      //
      //  REQUIRE: [first1, last1) and [first2, last2) should be
      //            ordered; any first1 should be not less than
      //            first2.
      //
      //  RESULT: smallest elements in [first1, last1), while left in
      //            [first2, last2).
      ///////////////////////////////////////////////////////////////////
      template <typename BiIter1, typename BiIter2>
      void
      combination_merge_right (BiIter1 first1, BiIter1 last1,
			       BiIter2 first2, BiIter2 last2)
      {
	if ((first1 == last1) || (first2 == last2))
	  return;
      
	BiIter1 m1 = last1;
	BiIter2 m2 = first2;
	while ((m1 != first1) && (m2 != last2))
	  std::iter_swap (--m1, m2++);
      
	std::reverse (first1, m1);
	std::reverse (first1, last1);
      
	std::reverse (m2, last2);
	std::reverse (first2, last2);
      }
    }
  }
}

#endif

