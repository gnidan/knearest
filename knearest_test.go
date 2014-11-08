package knearest

import (
  "testing"
)

type testComparison struct {
  id    int
  value float64
}

func (self testComparison) Compare(b Comparable) Comparison {
  other, ok := b.(testComparison)
  if ok {
    if self.value < other.value {
      return lt
    } else if self.value > other.value {
      return gt
    } else {
      return et
    }
  } else {
    panic("Deep inside distance comparisons, types went awry")
  }

}

var testData = []testComparison{
  testComparison{0, 26.841536502746},
  testComparison{1, 34.990746171675},
  testComparison{2, 12.288105912641},
  testComparison{3, 83.350508279796},
  testComparison{4, 12.13761624514},
  testComparison{5, 46.717980618923},
  testComparison{6, 82.157889465875},
  testComparison{7, 41.56469490452},
  testComparison{8, 68.41286964175},
  testComparison{9, 72.30240296214},
  testComparison{10, 47.388640068187},
  testComparison{11, 52.406732576157},
  testComparison{12, 50.383466971285},
  testComparison{13, 25.086573383346},
  testComparison{14, 0.28862706398993},
  testComparison{15, 98.351535107173},
  testComparison{16, 86.928612313666},
  testComparison{17, 2.5506918330447},
  testComparison{18, 72.762606187147},
  testComparison{19, 85.816451388326},
}

func TestList(t *testing.T) {
  kNearest := NewKNearestStorable(5)

  for _, comparison := range testData {
    kNearest.Add(comparison)

    list := kNearest.Get()
    if !isMonotonic(list) {
      t.Errorf("Added %v and list is no longer monotonic (%v)", comparison.value, list)

    }
  }
}

func isMonotonic(comparisons []Comparable) bool {
  last := -1.0
  for _, cmp := range comparisons {
    comparison, ok := cmp.(testComparison)
    if ok {
      if comparison.value >= last {
        last = comparison.value
      } else {
        return false
      }
    }
  }

  return true
}
