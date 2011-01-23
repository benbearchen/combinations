#ifndef __combi_test_compare_hpp
#define __combi_test_compare_hpp

#include <algorithm>
#include <iostream>
#include <vector>
using namespace std;

template <typename Iter>
bool
test_sequence_grow (Iter first, Iter last)
{
  if (first == last)
    return true;

  for (Iter a = first, b = first; ++b != last; ++a)
    {
      if (*b < *a)
	return false;
    }
  return true;
}

template <typename Iter, typename Comp>
bool
test_sequence_grow (Iter first, Iter last, Comp comp)
{
  if (first == last)
    return true;

  for (Iter a = first, b = first; ++b != last; ++a)
    {
      if (comp (*b, *a))
	return false;
    }
  return true;
}

template <typename Iter1, typename Iter2>
bool
test_two_sequence_less (Iter1 first1, Iter1 last1,
			Iter2 first2, Iter2 last2)
{
  return lexicographical_compare (first1, last1, 
				  first2, last2);
}

template <typename Iter1, typename Iter2, typename Comp>
bool
test_two_sequence_less (Iter1 first1, Iter1 last1,
			Iter2 first2, Iter2 last2,
			Comp comp)
{
  return lexicographical_compare (first1, last1, 
				  first2, last2,
				  comp);
}

template <typename RandomIter>
int
test_get_combi_count (RandomIter first, RandomIter middle, RandomIter last)
{
  if (first == middle)
    return 0;

  sort (first, last);
  
  int count = 0;
  do
    {
      sort (middle, last);
      reverse (middle, last);
      if (test_sequence_grow (first, middle))
	++count;
    }
  while (next_permutation (first, last));

  return count;
}

// warning: sequence1 < sequence2 may be differenct with sequence2 > sequence1
//          because sequence1's comp object may be differenct with sequence2's
template <typename Iter, 
	  typename Comp = std::less<typename std::iterator_traits<Iter>::value_type> >
class sequence
{
  typedef typename iterator_traits<Iter>::value_type value_type;

  vector<value_type> values;

  Comp compare;
public:
  typedef typename vector<value_type>::const_iterator iterator;
  
  explicit
  sequence (Comp comp = Comp())
    : compare(comp)
  {
  }

  template <typename Iter2>
  sequence (Iter2 begin, Iter2 end, Comp comp = Comp())
    : values(begin, end), compare(comp)
  {
  }

  template <typename Iter2, typename Comp2>
  sequence&
  operator = (const sequence<Iter2, Comp2>& seq)
  {
    values.assign (seq.begin(), seq.end());
    return *this;
  }

  template <typename Iter2>
  void
  assign (Iter2 begin, Iter2 end)
  {
    values.assign (begin, end);
  }
  
  template <typename Iter2, typename Comp2>
  bool
  operator == (const sequence<Iter2, Comp2>& seq) const
  {
    return (!(*this < seq) && !(seq < *this));
  }
  
  template <typename Iter2, typename Comp2>
  bool
  operator != (const sequence<Iter2, Comp2>& seq) const
  {
    return !(*this == seq);
  }
  
  template <typename Iter2, typename Comp2>
  bool
  operator < (const sequence<Iter2, Comp2>& seq) const
  {
    return test_two_sequence_less (begin(), end(),
				   seq.begin(), seq.end(), 
				   compare);
  }
  template <typename Iter2, typename Comp2>
  bool
  operator > (const sequence<Iter2, Comp2>& seq) const
  {
    return test_two_sequence_less (seq.begin(), seq.end(),
				   begin(), end(),
				   compare);
  }
  template <typename Iter2, typename Comp2>
  bool
  operator <= (const sequence<Iter2, Comp2>& seq) const
  {
    return !(*this > seq);
  }

  template <typename Iter2, typename Comp2>
  bool
  operator >= (const sequence<Iter2, Comp2>& seq) const
  {
    return !(*this < seq);
  }

  iterator
  begin() const
  {
    return values.begin();
  }

  iterator
  end() const
  {
    return values.end();
  }

  void
  clear ()
  {
    values.clear ();
  }

  void
  out () const
  {
    for (iterator i = begin(); i != end(); ++i)
      cout << *i << " ";
    cout << endl;
  }
};

template <typename Iter, 
	  typename Comp = std::less<typename std::iterator_traits<Iter>::value_type> >
class test_sequence_serial
{
  typedef typename iterator_traits<Iter>::value_type value_type;

