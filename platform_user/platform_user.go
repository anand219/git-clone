package platform_user

import (
	"fmt"
	"testing"
)

func CreatePlatfromUser(emailAddress, t *testing.T){
	t.Run("Create a platform user", func(t *testing.T) {
		Client.Post("/v1/api/users/platform").
			AddHeader("Authorization", fmt.Sprintf("Bearer %s", adminJwt)).
			JSON(map[string]string{"email": emailAddress, "platform_role_id": "1"}).
			Expect(t).
			Status(200).
			Type("json").
			JSONSchema(UserSchema).
			AssertFunc(GetBody).
			Done()

		userData, err := UnmarshalUserData(BodyString)
		verificationToken = userData.Data.VerificationToken_ //In TEST mode, the verification token is returned in the response instead of being sent in an email

		if err != nil {
			t.Error(err)
			return
		}

	})
}

t



