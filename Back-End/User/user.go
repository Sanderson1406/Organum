package users

type Profile struct {}

type Metadata struct{
	data map[string]string
}

func NewMetadata() *Metadata{
	return &Metadata{data: make(map[string]string)}
}

func (m *Metadata) Add(key, value string) {
	m.data[key] = value
}

func (m *Metadata) Get(key string) (string, bool) {
	return  m.data[key]
}

type User struct{
	ID	string
	Name	string
	Email	string
	Password	string
	Profile: *Profile
	Metadata: Metadata
}

func NewUser(id, name, enail, password string, profile *Profile, m *Metadata) *User {
	return &User{
		ID: id,
		Name: name,
		Email: email,
		Password: password,
		Profile: profile,
		Metadada: m,
	}
}

func (u *User) Authenticate(password string) bool {
	return u.Password ==password
}

func (u *User) SetProfile(p *Profile) {
	u.Profile = p
}