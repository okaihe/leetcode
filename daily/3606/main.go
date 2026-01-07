package main

import "slices"

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	n := len(code)
	validIndices := make([]int, 0, n)

	for i := range n {
		if !isActive[i] {
			continue
		}

		if getPriority(businessLine[i]) == 0 {
			continue
		}

		if isValidCouponName(code[i]) {
			validIndices = append(validIndices, i)
		}
	}

	slices.SortFunc(validIndices, func(a, b int) int {
		p1 := getPriority(businessLine[a])
		p2 := getPriority(businessLine[b])

		if p1 != p2 {
			return p1 - p2
		}
		if code[a] < code[b] {
			return -1
		}
		if code[a] > code[b] {
			return 1
		}
		return 0
	})

	result := make([]string, len(validIndices))
	for i, idx := range validIndices {
		result[i] = code[idx]
	}

	return result
}

func getPriority(line string) int {
	switch line {
	case "electronics":
		return 1
	case "grocery":
		return 2
	case "pharmacy":
		return 3
	case "restaurant":
		return 4
	default:
		return 0
	}
}

func isValidCouponName(s string) bool {
	if s == "" {
		return false
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '_') {
			return false
		}
	}
	return true
}
