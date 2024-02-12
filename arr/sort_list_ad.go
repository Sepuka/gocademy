package arr

//
// Есть два односвязных списка рекламной статистики (StatList),
// в каждом из которых элементы упорядочены по возрастанию даты.
// В них есть или показы (Shows) или клики (Clicks), но не оба значения одновременно.
// Длины списков и даты в них могут отличаться.
//
// Необходимо написать функцию, которая объединяет эти два списка по дате таким образом,
// чтобы результирующий список тоже был упорядочен по её возрастанию,
// но хранил в себе как показы, так и клики сразу.
//
// Пример:
// +---------------+    +---------------+    +---------------+
// |Date=2023-10-01|    |Date=2023-10-02|    |Date=2023-10-04|
// |Shows=1000     |--->|Shows=2000     |--->|Shows=3000     |
// |               |    |               |    |               |
// +---------------+    +---------------+    +---------------+
//
// +---------------+    +---------------+
// |Date=2023-10-01|    |Date=2023-10-04|
// |Clicks=250     |--->|Clicks=300     |
// |               |    |               |
// +---------------+    +---------------+
//
//	|
//	| CommonStatistics
//	v
//
// +---------------+    +---------------+    +---------------+
// |Date=2023-10-01|    |Date=2023-10-02|    |Date=2023-10-04|
// |Shows=1000     |--->|Shows=2000     |--->|Shows=3000     |
// |Clicks=250     |    |Clicks=0       |    |Clicks=300     |
// +---------------+    +---------------+    +---------------+
//
// Как проверять:
// go test -v

import (
	"fmt"
	"strings"
)

// StatList is a node of a linked list.
type StatList struct {
	Date   string
	Shows  uint64
	Clicks uint64
	Next   *StatList
}

// String returns a string representation of the list (helper-function).
func (s *StatList) String() string {
	if s == nil {
		return "nil"
	}

	builder := &strings.Builder{}
	builder.WriteString("StatList(")

	for node := s; node != nil; node = node.Next {
		builder.WriteString("[")
		builder.WriteString(node.Date)
		builder.WriteString(fmt.Sprintf("%d %d %p]", node.Shows, node.Clicks, node.Next))

		if node.Next != nil {
			builder.WriteString(", ")
		}
	}

	builder.WriteString(")")
	return builder.String()
}

func CommonStatistics(shows *StatList, clicks *StatList) *StatList {
	var res *StatList

	if shows == nil && clicks == nil {
		return nil
	}
	if shows == nil {
		return clicks
	}
	if clicks == nil {
		return shows
	}

	// находим самую ранюю статистику
	if shows.Date <= clicks.Date {
		res = shows
	} else {
		res = clicks
	}
	var head = res

	for res != nil {
		// если нет показов, возьмём их из другой структуры
		if res.Shows == 0 {
			// но только если дата такая же
			if res.Date == shows.Date {
				res.Shows = shows.Shows
				shows = shows.Next
			}
			// выберем наименьшую дату для ссылки на следующий узел
			if res.Next != nil {
				if res.Next.Date > shows.Date {
					res.Next = shows
				}
			} else {
				res.Next = shows

				return head
			}
		} else {
			if res.Date == clicks.Date {
				res.Clicks = clicks.Clicks
				clicks = clicks.Next
			}
			if res.Next != nil {
				if res.Next.Date > clicks.Date {
					res.Next = clicks
				}
			} else {
				res.Next = clicks

				return head
			}
		}

		res = res.Next
	}

	return head
}
