package signal_test

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func waitSig(c <-chan os.Signal, sig os.Signal) {
	select {
	case s := <-c:
		Expect(s).To(Equal(sig))
	case <-time.After(1 * time.Second):
		fmt.Printf("timeout waiting for %v", sig)
	}
}

var _ = Describe("TestGinServer", func() {

	Context("test basic signal handling works", func() {
		It("handles basic signal operations", func() {
			// ask for SIGHUP
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGHUP)
			defer signal.Stop(c)

			// send this process a SIGHUP
			fmt.Println("sighup...")
			syscall.Kill(syscall.Getpid(), syscall.SIGHUP)
			waitSig(c, syscall.SIGHUP)

			c1 := make(chan os.Signal, 1)
			signal.Notify(c1)

			// send this process a SIGWINCH
			fmt.Println("sigwinch...")
			syscall.Kill(syscall.Getpid(), syscall.SIGWINCH)
			waitSig(c1, syscall.SIGWINCH)

			fmt.Println("sighup...")
			syscall.Kill(syscall.Getpid(), syscall.SIGHUP)
			waitSig(c1, syscall.SIGHUP)
			fmt.Println("sighup...")
			syscall.Kill(syscall.Getpid(), syscall.SIGHUP)
			waitSig(c1, syscall.SIGHUP)
		})
	})
})
