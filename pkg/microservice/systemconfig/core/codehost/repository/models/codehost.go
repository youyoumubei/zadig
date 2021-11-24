package models

type CodeHost struct {
	ID            int    `bson:"id"               json:"id"`
	Type          string `bson:"type"             json:"type"`
	Address       string `bson:"address"          json:"address"`
	IsReady       string `bson:"is_ready"         json:"is_ready"`
	AccessToken   string `bson:"access_token"     json:"access_token"`
	RefreshToken  string `bson:"refresh_token"    json:"refresh_token"`
	Namespace     string `bson:"namespace"        json:"namespace"`
	ApplicationId string `bson:"application_id"   json:"application_id"`
	Region        string `bson:"region"           json:"region,omitempty"`
	Username      string `bson:"username"         json:"username,omitempty"`
	Password      string `bson:"password"         json:"password,omitempty"`
	ClientSecret  string `bson:"client_secret"    json:"client_secret"`
	CreatedAt     int64  `bson:"created_at"       json:"created_at"`
	UpdatedAt     int64  `bson:"updated_at"       json:"updated_at"`
	DeletedAt     int64  `bson:"deleted_at"       json:"deleted_at"`
}

func (CodeHost) TableName() string {
	return "code_host"
}
