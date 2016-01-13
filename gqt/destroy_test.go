package gqt_test

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/cloudfoundry-incubator/garden"
	"github.com/cloudfoundry-incubator/guardian/gqt/runner"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Destroying a Container", func() {
	var (
		client    *runner.RunningGarden
		container garden.Container
	)

	BeforeEach(func() {
		client = startGarden()
	})

	AfterEach(func() {
		Expect(client.DestroyAndStop()).To(Succeed())
	})

	JustBeforeEach(func() {
		Expect(client.Destroy(container.Handle())).To(Succeed())
	})

	Context("when running a process", func() {
		var (
			process     garden.Process
			initProcPid int
		)

		BeforeEach(func() {
			var err error

			container, err = client.Create(garden.ContainerSpec{})
			Expect(err).NotTo(HaveOccurred())

			initProcPid = initProcessPID(container.Handle())

			process, err = container.Run(garden.ProcessSpec{
				Path: "/bin/sh",
				Args: []string{
					"-c", "read x",
				},
			}, ginkgoIO)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should kill the containers init process", func() {
			var killExitCode = func() int {
				sess, err := gexec.Start(exec.Command("kill", "-0", fmt.Sprintf("%d", initProcPid)), GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				sess.Wait(1 * time.Second)
				return sess.ExitCode()
			}

			Eventually(killExitCode, "5s").Should(Equal(1))
		})

		It("should destroy the container's depot directory", func() {
			Expect(filepath.Join(client.DepotDir, container.Handle())).NotTo(BeAnExistingFile())
		})

		It("should destroy the container rootfs", func() {
			session, err := gexec.Start(exec.Command("du", "-d0", client.GraphPath), GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session).Should(gbytes.Say(`^0\s+`))
		})
	})

	Context("when using a static subnet", func() {
		var (
			contIfaceName     string
			contHandle        string
			existingContainer garden.Container
		)

		BeforeEach(func() {
			var err error

			container, err = client.Create(garden.ContainerSpec{
				Network: "177.100.10.30/24",
			})
			Expect(err).NotTo(HaveOccurred())
			contIfaceName = ethInterfaceName(container)
			contHandle = container.Handle()

			existingContainer, err = client.Create(garden.ContainerSpec{
				Network: "168.100.20.10/24",
			})
			Expect(err).NotTo(HaveOccurred())
		})

		AfterEach(func() {
			Expect(client.Destroy(existingContainer.Handle())).To(Succeed())
		})

		It("should remove iptable entries", func() {
			out, err := exec.Command("iptables", "-S", "-t", "filter").CombinedOutput()
			Expect(err).NotTo(HaveOccurred())
			Expect(string(out)).NotTo(ContainSubstring("177.100.10.0/24"))
			Expect(string(out)).To(ContainSubstring("168.100.20.0/24"))
		})

		It("should remove namespaces", func() {
			session, err := gexec.Start(
				exec.Command("ip", "netns", "list"),
				GinkgoWriter, GinkgoWriter,
			)
			Expect(err).NotTo(HaveOccurred())
			Consistently(session).ShouldNot(gbytes.Say(contHandle))
			Expect(session).Should(gbytes.Say(existingContainer.Handle()))
		})

		It("should remove virtual ethernet cards", func() {
			ifconfigExits := func() int {
				session, err := gexec.Start(exec.Command("ifconfig", contIfaceName), GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())

				return session.Wait().ExitCode()
			}
			Eventually(ifconfigExits).ShouldNot(Equal(0))

			ifaceName := ethInterfaceName(existingContainer)
			session, err := gexec.Start(exec.Command("ifconfig", ifaceName), GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session).Should(gexec.Exit(0))
		})

		It("should remove the network bridge", func() {
			session, err := gexec.Start(
				exec.Command("ifconfig"),
				GinkgoWriter, GinkgoWriter,
			)
			Expect(err).NotTo(HaveOccurred())
			Consistently(session).ShouldNot(gbytes.Say("br-177-100-10-0"))

			session, err = gexec.Start(
				exec.Command("ifconfig"),
				GinkgoWriter, GinkgoWriter,
			)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session).Should(gbytes.Say("br-168-100-20-0"))
		})
	})
})

func ethInterfaceName(container garden.Container) string {
	buffer := gbytes.NewBuffer()
	proc, err := container.Run(
		garden.ProcessSpec{
			Path: "sh",
			Args: []string{"-c", "ifconfig | grep 'Ethernet' | cut -f 1 -d ' '"},
			User: "root",
		},
		garden.ProcessIO{
			Stdout: buffer,
			Stderr: GinkgoWriter,
		},
	)
	Expect(err).NotTo(HaveOccurred())
	Expect(proc.Wait()).To(Equal(0))

	contIfaceName := string(buffer.Contents()) // w3-abc-1

	return contIfaceName[:len(contIfaceName)-2] + "0" // w3-abc-0
}
