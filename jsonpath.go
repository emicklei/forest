package rat

import "strconv"

func pathFindIn(index int, tokens []string, here interface{}) interface{} {
	//.Printf("%d %q %d, %v\n", index, tokens, len(tokens), here)
	if here == nil {
		return here
	}
	if index == len(tokens) {
		return here
	}
	token := tokens[index]
	if len(token) == 0 {
		return here
	}
	i, err := strconv.Atoi(token)
	if err == nil {
		// try index into array
		array, ok := here.([]interface{})
		if ok {
			if i >= len(array) {
				return nil
			}
			return pathFindIn(index+1, tokens, array[i])
		}
		return nil
	}
	// try key into hash
	hash, ok := here.(map[string]interface{})
	if ok {
		return pathFindIn(index+1, tokens, hash[token])
	}
	return nil
}
