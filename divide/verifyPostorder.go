package divide

const invalid = -1

func verifyPostorder(postOrder []int) bool {
	if len(postOrder) <= 1 {
		return true
	}

	boudary := invalid
	root := postOrder[len(postOrder)-1]
	for index, v := range postOrder {
		if v > root {
			boudary = index
		}
	}

	res := true
	if boudary != invalid {
		for i := boudary; i < len(postOrder); i++ {
			if postOrder[i] <= root {
				if i < len(postOrder)-1 {
					return false
				}
			}
		}
	}

	if boudary == invalid {
		res = res && verifyPostorder(postOrder[0:len(postOrder)-1])
		if !res {
			return false
		}
	} else {
		res = res && verifyPostorder(postOrder[0:boudary]) && verifyPostorder(postOrder[boudary:len(postOrder)-1])
	}

	return res
}
