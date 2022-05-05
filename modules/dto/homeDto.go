package dto

type UserAssociation struct {
	//用户id
	Id string `json:"id"`
	//想要加入社团的id
	AssociationId string `json:"association_id"`
}

type AssociationId struct {
	Id string `json:"id"`
}
