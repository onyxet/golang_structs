package post

type PostSortDept struct {
	DeptName string
}

func (d *PostSortDept) Sort(p Parcel) {
	p.Send()
}
