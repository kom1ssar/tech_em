package api

type PersonV1Implementation interface {
	ById(id int64)
	List()
	Create()
	Update()
}
