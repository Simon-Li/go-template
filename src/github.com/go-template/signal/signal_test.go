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

var _ = Describe("Test signal handling", func() {

	Context("test basic signal handling works", func() {
		It("handles basic signal operations", func() {
			// ask for SIGHUP
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGHUP)
			defer signal.Stop(c)

			// send this process a SIGHUP
			fmt.Println("sighup... on c")
			syscall.Kill(syscall.Getpid(), syscall.SIGHUP)
			waitSig(c, syscall.SIGHUP)

			c1 := make(chan os.Signal, 1)
			signal.Notify(c1)

			// send this process a SIGWINCH
			fmt.Println("sigwinch... on c1")
			syscall.Kill(syscall.Getpid(), syscall.SIGWINCH)
			waitSig(c1, syscall.SIGWINCH)

			// Send two more SIGHUPs, to make sure that
			// they get delivered on c1 and that not reading
			// from c does not block everything.
			fmt.Println("sighup... on c1")
			syscall.Kill(syscall.Getpid(), syscall.SIGHUP)
			waitSig(c1, syscall.SIGHUP)
			fmt.Println("sighup... on c1")
			syscall.Kill(syscall.Getpid(), syscall.SIGHUP)
			waitSig(c1, syscall.SIGHUP)

			// The first SIGHUP should be waiting for us on c.
			fmt.Println("sighup... on c")
			waitSig(c, syscall.SIGHUP)
		})
	})
})
