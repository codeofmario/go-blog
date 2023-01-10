package util

func FromModelToDtoList[MODEL any, DTO any](models []*MODEL, transform func(*MODEL) *DTO) []*DTO {
	dtoList := make([]*DTO, len(models))

	for i, model := range models {
		dtoList[i] = transform(model)
	}

	return dtoList
}
