package models

// Complaint Mapping

func MapComplaintToDTO(c Complaint) ComplaintDTO {
	return ComplaintDTO{
		ID:          c.ID,
		Title:       c.Title,
		Description: c.Description,
		UserID:      c.UserID,
		CategoryID:  c.CategoryID,
	}
}

func MapComplaintsToDTO(c []Complaint) []ComplaintDTO {
	var dtos []ComplaintDTO

	for _, complaint := range c {
		dtos = append(dtos, MapComplaintToDTO(complaint))
	}

	return dtos
}

// Category Mapping

func MapCategoryToDTO(c ComplaintCategory) CategoryDTO {
	return CategoryDTO{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
	}
}

func MapCategoriesToDTO(c []ComplaintCategory) []CategoryDTO {
	var dtos []CategoryDTO

	for _, category := range c {
		dtos = append(dtos, MapCategoryToDTO(category))
	}

	return dtos
}

// Comment Mapping

func MapCommentToDTO(c Comment) CommentDTO {
	return CommentDTO{
		ID:          c.ID,
		ComplaintID: c.ComplaintID,
		UserID:      c.UserID,
		CommentText: c.CommentText,
		CreatedAt:   c.CreatedAt,
	}
}

func MapCommentsToDTO(c []Comment) []CommentDTO {
	var dtos []CommentDTO

	for _, comment := range c {
		dtos = append(dtos, MapCommentToDTO(comment))
	}

	return dtos
}
