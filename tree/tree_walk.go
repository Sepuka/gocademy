package tree

import "errors"

// Обходит дерево в глубину и возвращает среднее арифметическое значений всех листьев заданного уровня
func WalkAvgLeafRecursive(tree *Tree, level uint) (float32, error) {
	if level == 0 {
		return float32(tree.value), nil
	}

	if level == 1 {
		if tree.left != nil && tree.right != nil {
			return float32(tree.left.value+tree.right.value) / 2, nil
		} else if tree.left != nil {
			return float32(tree.left.value), nil
		} else if tree.right != nil {
			return float32(tree.right.value), nil
		} else {
			return 0, errors.New(`too deep`)
		}
	}

	left, err := WalkAvgLeafRecursive(tree.left, level-1)
	if err != nil {
		return 0, err
	}

	right, err := WalkAvgLeafRecursive(tree.right, level-1)
	if err != nil {
		return 0, err
	}

	return float32(left+right) / 2, nil
}

// Обходит дерево запоминая значения каждого уровня
// возвращает среднее арифметическое этого уровня
func WalkAvgLeafIter(tree *Tree, level uint) (res float32, err error) {
	var queue = map[uint][]*Tree{}
	queue[0] = []*Tree{tree}
	var lvl uint = 0

	for len(queue[lvl]) > 0 {
		var row = []*Tree{}
		for _, tree = range queue[lvl] {
			if tree.left != nil {
				row = append(row, tree.left)
			}
			if tree.right != nil {
				row = append(row, tree.right)
			}
		}

		lvl++
		if len(row) > 0 {
			queue[lvl] = row
		}
	}

	if val, ok := queue[level]; ok {
		for _, num := range val {
			res += float32(num.value)
		}

		return res / float32(len(val)), nil
	}

	return 0, errors.New(`too deep`)
}

// Возвращает среднюю глубину дерева
func WalkAvgDeep(tree *Tree) float32 {
	if tree.left == nil && tree.right == nil {
		return 0
	}

	return float32(deep(tree.left, 1)+deep(tree.right, 1)) / 2
}

func deep(tree *Tree, cnt float32) float32 {
	if tree == nil {
		return 0
	}

	if tree.left == nil && tree.right == nil {
		return cnt
	}

	cnt++

	if tree.left != nil && tree.right != nil {
		return float32(deep(tree.left, cnt)+deep(tree.right, cnt)) / 2
	}

	if tree.left != nil {
		return deep(tree.left, cnt)
	}

	return deep(tree.right, cnt)
}
