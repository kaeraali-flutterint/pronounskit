package lib

import "time"

type User struct {
	ID                 string        `json:"id"`
	FirstName          string        `json:"first_name"`
	LastName           string        `json:"last_name"`
	Email              string        `json:"email"`
	Type               int           `json:"type"`
	RoleName           string        `json:"role_name"`
	Pmi                int64         `json:"pmi"`
	UsePmi             bool          `json:"use_pmi"`
	PersonalMeetingURL string        `json:"personal_meeting_url"`
	Timezone           string        `json:"timezone"`
	Verified           int           `json:"verified"`
	Dept               string        `json:"dept"`
	CreatedAt          time.Time     `json:"created_at"`
	LastLoginTime      time.Time     `json:"last_login_time"`
	LastClientVersion  string        `json:"last_client_version"`
	PicURL             string        `json:"pic_url"`
	HostKey            string        `json:"host_key"`
	CmsUserID          string        `json:"cms_user_id"`
	Jid                string        `json:"jid"`
	GroupIds           []interface{} `json:"group_ids"`
	ImGroupIds         []interface{} `json:"im_group_ids"`
	AccountID          string        `json:"account_id"`
	Language           string        `json:"language"`
	PhoneCountry       string        `json:"phone_country"`
	PhoneNumber        string        `json:"phone_number"`
	Status             string        `json:"status"`
	JobTitle           string        `json:"job_title"`
	Company            string        `json:"company"`
	Location           string        `json:"location"`
	LoginTypes         []int         `json:"login_types"`
	RoleID             string        `json:"role_id"`
	AccountNumber      int           `json:"account_number"`
	Pronouns           string        `json:"pronouns"`
	Cluster            string        `json:"cluster"`
	PronounsOption     int           `json:"pronouns_option"`
}
