package arr

import (
  "math"
)

/**
 * Раскидывает самые большие кучи равномерно
 * не находит минимальный K, но имеет не плохой перфоманс
 */
func minEatingSpeedFlattening(piles []int, h int) int {
  var max int
  if len(piles) == h {
    max, _ = getMax(piles)
    return max
  }

  var over = h - len(piles)
  var result = make([]int, len(piles) + over)
  copy(result, piles)

  for over > 0 {
    var max, pos = getMax(result)
    var half = float64(max) / 2
    result[pos] = int(math.Ceil(half))
    result[len(piles) + over - 1] = int(math.Floor(half))
    over--
  }

  max, _ = getMax(result)

  return max
}

func minEatingSpeedGrowing(piles []int, h int) int {
  var max, i, pile int
  max, _ = getMax(piles)
  if len(piles) == h {
    return max
  }

  var k = int(max / h)
  var attempts = h

  for {
    for i, pile = range piles {
      frac := float64(pile) / float64(k)
      attempts -= int(math.Ceil(frac))
      if attempts < 0 {
        k++
        break;
      }
    }

    if i == len(piles) - 1 && attempts >= 0 {
      return k
    }

    attempts = h
  }

  return k
}

func getMax(piles []int) (max int, pos int) {
  for cur, pile := range piles {
    if pile > max {
      max = pile
      pos = cur
    }
  }

  return max, pos
}
