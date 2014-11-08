knearest
========

Go package for finding the k-nearest (or largest/smallest) values in a list.

Usage:

    kNearestStorable := knearest.NewKNearestStorable(11)
    for _, comparison := range values {
      kNearest.Add(comparison)
    }

    kNearest := kNearestStorable.Get()


