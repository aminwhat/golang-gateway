package db_package

type collectionInterface interface {
	Create()
	Find()
	Update()
	Delete()
}
