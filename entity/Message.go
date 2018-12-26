package entity

type MseeageEntity struct {
	UserName string `json:"name",db:"username"`
	Mseeage string `json:"mseeage",db:"message"` // `json:"-"` 这个字段不会序列化到json中 omitempty 如果字段为空则不会序列化到json中

}
