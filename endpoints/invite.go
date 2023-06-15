package endpoints

import "fmt"

const (
	InvitesEndpoint = Endpoint + "invites"
)

func Invite(inviteCode string) string {
	return fmt.Sprintf("%s/%s", InvitesEndpoint, inviteCode)
}
