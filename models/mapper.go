package models

//Complaint Mapping

func MapComplaintToDTO(c Complaint) ComplaintDTO {
	return ComplaintDTO{
		ID:          c.ID,
		Title:       c.Title,
		Description: c.Description,
		CustomerID:  c.CustomerID,
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