  sequence<Iter, Comp> values;
  bool direction;
  int count;
public:
  explicit
  test_sequence_serial (bool dir = true, Comp comp = Comp())
    : direction(dir), values(comp), count(0)
  {
  }

  template <typename Iter2>
  void push_sequence (Iter2 first, Iter2 last)
  {
    sequence<Iter> next(first, last);

    if (values.begin() == values.end())
      {
	values = next;
	++count;
	return;
      }


    if ((direction && !(values < next))
	|| (!direction && !(values > next)))
      {
	values.out ();
	next.out ();
	cout << "direction:" << direction << endl;
	throw "last sequence is not small than new sequence";
      }
    values = next;
    ++count;
  }

  void clear (bool dir)
  {
    values.clear ();
    direction = dir;
    count = 0;
  }

  int get_count () const
  {
    return count;
  }
};

template <class Picker>
int
test_picker_next_count (Picker& picker)
{
  picker.init (true);

  test_sequence_serial<typename Picker::iterator> serial(true);

  int count = 0;
  do
    {
      serial.push_sequence (picker.begin(), picker.end());
      ++count;
    }
  while (picker.next ());

  return count;
}

template <class Picker>
int
test_picker_prev_count (Picker& picker)
{
  picker.init (false);

  test_sequence_serial<typename Picker::iterator> serial(false);

  int count = 0;
  do
    {
      serial.push_sequence (picker.begin(), picker.end());
      ++count;
    }
  while (picker.prev ());

  return count;
}

template <class Picker1, class Picker2>
bool
test_two_picker_equal (Picker1& picker1, Picker2& picker2)
{
  vector<sequence<typename Picker1::iterator> > stack;

  picker1.init (true);
  picker2.init (true);

  bool finish1, finish2;
  do
    {
      sequence<typename Picker1::iterator> seq1(picker1.begin(),
						picker1.end());
      sequence<typename Picker2::iterator> seq2(picker2.begin(),
						picker2.end());
      if (seq1 != seq2)
	{
	  seq1.out ();
	  seq2.out ();
	  throw "don't equal while next of picker1 and picker2";
	}

      stack.push_back (seq1);

      finish1 = picker1.next ();
      finish2 = picker2.next ();
    }
  while (finish1 && finish2);

  if (finish1 != finish2)
    {
      throw "the times of picker1 and picker2'next are not same";
    }

  picker1.init (false);
  picker2.init (false);

  do
    {
      sequence<typename Picker1::iterator> seq1(picker1.begin(),
						picker1.end());
      sequence<typename Picker2::iterator> seq2(picker2.begin(),
						picker2.end());
      if (seq1 != seq2)
	{
	  seq1.out ();
	  seq2.out ();
	  throw "don't equal while prev of picker1 and picker2";
	}

      if (stack.empty())
	{
	  throw "prev is longer than next";
	}
      if (seq1 != stack.back())
	{
	  seq1.out ();
	  stack.back().out ();
	  throw "prev is not equal next";
	}

      stack.pop_back ();

      finish1 = picker1.prev ();
      finish2 = picker2.prev ();
    }
  while (finish1 && finish2);

  if (finish1 != finish2)
    {
      throw "the times of picker1 and picker2'prev are not same";
    }

  if (!stack.empty())
    {
      throw "next is longer than prev";
    }

  return test_picker_next_prev (picker1) && test_picker_next_prev (picker2);
}

template <class Picker>
bool
test_picker_next_prev (Picker& picker)
{
  typedef sequence<typename Picker::iterator> seq_t;

  picker.init (true);
  
  seq_t seq_begin(picker.begin(), picker.end());
  seq_t seq_last = seq_begin;

  do
    {
      picker.next ();
      seq_t seq_next(picker.begin(), picker.end());

      picker.prev ();
      seq_t seq_prev(picker.begin(), picker.end());

      if (seq_last != seq_prev)
	{
	  cout << "orig:"; seq_last.out ();
	  cout << "next:"; seq_next.out ();
	  cout << "prev:"; seq_prev.out ();
	  throw "seq_last != (picker.next().prev())";
	}

      picker.next ();
      seq_last = seq_t(picker.begin(), picker.end());
      
      if (seq_last != seq_next)
	{
	  cout << "next:"; seq_next.out ();
	  cout << "prev:"; seq_prev.out ();
	  cout << "next:"; seq_last.out ();
	  throw "seq_next != (picker.next().prev().next())";
	}
    }
  while (seq_last != seq_begin);

  return true;
}

#endif
