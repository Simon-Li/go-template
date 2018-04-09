package main_test

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Integration Tests", func() {
	var (
		client *http.Client
	)

	BeforeEach(func() {
		client = newHTTPClient()
	})

	It("responds to web-site with valid HTML", func() {
		resp, err := client.Get(fmt.Sprintf("http://%s/", "115.28.28.127"))
		Expect(err).NotTo(HaveOccurred())

		body, err := ioutil.ReadAll(resp.Body)
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.Body.Close()).To(Succeed())
		//Expect(string(body)).To(ContainSubstring("<html>"))
		Expect(string(body)).To(ContainSubstring("</html>"))
	})

	It("Expect variable increment", func() {
		var i int = 999
		i++
		Expect(i).To(Equal(999 + 1))
	})
})

func newHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}
