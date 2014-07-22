package main

import (
	"testing"

	//"fmt"

	cv "github.com/smartystreets/goconvey/convey"
)

func TestCreateTrex(t *testing.T) {
	cv.Convey("Given a mysql with a gauntlet database and a trex table", t, func() {
		InjectDummyGoCDEnvironment()

		cv.Convey("we grab the GoCD environment vars and insert them as a row into the trex table", func() {
			env := GrabGoCDEnv()
			insertSucceeded := ExportEnvToTrex(env)
			row := GetLastRow()
			cv.So(row.Pipeline, cv.ShouldEqual, "pipen")
		})
	})
}
