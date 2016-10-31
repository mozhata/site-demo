package api

import (
	"fmt"
	"net/http"
	"strings"

	"bitbucket.org/applysquare/applysquare-go/pkg/common"
)

type User struct {
	id                     string
	permissions            map[string]bool
	editableInstituteSlugs map[string]bool

	// if true, the user has some biz permission. Furthermore, the user has permission tags like "biz:cn.tsinghua:admin".
	hasSomeBizPermission bool
	// If true, the user is a tag admin. The user has permission tags like "tag_staff:fancytag:admin"
	isTagStaff bool
}

func NewUserFromRequest(req *http.Request) *User {
	id := req.Header.Get("X-A2-USER-UUID")
	permissions := make(map[string]bool)

	isTagStaff := false
	isBiz := false
	for _, p := range common.SplitAndTrim(req.Header.Get("X-A2-USER-PERMISSIONS"), ",") {
		perm := common.SplitAndTrim(p, ":")
		if perm[0] == "tag_staff" {
			isTagStaff = true
		}
		if perm[0] == "biz" {
			isBiz = true
		}
		permissions[p] = true
	}
	editableInstituteSlugs := make(map[string]bool)
	for _, s := range common.SplitAndTrim(req.Header.Get("X-A2-USER-EDITABLE-INSTITUTE-SLUGS"), ",") {
		editableInstituteSlugs[s] = true
	}
	return &User{
		id:                     id,
		permissions:            permissions,
		editableInstituteSlugs: editableInstituteSlugs,
		hasSomeBizPermission:   isBiz,
		isTagStaff:             isTagStaff,
	}
}

func (u *User) ID() string {
	return u.id
}

func (u *User) HasPermission(p string) bool {
	_, ok := u.permissions[p]
	return ok
}

func (u *User) HasSomeEditableInstituteSlug() bool {
	if len(u.editableInstituteSlugs) > 0 {
		return true
	}
	return false
}

func (u *User) HasEditableInstituteSlug(s string) bool {
	_, ok := u.editableInstituteSlugs[s]
	if ok {
		return true
	}
	countrySlug := func(slug string) string {
		countryKey := "us"
		if parts := strings.Split(slug, "."); len(parts) > 1 {
			countryKey = parts[0]
		}
		return fmt.Sprintf("country:%s", countryKey)
	}(s)
	_, ok = u.editableInstituteSlugs[countrySlug]
	return ok
}

func (u *User) IsStaff() bool {
	return u.HasPermission("staff")
}

func (u *User) IsDev() bool {
	return u.HasPermission("dev")
}

func (u *User) IsMsgStaff() bool {
	return u.HasPermission("msg_staff")
}

func (u *User) IsDiscussionStaff() bool {
	return u.HasPermission("qa_staff") || u.IsStaff()
}

func (u *User) IsDituiStaff() bool {
	return u.HasPermission("ditui_staff") || u.IsStaff()
}

func (u *User) IsConfigFosStaff() bool {
	return u.HasPermission("fos_staff") || u.IsStaff()
}

func (u *User) IsConfigCareerStaff() bool {
	return u.HasPermission("career_staff") || u.IsStaff()
}

func (u *User) IsConfigCourseStaff() bool {
	return u.HasPermission("course_staff") || u.IsStaff()
}

func (u *User) IsHomeStaff() bool {
	return u.HasPermission("home_staff") || u.IsStaff()
}

func (u *User) HasAnyTagPermission() bool {
	return u.isTagStaff
}

func (u *User) IsTagEditor(tag string) bool {
	return u.IsTagAdmin(tag) || u.HasPermission("tag_staff:_all:editor") || u.HasPermission("tag_staff:"+tag+":editor")
}

func (u *User) IsTagAdmin(tag string) bool {
	return u.HasPermission("tag_staff:_all:admin") || u.HasPermission("tag_staff:"+tag+":admin")
}

func (u *User) HasSomeBizPermission() bool {
	return u.hasSomeBizPermission
}

func (u *User) BizSlugs() []string {
	slugs := []string{}
	for p, _ := range u.permissions {
		perm := common.SplitAndTrim(p, ":")
		if perm[0] == "biz" {
			slugs = append(slugs, perm[1])
		}
	}
	return slugs
}

func (u *User) IsBizAdmin(slug string) bool {
	return u.HasPermission("biz:" + slug + ":admin")
}
