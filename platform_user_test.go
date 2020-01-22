package end_to_end

import (
	"testing"
)

func TestPlatformUsers(t *testing.T) {
	/*
		helper.MakeClient()

		t.Run("Sign in as admin", func(t *testing.T) {
			helper.Client.Post("/v1/api/users/auth").
				JSON(map[string]string{
					"email":    "admin@example.com",
					"password": "adminsecret",
				}).
				Expect(t).
				Status(200).
				Type("json").
				AssertFunc(helper.GetBody).
				Done()

			adminJwt, err = helper.UnmarshalStringData(BodyString)
			if err != nil {
				t.Error(err)
				return
			}
		})

		t.Run("Create a platform user", func(t *testing.T) {
			platformUserEmailAddress = MakeEmailAddress()
			Client.Post("/v1/api/users/platform").
				AddHeader("Authorization", fmt.Sprintf("Bearer %s", adminJwt)).
				JSON(map[string]string{"email": platformUserEmailAddress, "platform_role_id": "1"}).
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

		t.Run("Activate a platform user", func(t *testing.T) {
			Client.Post("/v1/api/users/platform/activate").
				AddHeader("Authorization", fmt.Sprintf("Bearer %s", adminJwt)).
				JSON(map[string]string{"token": verificationToken, "password": PASSWORD}).
				Expect(t).
				Status(200).
				Type("json").
				JSONSchema(UserSchema).
				AssertFunc(GetBody).
				Done()

			fmt.Printf("Activate response %s\n", BodyString)
		})

		t.Run("Sign in a platform user", func(t *testing.T) {
			Client.Post("/v1/api/users/auth").
				JSON(map[string]string{
					"email":    platformUserEmailAddress,
					"password": PASSWORD,
				}).
				Expect(t).
				Status(200).
				Type("json").
				//JSONSchema(userSchema).
				AssertFunc(GetBody).
				Done()

			fmt.Printf("Sign in response %s\n", BodyString)

			_, err = UnmarshalStringData(BodyString)
			if err != nil {
				t.Error(err)
				return
			}

		})*/
}
