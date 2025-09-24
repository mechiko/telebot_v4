package entity

// для одного пользователя
type UserStatesMaps struct {
	UserId  int64
	Intro   map[string]string
	Examens map[string]string
}

// для всех пользователей мап от uid
type UsersStatesMap map[int64]UserStatesMaps
