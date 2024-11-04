package datatypes

type actor struct{
	Id string
	Login string
	Display_login string
	Gravatar
}

type Userdata struct {
	Id string
	Type string
	Actor actor


}