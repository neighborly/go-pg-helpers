package pghelpers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	pgh "github.com/neighborly/go-pghelpers"
)

var _ = Describe("Connection Test", func() {

	var (
		testConfig = pgh.PostgresConfig{
			Host:       "localhost",
			Port:       5432,
			Username:   "postgres",
			Password:   "",
			Database:   "postgres",
			SSLEnabled: false,
		}
	)

	It("should connect to a database", func() {
		db, err := pgh.ConnectPostgres(testConfig)
		Expect(db).To(Not(BeNil()))
		Expect(err).To(BeNil())
		Expect(db.Ping()).To(Succeed())
	})
})
