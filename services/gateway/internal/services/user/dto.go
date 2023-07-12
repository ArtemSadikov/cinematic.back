package user

import (
	"cinematic.back/api/users/pb"
	"github.com/google/uuid"
	"time"
)

type Profile struct {
	Email    string
	Username string
}

func (p *Profile) fromPB(data *pb.UserProfile) {
	p.Email = data.Email
	p.Username = data.Username
}

type User struct {
	Id        uuid.UUID
	Profile   Profile
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (u *User) fromPB(data *pb.User) {
	u.Id = uuid.MustParse(data.Id)

	u.Profile.fromPB(data.Profile)

	u.CreatedAt = data.CreatedAt.AsTime()
	u.UpdatedAt = data.UpdatedAt.AsTime()

	if data.DeletedAt.IsValid() {
		deletedAt := data.DeletedAt.AsTime()
		u.DeletedAt = &deletedAt
	}
}
