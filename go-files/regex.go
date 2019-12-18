package myregex

func matchOne(pattern, input byte) bool {
  if pattern == ' ' {
	  return true
  }

  if input == ' ' {
    return false
  }

  if pattern == '.' {
    return true
  }
  return pattern == input
}

func matchEqualLen(pattern, input string) bool {
  if pattern == "" {
    return true
  }

  return matchOne(pattern[0], input[0]) && matchEqualLen(pattern[1:], input[1:])
}