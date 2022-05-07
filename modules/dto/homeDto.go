package dto

type UserAssociation struct {
	//用户id
	Id string `json:"id"`
	//想要加入社团的id
	AssociationId string `json:"association_id"`
}

type AssociationId struct {
	//社团id
	Id string `json:"id"`
}

type SendNoticeDto struct {
	Title         string `json:"title"`
	Content       string `json:"content"`
	AssociationId string `json:"association_id"`
}

type AssociationContent struct {
	College string `json:"college"`
	Role    string `json:"role"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
}
