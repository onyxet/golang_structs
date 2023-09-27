package post

type PostSortDept struct {
	DeptName string
}

func (d *PostSortDept) Sort(p Parcel) {
	switch p.(type) {
	case *Box:
		p.Send()
	case *Envelope:
		p.Send()
	default:
		panic("unknown parcel type")
	}
}
