package database

type Role string

const (
	Admin     Role = "admin"
	Moderator Role = "mod"
	Booster   Role = "booster"
)

type Entity struct {
	Role   Role   `bson:"role"`
	Member Member `bson:"member"`
}

type Member struct {
	Id        string `bson:"id"`
	Name      string `bson:"name"`
	AvatarURL string `bson:"avatarUrl"`
}

type General struct {
	ChannelCount   int      `bson:"channelCount"`
	MemberCount    int      `bson:"memberCount"`
	SpecialMembers []Entity `bson:"specialMembers"`
}

type InfoEntity struct {
	Bio        string   `bson:"bio"`
	Occupation string   `bson:"occupation"`
	Github     string   `bson:"github"`
	Skills     []string `bson:"skills"`
	MemberID   string   `bson:"member_id"`
}
