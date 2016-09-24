package database

import (
	"testing"

	"bitbucket.org/SummerCampDev/summercamp/models"
	"bitbucket.org/SummerCampDev/summercamp/models/forms"
	"bitbucket.org/SummerCampDev/summercamp/tests/setup"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUserModel(t *testing.T) {
	Convey("Test User model", t, func() {
		Convey("Test FetchAll", func() {
			c := setup.GetFixture("users").Count()
			users, ok := models.Users.FetchAll()
			So(ok, ShouldBeTrue)
			So(users, ShouldHaveLength, c)
		})
	})
}

func TestRegistrationForm(t *testing.T) {
	Convey("Test registration form", t, func() {
		ur := forms.UserRegistration{
			Email:           "valid@mail.com",
			Type:            models.SpecTypeExecutor,
			FirstName:       "oleh",
			LastName:        "holovko",
			Password:        "1234%",
			PasswordConfirm: "1234%",
			Country:         "country",
			City:            "city",
		}

		// TODO: add Delete method to user and remove user here if everything is okay so the tests with fixtures
		// can succeed
		Convey("Everything is okay", func() {
			ur.Type = models.SpecTypeExecutor

			user, ok := ur.Register()
			So(user, ShouldNotBeNil)
			So(ok, ShouldBeTrue)
			So(ur.Errors, ShouldHaveLength, 0)
		})

		Convey("Got invalid email", func() {
			ur.Email = "bad_email"
			user, ok := ur.Register()

			So(user, ShouldBeNil)
			So(ok, ShouldBeFalse)

			So(len(ur.Errors), ShouldNotEqual, 0)
		})

		Convey("Password and password confirmation mismatch", func() {
			ur.Password = "1"
			user, ok := ur.Register()

			So(user, ShouldBeNil)
			So(ok, ShouldBeFalse)
			So(len(ur.Errors), ShouldNotEqual, 0)
		})

		Convey("Invalid speciality", func() {
			ur.Type = "bad_spec"
			spec := models.Speciality(ur.Type)

			So(spec.Valid(), ShouldBeFalse)

			user, ok := ur.Register()

			So(user, ShouldBeNil)
			So(ok, ShouldBeFalse)
			So(len(ur.Errors), ShouldNotEqual, 0)
		})

		Convey("Very long password", func() {
			ur.Password = "%1111111111111111111111111111111111111111111111111111111111111111111111111111"
			ur.PasswordConfirm = "%1111111111111111111111111111111111111111111111111111111111111111111111111111"

			user, ok := ur.Register()

			So(user, ShouldBeNil)
			So(ok, ShouldBeFalse)
			So(len(ur.Errors), ShouldNotEqual, 0)
		})
	})

}
