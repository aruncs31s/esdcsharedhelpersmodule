package utils

func GetPaginationMeta(limit int, offset int, totalRecords uint) interface{} {
	return map[string]interface{}{
		"totalRecords": totalRecords,
		"totalPages":   getTotalPages(totalRecords, limit),
		"currentPage":  getCurrentPage(offset, limit),
		"pageSize":     limit,
	}
}
func getCurrentPage(offset int, limit int) int {
	if limit == 0 {
		return 0
	}
	return (offset / limit) + 1
}
func getTotalPages(totalRecords uint, limit int) int {
	totalRecordsInt := int(totalRecords)
	if limit == 0 {
		return 0
	}
	if totalRecordsInt%limit == 0 {
		return totalRecordsInt / limit
	}
	return (totalRecordsInt / limit) + 1
}
