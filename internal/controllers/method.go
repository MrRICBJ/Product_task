package controllers

import (
	"strconv"
)

func GetLimOff(limStr, offStr string) (int32, int32, error) {
	var limit, offset int
	var err error

	if limStr == "" {
		limit = 1
	} else {
		limit, err = strconv.Atoi(limStr)
		if err != nil {
			return 0, 0, err
		}
	}

	if offStr != "" {
		offset, err = strconv.Atoi(offStr)
		if err != nil {
			return 0, 0, err
		}
	}

	return int32(limit), int32(offset), nil
}
